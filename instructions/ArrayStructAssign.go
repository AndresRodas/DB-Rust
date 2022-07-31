package instructions

import (
	"OLC2/environment"
	"OLC2/interfaces"
	"fmt"
	arrayList "github.com/colegno/arraylist"
)

type ArrayStructAssign struct {
	Lin     int
	Col     int
	Arr     string
	Index   interfaces.Expression
	ListStr *arrayList.List
	Exp     interfaces.Expression
}

func NewArrayStructAssign(lin int, col int, iarr string, ind interfaces.Expression, listst *arrayList.List, expr interfaces.Expression) ArrayStructAssign {
	instr := ArrayStructAssign{lin, col, iarr, ind, listst, expr}
	return instr
}

func (p ArrayStructAssign) Ejecutar(ast *environment.AST, env interface{}) environment.Symbol {
	var result, arrExp, indexExp environment.Symbol
	//lista temporal
	var tmpList *arrayList.List
	tmpList = arrayList.New()
	//obteniendo arreglo
	arrExp = env.(environment.Environment).GetVariable(p.Arr)
	//obteniendo indice
	indexExp = p.Index.(interfaces.Expression).Ejecutar(ast, env)
	//validando tipos
	if (arrExp.Tipo == environment.ARRAY || arrExp.Tipo == environment.VECTOR) && indexExp.Tipo == environment.INTEGER {
		//recorriendo arreglo array o vector
		for i := 0; i < arrExp.Valor.(*arrayList.List).Len(); i++ {
			//validando indice, cuando indice sea igual
			if i == indexExp.Valor.(int) {
				//creacion de diccionario temporal
				var tmpDic map[string]environment.Symbol
				tmpDic = make(map[string]environment.Symbol)
				//asignacion de diccionario (valor del array e index que acabamos de obtener)
				tmpDic = arrExp.Valor.(*arrayList.List).GetValue(i).(environment.Symbol).Valor.(map[string]environment.Symbol)
				//recorro la lista de id
				for _, s := range p.ListStr.ToArray() { //recorremos lista de ides (acceso struct)
					//validando variable
					if variable, ok := tmpDic[s.(string)]; ok {
						//validando tipo struct
						if variable.Tipo == environment.STRUCT {
							//validando mutabilidad
							if variable.Mutable == true {
								tmpDic = variable.Valor.(map[string]environment.Symbol)
							} else {
								fmt.Println("La variable no es mutable")
								break
							}
						} else if tmpDic[s.(string)].Mutable {
							tmpDic[s.(string)] = p.Exp.Ejecutar(ast, env) //asignacion de struct
							break
						} else {
							fmt.Println("La variable no es mutable")
							break
						}
					}

				}
				//agregando struct editado
				tmpList.Add(arrExp.Valor.(*arrayList.List).GetValue(i))
			} else {
				//de lo contrario agregar los demas indices
				tmpList.Add(arrExp.Valor.(*arrayList.List).GetValue(i))
			}
		}
	} else {
		fmt.Println("Tipo de dato erroneo")
		return result
	}
	//asignando el result
	result = arrExp
	result.Valor = tmpList
	//seteando variable final
	env.(environment.Environment).SetVariable(p.Arr, result)
	return result
}
