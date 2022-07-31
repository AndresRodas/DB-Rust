package instructions

import (
	"OLC2/environment"
	"OLC2/interfaces"
	"fmt"
	arrayList "github.com/colegno/arraylist"
)

type ArrayAssign struct {
	Lin     int
	Col     int
	Id      string
	ListExp *arrayList.List
	Exp     interfaces.Expression
}

func NewArrayAssign(lin int, col int, iden string, list *arrayList.List, expr interfaces.Expression) ArrayAssign {
	instr := ArrayAssign{lin, col, iden, list, expr}
	return instr
}

func (p ArrayAssign) Ejecutar(ast *environment.AST, env interface{}) environment.Symbol {
	var result, newVal environment.Symbol
	var expList *arrayList.List

	newVal = p.Exp.Ejecutar(ast, env)
	result = env.(environment.Environment).GetVariable(p.Id) //obtener array
	expList = arrayList.New()
	for _, exp := range p.ListExp.ToArray() {
		tempIndex := exp.(interfaces.Expression).Ejecutar(ast, env)
		if tempIndex.Tipo != environment.INTEGER {
			fmt.Println("Error indice incorrecto")
			return result
		}
		expList.Add(tempIndex.Valor) //agregando valores a lista
	}

	result = p.ReplaceArray(result, expList, newVal)        //asignando nuevo valor a index
	env.(environment.Environment).SetVariable(p.Id, result) //seteando arreglo final
	return result

}

func (p ArrayAssign) ReplaceArray(array environment.Symbol, indexList *arrayList.List, newVal environment.Symbol) environment.Symbol {
	var result environment.Symbol
	var tempList *arrayList.List
	tempList = arrayList.New()
	cloneIndex := indexList.Clone()                               //clonando lista
	index := indexList.GetValue(0)                                //extrayendo indice
	contIndex := 0                                                //contador
	for _, val := range array.Valor.(*arrayList.List).ToArray() { //recorrer el array de simbolos
		if contIndex == index { //comparando indices
			if (val.(environment.Symbol).Tipo == environment.ARRAY) && (cloneIndex.Len() > 1) { //es array? y hay mas indices?
				cloneIndex.RemoveAtIndex(0)                                                //eliminando indice matcheado
				tempList.Add(p.ReplaceArray(val.(environment.Symbol), cloneIndex, newVal)) //recursividad
			} else {
				if val.(environment.Symbol).Tipo == newVal.Tipo {
					tempList.Add(newVal) //hace match con no lista y lo remplaza
				} else {
					fmt.Println("Error los tipos no coinciden") //error
					tempList.Add(val)
				}
			}
		} else {
			tempList.Add(val) //no hace match y agrega
		}
		contIndex++ //suma contador indice
	}
	result = environment.Symbol{Lin: array.Lin, Col: array.Col, Id: array.Id, Tipo: array.Tipo, Valor: tempList, Mutable: array.Mutable}
	return result
}
