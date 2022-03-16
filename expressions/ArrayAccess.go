package expressions

import (
	"OLC2/environment"
	"OLC2/interfaces"
	"fmt"

	arrayList "github.com/colegno/arraylist"
)

type ArrayAccess struct {
	Lin   int
	Col   int
	Array interfaces.Expression
	Index interfaces.Expression
}

func NewArrayAccess(lin int, col int, array interfaces.Expression, index interfaces.Expression) ArrayAccess {
	exp := ArrayAccess{lin, col, array, index}
	return exp
}

func (p ArrayAccess) Ejecutar(ast *environment.AST, env interface{}) environment.Symbol {

	var tempArray environment.Symbol
	tempArray = p.Array.Ejecutar(ast, env) //sym{[1,2,3,4]}

	if tempArray.Tipo == environment.ARRAY {
		var tempIndex environment.Symbol
		tempIndex = p.Index.Ejecutar(ast, env) //sym{3}

		var tempValue interface{}
		tempValue = tempArray.Valor

		return tempValue.(*arrayList.List).GetValue(tempIndex.Valor.(int)).(environment.Symbol)
	}
	fmt.Println("No es un arreglo")

	return environment.Symbol{
		Lin:   p.Lin,
		Col:   p.Col,
		Id:    "",
		Tipo:  environment.NULL,
		Valor: "ERROR",
	}
}
