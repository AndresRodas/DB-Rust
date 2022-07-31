package instructions

import (
	"OLC2/environment"
	arrayList "github.com/colegno/arraylist"
)

type Struct struct {
	Lin     int
	Col     int
	Id      string
	ListAtr *arrayList.List
}

func NewStruct(lin int, col int, id string, list *arrayList.List) Struct {
	instr := Struct{lin, col, id, list}
	return instr
}

func (p Struct) Ejecutar(ast *environment.AST, env interface{}) environment.Symbol {
	env.(environment.Environment).SaveStruct(p.Id, p.ListAtr, true)
	return environment.Symbol{Lin: p.Lin, Col: p.Col, Id: "", Tipo: environment.STRUCT, Valor: p.ListAtr, Mutable: true}
}
