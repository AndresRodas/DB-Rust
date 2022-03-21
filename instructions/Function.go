package instructions

import (
	"OLC2/environment"
	arrayList "github.com/colegno/arraylist"
)

type Function struct {
	Lin     int
	Col     int
	Id      string
	ListDec *arrayList.List
	Tipo    environment.TipoExpresion
	Bloque  *arrayList.List
}

func NewFunction(lin int, col int, id string, listd *arrayList.List, tipo environment.TipoExpresion, bloc *arrayList.List) Function {
	instr := Function{lin, col, id, listd, tipo, bloc}
	return instr
}

func (p Function) Ejecutar(ast *environment.AST, env interface{}) environment.Symbol {
	var result environment.Symbol
	var function environment.FunctionSymbol

	//creando nuevo entorno de funcion
	var funcEnv environment.Environment
	funcEnv = environment.NewEnvironment(env, "FUNCTION-"+p.Id)

	//creando simbolo funcion
	function = environment.FunctionSymbol{
		Lin:     p.Lin,
		Col:     p.Col,
		Id:      p.Id,
		Ent:     funcEnv,
		ListDec: p.ListDec,
		Block:   p.Bloque,
	}
	//guardando simbolo funcion
	env.(environment.Environment).SaveFunction(p.Id, function)
	return result
}
