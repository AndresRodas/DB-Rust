package instructions

import (
	"OLC2/environment"
	"OLC2/interfaces"
	"fmt"
	arrayList "github.com/colegno/arraylist"
)

type Push struct {
	Lin       int
	Col       int
	Id        string
	Expresion interfaces.Expression
}

func NewPush(lin int, col int, id string, exp interfaces.Expression) Push {
	instr := Push{lin, col, id, exp}
	return instr
}

func (p Push) Ejecutar(ast *environment.AST, env interface{}) environment.Symbol {
	var result, tmpSymbol environment.Symbol
	tmpSymbol = env.(environment.Environment).GetVariable(p.Id) //variable guardada
	tmpExp := p.Expresion.Ejecutar(ast, env)
	//validar que id sea vector
	if tmpSymbol.Tipo == environment.VECTOR {
		//validar tipo de expresion
		if tmpExp.Tipo == tmpSymbol.TipoArr || tmpSymbol.Id == tmpExp.Id {
			//push valor
			if tmpSymbol.Mutable {
				//setear capacidad
				if tmpSymbol.Valor.(*arrayList.List).Len() == tmpSymbol.Capacity {
					tmpSymbol.Capacity = tmpSymbol.Capacity * 2
				}
				tmpSymbol.Valor.(*arrayList.List).Add(tmpExp)
			} else {
				fmt.Println("La variable no es mutable")
				return result
			}
		} else {
			fmt.Println("El tipo de dato es incorrecto")
			return result
		}
	} else {
		fmt.Println("Solo se puede hacer push sobre vectores")
		return result
	}
	//setear variable
	env.(environment.Environment).SetVariable(p.Id, tmpSymbol)

	return result
}
