package environment

import arrayList "github.com/colegno/arraylist"

type FunctionSymbol struct {
	Lin         int
	Col         int
	Id          string
	Ent         interface{}
	ListDec     *arrayList.List
	Block       *arrayList.List
	TipoRetorno TipoExpresion
	IdTipo      string
}
