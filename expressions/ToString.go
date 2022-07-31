package expressions

import (
	"OLC2/environment"
	"OLC2/interfaces"
	"fmt"
)

type ToString struct {
	Lin int
	Col int
	Exp interface{}
}

func NewToString(lin int, col int, ex interface{}) ToString {
	expr := ToString{lin, col, ex}
	return expr
}

func (p ToString) Ejecutar(ast *environment.AST, env interface{}) environment.Symbol {
	var result environment.Symbol

	tmpExp := p.Exp.(interfaces.Expression).Ejecutar(ast, env)

	newVal := fmt.Sprintf("%v", tmpExp.Valor)

	result = environment.Symbol{
		Lin:     p.Lin,
		Col:     p.Col,
		Id:      "",
		Tipo:    environment.STRING,
		Valor:   newVal,
		Mutable: true,
	}

	return result
}
