package instructions

import (
	"OLC2/environment"
	"OLC2/interfaces"
	arrayList "github.com/colegno/arraylist"
)

type StructAssign struct {
	Lin        int
	Col        int
	ListStruct *arrayList.List
	Exp        interfaces.Expression
}

func NewStructAssign(lin int, col int, list *arrayList.List, expr interfaces.Expression) StructAssign {
	instr := StructAssign{lin, col, list, expr}
	return instr
}

func (p StructAssign) Ejecutar(ast *environment.AST, env interface{}) environment.Symbol {
	var result environment.Symbol
	exp := p.Exp.Ejecutar(ast, env)
	result = env.(environment.Environment).SetStruct(p.ListStruct, exp)
	return result
}
