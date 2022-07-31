package expressions

import (
	"OLC2/environment"
	"OLC2/interfaces"
	"fmt"
	"math"
)

type Sqrt struct {
	Lin int
	Col int
	Exp interface{}
}

func NewSqrt(lin int, col int, ex interface{}) Sqrt {
	expr := Sqrt{lin, col, ex}
	return expr
}

func (p Sqrt) Ejecutar(ast *environment.AST, env interface{}) environment.Symbol {
	var result environment.Symbol

	tmpExp := p.Exp.(interfaces.Expression).Ejecutar(ast, env)

	if tmpExp.Tipo == environment.INTEGER {
		var newVal float64
		newVal = math.Sqrt(float64(tmpExp.Valor.(int)))
		result = environment.Symbol{
			Lin:     p.Lin,
			Col:     p.Col,
			Id:      "",
			Tipo:    tmpExp.Tipo,
			Valor:   int(newVal),
			Mutable: true,
		}
	} else if tmpExp.Tipo == environment.FLOAT {
		var newVal float64
		newVal = math.Sqrt(tmpExp.Valor.(float64))
		result = environment.Symbol{
			Lin:     p.Lin,
			Col:     p.Col,
			Id:      "",
			Tipo:    tmpExp.Tipo,
			Valor:   newVal,
			Mutable: true,
		}
	} else {
		fmt.Println("Tipo de dato incorrecto")
		return tmpExp
	}

	return result
}
