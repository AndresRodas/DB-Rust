package expressions

import (
	"OLC2/environment"
	"OLC2/interfaces"
	"fmt"
	"math"
)

type Pow struct {
	Lin   int
	Col   int
	Tipo1 environment.TipoExpresion
	Tipo2 environment.TipoExpresion
	Exp1  interface{}
	Exp2  interface{}
}

func NewPow(lin int, col int, tipo1, tipo2 environment.TipoExpresion, exp1, exp2 interface{}) Pow {
	exp := Pow{lin, col, tipo1, tipo2, exp1, exp2}
	return exp
}

func (p Pow) Ejecutar(ast *environment.AST, env interface{}) environment.Symbol {
	var result environment.Symbol
	e1 := p.Exp1.(interfaces.Expression).Ejecutar(ast, env)
	e2 := p.Exp2.(interfaces.Expression).Ejecutar(ast, env)

	if p.Tipo1 == p.Tipo2 && e1.Tipo == p.Tipo1 && e1.Tipo == e2.Tipo {
		if p.Tipo1 == environment.INTEGER {
			val := math.Pow(float64(e1.Valor.(int)), float64(e2.Valor.(int)))
			result = environment.Symbol{Lin: p.Lin, Col: p.Col, Tipo: p.Tipo1, Valor: int(val)}
		} else {
			val := math.Pow(e1.Valor.(float64), e2.Valor.(float64))
			result = environment.Symbol{Lin: p.Lin, Col: p.Col, Tipo: p.Tipo1, Valor: val}
		}
	} else {
		fmt.Println("Los tipos indicados son incorrectos")
	}
	return result
}
