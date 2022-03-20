package instructions

import (
	"OLC2/environment"
	"OLC2/interfaces"
	"fmt"
	arrayList "github.com/colegno/arraylist"
	"strings"
)

type Print struct {
	Lin    int
	Col    int
	Values *arrayList.List
}

func NewPrint(lin int, col int, val *arrayList.List) Print {
	return Print{lin, col, val}
}

func (p Print) Ejecutar(ast *environment.AST, env interface{}) environment.Symbol {
	var result environment.Symbol
	var ToPrint string
	if p.Values.Len() > 1 {
		//Con formato
		format := fmt.Sprintf("%v", p.Values.GetValue(0).(interfaces.Expression).Ejecutar(ast, env).Valor)
		p.Values.RemoveAtIndex(0)                // removemos el formato
		for _, val := range p.Values.ToArray() { //recorremos el resto
			valToPrint := val.(interfaces.Expression).Ejecutar(ast, env)
			if (valToPrint.Tipo == environment.ARRAY) || (valToPrint.Tipo == environment.VECTOR) {
				format = strings.Replace(format, "{:?}", p.ArrayToString(valToPrint.Valor.(*arrayList.List)), 1)
			} else {
				format = strings.Replace(format, "{}", fmt.Sprintf("%v", valToPrint.Valor), 1)
			}
		}
		ToPrint = ToPrint + format
		ast.SetPrint(ToPrint + "\n")
		return result
	}
	//Sin formato
	for _, p := range p.Values.ToArray() {
		ToPrint = ToPrint + fmt.Sprintf("%v", p.(interfaces.Expression).Ejecutar(ast, env).Valor)
	}
	ast.SetPrint(ToPrint + "\n")
	return result
}

func (p Print) ArrayToString(array *arrayList.List) string {
	var strBuffer string
	strBuffer = "["
	contInd := 0
	for _, val := range array.ToArray() {
		contInd++
		if (val.(environment.Symbol).Tipo == environment.ARRAY) || (val.(environment.Symbol).Tipo == environment.VECTOR) {
			if contInd < array.Len() {
				strBuffer = strBuffer + p.ArrayToString(val.(environment.Symbol).Valor.(*arrayList.List)) + ","
			} else {
				strBuffer = strBuffer + p.ArrayToString(val.(environment.Symbol).Valor.(*arrayList.List))
			}

		} else {
			if contInd < array.Len() {
				strBuffer = strBuffer + fmt.Sprintf("%v", val.(environment.Symbol).Valor) + ", "
			} else {
				strBuffer = strBuffer + fmt.Sprintf("%v", val.(environment.Symbol).Valor)
			}

		}
	}
	return strBuffer + "]"

}
