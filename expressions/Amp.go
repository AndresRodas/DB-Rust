package expressions

import (
	"OLC2/environment"
	"OLC2/interfaces"
	"fmt"
)

type Amp struct {
	Lin int
	Col int
	Exp interface{}
}

func NewAmp(lin int, col int, ex interface{}) Amp {
	expr := Amp{lin, col, ex}
	return expr
}

func (p Amp) Ejecutar(ast *environment.AST, env interface{}) environment.Symbol {
	var result environment.Symbol

	tmpExp := p.Exp.(interfaces.Expression).Ejecutar(ast, env)

	newVal := fmt.Sprintf("%v", tmpExp.Valor)

	result = environment.Symbol{
		Lin:     p.Lin,
		Col:     p.Col,
		Id:      "",
		Tipo:    environment.STR,
		Valor:   newVal,
		Mutable: true,
	}

	return result
}
