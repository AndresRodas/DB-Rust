package environment

import (
	arrayList "github.com/colegno/arraylist"
)

type AST struct {
	Main         *arrayList.List
	Instructions *arrayList.List
	Print        string
	Errors       string
}

func NewAST(main *arrayList.List, inst *arrayList.List, print string) AST {
	ast := AST{Main: main, Instructions: inst, Print: print}
	return ast
}

func (a *AST) GetPrint() string {
	return a.Print
}

func (a *AST) SetPrint(ToPrint string) {
	a.Print = a.Print + ToPrint
}

//func (a *AST) AddSymbol(id, symtype, datatype, envi string, row, column int) {
//	a.Symbols.Add("{ ID: '" + id + "', SymType: '" + symtype + "', DataType: '" + datatype + "', Environment: '" + envi + "', Row: '" + string(row) + "', Column: '" + string(column) + " }")
//}
