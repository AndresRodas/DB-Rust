package expressions

import (
	"OLC2/environment"
	"OLC2/interfaces"
	arrayList "github.com/colegno/arraylist"
)

type Array struct {
	Lin     int
	Col     int
	ListExp *arrayList.List
}

func NewArray(lin int, col int, list *arrayList.List) Array {
	exp := Array{lin, col, list}
	return exp
}

func (p Array) Ejecutar(ast *environment.AST, env interface{}) environment.Symbol {

	var tempExp *arrayList.List
	tempExp = arrayList.New()

	for _, s := range p.ListExp.ToArray() {
		tempExp.Add(s.(interfaces.Expression).Ejecutar(ast, env))
	}

	return environment.Symbol{
		Lin:     p.Lin,
		Col:     p.Col,
		Id:      "",
		Tipo:    environment.ARRAY,
		Valor:   tempExp,
		Mutable: true,
	}
}
