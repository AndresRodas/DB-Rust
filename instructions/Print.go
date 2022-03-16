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

func (p Print) Ejecutar(ast *environment.AST, env interface{}) environment.Symbol {
	var result environment.Symbol
	var ToPrint string
	if p.Values.Len() > 1 {
		//Con formato
		format := fmt.Sprintf("%v", p.Values.GetValue(0).(interfaces.Expression).Ejecutar(ast, env).Valor)
		p.Values.RemoveAtIndex(0)
		for _, p := range p.Values.ToArray() {
			format = strings.Replace(format, "{}", fmt.Sprintf("%v", p.(interfaces.Expression).Ejecutar(ast, env).Valor), 1)
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

func NewPrint(lin int, col int, val *arrayList.List) Print {
	return Print{lin, col, val}
}
