package expressions

import (
	"OLC2/environment"
	"OLC2/interfaces"
)

type Clone struct {
	Lin int
	Col int
	Exp interface{}
}

func NewClone(lin int, col int, ex interface{}) Clone {
	expr := Clone{lin, col, ex}
	return expr
}

func (p Clone) Ejecutar(ast *environment.AST, env interface{}) environment.Symbol {
	var result environment.Symbol

	tmpExp := p.Exp.(interfaces.Expression).Ejecutar(ast, env)

	//generando copia
	result = environment.Symbol{
		Lin:       tmpExp.Lin,
		Col:       tmpExp.Col,
		Id:        tmpExp.Id,
		Tipo:      tmpExp.Tipo,
		Valor:     tmpExp.Valor,
		Mutable:   tmpExp.Mutable,
		Capacity:  tmpExp.Capacity,
		TipoArr:   tmpExp.TipoArr,
		ExtraTipo: tmpExp.ExtraTipo,
	}

	return result
}
