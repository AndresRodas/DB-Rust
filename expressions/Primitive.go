package expressions

import (
	"OLC2/environment"
)

type Primitive struct {
	Lin   int
	Col   int
	Valor interface{}
	Tipo  environment.TipoExpresion
}

func NewPrimitive(lin int, col int, valor interface{}, tipo environment.TipoExpresion) Primitive {
	exp := Primitive{lin, col, valor, tipo}
	return exp
}

func (p Primitive) Ejecutar(ast *environment.AST, env interface{}) environment.Symbol {

	return environment.Symbol{
		Lin:     p.Lin,
		Col:     p.Col,
		Id:      "",
		Tipo:    p.Tipo,
		Valor:   p.Valor,
		Mutable: true,
	}
}
