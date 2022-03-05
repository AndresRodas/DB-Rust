package instructions

import (
	"OLC2/environment"
	"OLC2/interfaces"
	"fmt"
	arrayList "github.com/colegno/arraylist"
)

type Print struct {
	Lin    int
	Col    int
	Values *arrayList.List
}

func (p Print) Ejecutar(ast *environment.AST, env interface{}) interface{} {

	var ToPrint string
	for _, p := range p.Values.ToArray() {
		//fmt.Println(p.(interfaces.Expression).Ejecutar(ast, env))
		ToPrint = ToPrint + fmt.Sprintf("%v", p.(interfaces.Expression).Ejecutar(ast, env).Valor)
	}
	ast.SetPrint(ToPrint + "\n")
	return nil
}

func NewPrint(lin int, col int, val *arrayList.List) Print {
	return Print{lin, col, val}
}
