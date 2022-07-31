package instructions

import (
	"OLC2/environment"
	"OLC2/interfaces"
	"fmt"
	arrayList "github.com/colegno/arraylist"
)

type Insert struct {
	Lin  int
	Col  int
	Id   string
	Exp1 interfaces.Expression
	Exp2 interfaces.Expression
}

func NewInsert(lin int, col int, id string, exp1, exp2 interfaces.Expression) Insert {
	instr := Insert{lin, col, id, exp1, exp2}
	return instr
}

func (p Insert) Ejecutar(ast *environment.AST, env interface{}) environment.Symbol {
	var result, tmpSymbol environment.Symbol
	var tmpArr *arrayList.List
	tmpSymbol = env.(environment.Environment).GetVariable(p.Id)
	tmpExp1 := p.Exp1.Ejecutar(ast, env) //indice
	tmpExp2 := p.Exp2.Ejecutar(ast, env) //valor
	tmpArr = arrayList.New()
	//validar que id sea vector
	if tmpSymbol.Tipo == environment.VECTOR {
		//validar tipo de expresion
		if ((tmpExp2.Tipo == tmpSymbol.TipoArr) || ((tmpExp2.Id == tmpSymbol.Id) && (tmpSymbol.Id != ""))) && (tmpExp1.Tipo == environment.INTEGER) {
			//validar indice
			if (tmpExp1.Valor.(int) <= tmpSymbol.Valor.(*arrayList.List).Len()) && (tmpExp1.Valor.(int) >= 0) {
				//recorrer indice
				for i := 0; i <= tmpSymbol.Valor.(*arrayList.List).Len(); i++ {
					//validar mutabilidad
					if tmpSymbol.Mutable {
						//asignar
						if tmpExp1.Valor.(int) == i {
							//setear capacidad
							if tmpSymbol.Valor.(*arrayList.List).Len() == tmpSymbol.Capacity {
								if tmpSymbol.Capacity == 0 {
									tmpSymbol.Capacity = tmpSymbol.Capacity + 1
								} else {
									tmpSymbol.Capacity = tmpSymbol.Capacity * 2
								}
							}
							inStackVal := tmpSymbol.Valor.(*arrayList.List).GetValue(i)
							if tmpExp2.Valor != nil {
								tmpArr.Add(tmpExp2)
							}
							if inStackVal != nil {
								tmpArr.Add(inStackVal)
							}
						} else {
							tmpArr.Add(tmpSymbol.Valor.(*arrayList.List).GetValue(i))
						}
					} else {
						fmt.Println("La variable no es mutable")
						return result
					}
				}
			} else {
				fmt.Println("Indice fuera del rango")
				return result
			}
		} else {
			fmt.Println("El tipo de dato es incorrecto")
			return result
		}
	} else {
		fmt.Println("Solo se puede hacer Insert sobre vectores")
		return result
	}
	//remplazar lista
	tmpSymbol.Valor = tmpArr
	//setear variable
	env.(environment.Environment).SetVariable(p.Id, tmpSymbol)

	return result
}
