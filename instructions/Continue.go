package instructions

import (
	"OLC2/environment"
	"fmt"
)

type Continue struct {
	Lin int
	Col int
}

func NewContinue(lin int, col int) Continue {
	conInstr := Continue{lin, col}
	return conInstr
}

func (p Continue) Ejecutar(ast *environment.AST, env interface{}) environment.Symbol {
	var result environment.Symbol
	//validar ciclo entorno
	if env.(environment.Environment).LoopValidation() {
		result = environment.Symbol{
			Lin:     p.Lin,
			Col:     p.Col,
			Id:      "CONTINUE",
			Tipo:    environment.NULL,
			Valor:   nil,
			Mutable: false,
		}
	} else {
		fmt.Println("Error sentencia continue")
	}
	return result
}
