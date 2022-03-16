package expressions

import (
	"OLC2/environment"
)

type CallVar struct {
	Lin int
	Col int
	Id  string
}

func (p CallVar) Ejecutar(ast *environment.AST, env interface{}) environment.Symbol {
	result := env.(environment.Environment).GetVariable(p.Id)
	return result
}

func NewCallVar(lin int, col int, id string) CallVar {
	exp := CallVar{lin, col, id}
	return exp
}
