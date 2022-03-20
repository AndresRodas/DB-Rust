package instructions

import (
	"OLC2/environment"
)

type ParamsDeclaration struct {
	Lin  int
	Col  int
	Id   string
	Tipo environment.TipoExpresion
}

func NewParamsDeclaration(lin int, col int, id string, tipo environment.TipoExpresion) ParamsDeclaration {
	instr := ParamsDeclaration{lin, col, id, tipo}
	return instr
}

func (p ParamsDeclaration) Ejecutar(ast *environment.AST, env interface{}) environment.Symbol {

	var result environment.Symbol

	result = environment.Symbol{
		Lin:     p.Lin,
		Col:     p.Col,
		Id:      p.Id,
		Tipo:    p.Tipo,
		Valor:   0,
		Mutable: true,
	}

	env.(environment.Environment).SaveVariable(p.Id, result, p.Tipo, true)

	return result
}
