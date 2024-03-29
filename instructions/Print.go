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
	valuesCopy := p.Values.Clone()
	if valuesCopy.Len() > 1 {
		//Con formato
		format := fmt.Sprintf("%v", valuesCopy.GetValue(0).(interfaces.Expression).Ejecutar(ast, env).Valor)
		valuesCopy.RemoveAtIndex(0)                // removemos el formato
		for _, val := range valuesCopy.ToArray() { //recorremos el resto
			valToPrint := val.(interfaces.Expression).Ejecutar(ast, env)
			if (valToPrint.Tipo == environment.ARRAY) || (valToPrint.Tipo == environment.VECTOR) {
				format = strings.Replace(format, "{:?}", p.ArrayToString(valToPrint.Valor.(*arrayList.List)), 1)
			} else if valToPrint.Tipo == environment.STRUCT { //imprimir struct
				format = strings.Replace(format, "{}", fmt.Sprintf("%v", p.StructToString(valToPrint)), 1)
			} else {
				format = strings.Replace(format, "{}", fmt.Sprintf("%v", valToPrint.Valor), 1)
			}
		}
		ToPrint = ToPrint + format
		ast.SetPrint(ToPrint + "\n")
		return result
	}
	//Sin formato
	for _, val := range p.Values.ToArray() {
		valToPrint := val.(interfaces.Expression).Ejecutar(ast, env)
		if valToPrint.Tipo == environment.STRUCT { //imprimir struct
			ToPrint = ToPrint + p.StructToString(valToPrint)
		} else {
			ToPrint = ToPrint + fmt.Sprintf("%v", valToPrint.Valor)
		}
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

func (p Print) StructToString(structSymbol environment.Symbol) string {
	tmpStructString := "{ "
	for key, element := range structSymbol.Valor.(map[string]environment.Symbol) {
		tmpStructString = tmpStructString + key + ": "
		if element.Tipo == environment.STRUCT {
			tmpStructString = tmpStructString + p.StructToString(element)
		} else {
			tmpStructString = tmpStructString + fmt.Sprintf("%v", element.Valor) + ", "
		}
	}
	tmpStructString = strings.TrimRight(tmpStructString, ", ")
	return tmpStructString + " }"
}
