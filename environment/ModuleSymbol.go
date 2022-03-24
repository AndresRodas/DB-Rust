package environment

import arrayList "github.com/colegno/arraylist"

type ModuleSymbol struct {
	Lin    int
	Col    int
	Id     string
	Childs *arrayList.List
}
