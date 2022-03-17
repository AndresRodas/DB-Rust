package instructions

import (
	"OLC2/environment"
	arrayList "github.com/colegno/arraylist"
)

type Arm struct {
	Lin     int
	Col     int
	Matches *arrayList.List
	Block   *arrayList.List
}

func NewArm(lin int, col int, match *arrayList.List, bloque *arrayList.List) Arm {
	brazo := Arm{lin, col, match, bloque}
	return brazo
}

func (p Arm) Ejecutar(ast *environment.AST, env interface{}) environment.Symbol {
	var result environment.Symbol
	//ToDo ejecutar
	return result
}
