package expressions

import (
	"OLC2/environment"
	"OLC2/interfaces"
	"fmt"
	arrayList "github.com/colegno/arraylist"
)

type Remove struct {
	Lin       int
	Col       int
	Id        string
	Expresion interfaces.Expression
}

func NewRemove(lin int, col int, id string, exp interfaces.Expression) Remove {
	expr := Remove{lin, col, id, exp}
	return expr
}

func (p Remove) Ejecutar(ast *environment.AST, env interface{}) environment.Symbol {
	var result, tmpSymbol, tmpExp environment.Symbol
	var tmpList *arrayList.List
	tmpList = arrayList.New()                                   //lista de los nuevos valores
	tmpSymbol = env.(environment.Environment).GetVariable(p.Id) //variable tipo vector
	tmpExp = p.Expresion.Ejecutar(ast, env)                     //indice

	//validar que id sea vector
	if tmpSymbol.Tipo == environment.VECTOR {
		//validar tipo de expresion
		if tmpExp.Tipo == environment.INTEGER {
			//validar indice
			if tmpExp.Valor.(int) < tmpSymbol.Valor.(*arrayList.List).Len() && tmpExp.Valor.(int) >= 0 {
				//recorrer indice
				for i := 0; i < tmpSymbol.Valor.(*arrayList.List).Len(); i++ {
					//asignar
					if tmpExp.Valor.(int) == i {
						result = tmpSymbol.Valor.(*arrayList.List).GetValue(i).(environment.Symbol)
					} else {
						tmpList.Add(tmpSymbol.Valor.(*arrayList.List).GetValue(i))
					}
				}
			} else {
				fmt.Println("Indice fuera del rango")
				return result
			}
		} else {
			fmt.Println("Tipo de indice incorrecto")
			return result
		}
	} else {
		fmt.Println("Solo se puede hacer Remove sobre vectores")
		return result
	}

	//remplazar lista
	tmpSymbol.Valor = tmpList
	//setear variable
	env.(environment.Environment).SetVariable(p.Id, tmpSymbol)
	return result
}
