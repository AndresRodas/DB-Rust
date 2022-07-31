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

	if tempArray.Tipo == environment.ARRAY || tempArray.Tipo == environment.VECTOR {
		var tempIndex environment.Symbol
		tempIndex = p.Index.Ejecutar(ast, env) //sym{3}
		var tempValue interface{}
		tempValue = tempArray.Valor
		//validando indice
		if tempIndex.Valor.(int) >= 0 && tempIndex.Valor.(int) <= tempValue.(*arrayList.List).Len() {
			valret := tempValue.(*arrayList.List).GetValue(tempIndex.Valor.(int)).(environment.Symbol)
			return valret
		} else {
			fmt.Println("Indice fuera de los limites ", p.Lin, p.Col)
			fmt.Println("indice: ", tempIndex.Valor.(int))
			fmt.Println("tamaño: ", tempValue.(*arrayList.List).Len())
		}

	} else {
		fmt.Println("No es un arreglo")
	}
	return environment.Symbol{
		Lin:     p.Lin,
		Col:     p.Col,
		Id:      "",
		Tipo:    environment.NULL,
		Valor:   0,
		Mutable: true,
	}
}
