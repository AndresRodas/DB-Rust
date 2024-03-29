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
	FuncId  string
}

func NewFunction(lin int, col int, id string, listd *arrayList.List, tipo environment.TipoExpresion, bloc *arrayList.List, fid string) Function {
	instr := Function{lin, col, id, listd, tipo, bloc, fid}
	return instr
}

func (p Function) Ejecutar(ast *environment.AST, env interface{}) environment.Symbol {
	var result environment.Symbol
	var function environment.FunctionSymbol

	//creando simbolo funcion
	function = environment.FunctionSymbol{
		Lin:         p.Lin,
		Col:         p.Col,
		Id:          p.Id,
		Ent:         nil,
		ListDec:     p.ListDec,
		Block:       p.Bloque,
		TipoRetorno: p.Tipo,
		IdTipo:      p.FuncId,
	}
	//fmt.Println("ESTA ES LA FUNCION ", p.Id)
	//fmt.Println(p.ListDec.ToArray())
	//guardando simbolo funcion
	env.(environment.Environment).SaveFunction(p.Id, function)
	return result
}
