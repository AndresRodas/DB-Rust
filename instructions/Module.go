package instructions

import (
	"OLC2/environment"
	arrayList "github.com/colegno/arraylist"
)

type Module struct {
	Lin    int
	Col    int
	Id     string
	Childs *arrayList.List
}

func NewModule(lin int, col int, id string, hijos *arrayList.List) Module {
	inst := Module{lin, col, id, hijos}
	return inst
}

func (p Module) Ejecutar(ast *environment.AST, env interface{}) environment.Symbol {
	var result environment.ModuleSymbol

	//creando simbolo modulo
	result = environment.ModuleSymbol{Lin: p.Lin, Col: p.Col, Id: p.Id, Childs: p.Childs}

	//guardando modulo
	env.(environment.Environment).SaveModule(p.Id, result)

	return environment.Symbol{}
}
