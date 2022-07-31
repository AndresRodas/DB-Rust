package expressions

import (
	"OLC2/environment"
	"OLC2/interfaces"
	"fmt"
	arrayList "github.com/colegno/arraylist"
)

type StructExp struct {
	Lin     int
	Col     int
	Id      string
	ListExp *arrayList.List
}

func NewStructExp(lin int, col int, id string, list *arrayList.List) StructExp {
	exp := StructExp{lin, col, id, list}
	return exp
}

func (p StructExp) Ejecutar(ast *environment.AST, env interface{}) environment.Symbol {
	var valor map[string]environment.Symbol
	var result environment.Symbol
	var tempExp *arrayList.List
	tempExp = arrayList.New()
	//se guarda el listado de valores nuevos
	for _, s := range p.ListExp.ToArray() {
		tempExp.Add(s)
	}
	//se obtiene la estructura guardada
	resultStruct := env.(environment.Environment).GetStruct(p.Id)

	//se valida si existe struct
	if resultStruct.Tipo == environment.STRUCT {
		//Validar tamaño y que cada tipo coincida con el struct existente
		if resultStruct.Valor.(*arrayList.List).Len() == p.ListExp.Len() {
			valor = make(map[string]environment.Symbol)
			//contador de definiciones
			var contDef = 0
			//recorrer el struct almacenado
			for _, strAlm := range resultStruct.Valor.(*arrayList.List).ToArray() { //StructType
				//recorrer los valores de entrada
				for _, strEnt := range p.ListExp.ToArray() { // StructContent

					//validar las coincidencias y los tipos
					if strAlm.(environment.StructType).Id == strEnt.(environment.StructContent).Id {
						tempVal := strEnt.(environment.StructContent).Exp.(interfaces.Expression).Ejecutar(ast, env)
						if strAlm.(environment.StructType).Tipo == tempVal.Tipo {
							contDef++
							valor[strAlm.(environment.StructType).Id] = tempVal
							break
						} else if strAlm.(environment.StructType).Tipo == environment.WILDCARD {
							//si es tipo struct (trae id en vez de tipo)
							if strAlm.(environment.StructType).StructId == tempVal.Id {
								contDef++
								valor[strAlm.(environment.StructType).Id] = tempVal
								break
							} else {
								fmt.Println("El tipo de dato del struct no coincide")
								return result
							}
						} else {
							fmt.Println("El tipo de dato del struct no coincide")
							return result
						}
					}
				}
			}
			//si es valido, guardarlo
			if p.ListExp.Len() == contDef {
				//fmt.Println("ASIGNANDO EL EXTRA TIPO: ", p.Id)
				result = environment.Symbol{Lin: p.Lin, Col: p.Col, Id: p.Id, Tipo: environment.STRUCT, Valor: valor, Mutable: true, ExtraTipo: p.Id}
			} else {
				fmt.Println("Faltan parámetros en el struct")
			}
		} else {
			fmt.Println("La cantidad de valores en el struct es incorrecta")
		}
	}
	return result
}
