package expressions

import (
	"OLC2/environment"
	"OLC2/interfaces"
	"fmt"
	arrayList "github.com/colegno/arraylist"
)

type Range struct {
	Lin  int
	Col  int
	Exp1 interfaces.Expression
	Exp2 interfaces.Expression
	Tipo environment.TipoExpresion
}

func NewRange(lin int, col int, expuno interfaces.Expression, expdos interfaces.Expression) Range {
	exp := Range{lin, col, expuno, expdos, 0}
	return exp
}

func (p Range) Ejecutar(ast *environment.AST, env interface{}) environment.Symbol {

	var val1, val2 environment.Symbol
	var arr *arrayList.List
	arr = arrayList.New()
	val1 = p.Exp1.Ejecutar(ast, env)
	val2 = p.Exp2.Ejecutar(ast, env)
	p.Tipo = val1.Tipo

	if (val1.Valor.(int) < val2.Valor.(int)) && (val1.Tipo == environment.INTEGER) && (val2.Tipo == environment.INTEGER) {

		cont := val1.Valor.(int)
		arr.Add(cont)
		for {
			if cont < val2.Valor.(int) {
				cont++
				arr.Add(cont)
			} else {
				break
			}
		}
	} else {
		fmt.Println("Indices incorrectos")
	}

	return environment.Symbol{
		Lin:     p.Lin,
		Col:     p.Col,
		Id:      "RANGE",
		Tipo:    environment.ARRAY,
		Valor:   arr,
		Mutable: false,
	}
}
