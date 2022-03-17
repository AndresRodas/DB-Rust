package instructions

import (
	"OLC2/environment"
	"OLC2/interfaces"
	"fmt"
)

type Break struct {
	Lin       int
	Col       int
	Expresion interfaces.Expression
}

func NewBreak(lin int, col int, exp interfaces.Expression) Break {
	breakInstr := Break{lin, col, exp}
	return breakInstr
}

func (p Break) Ejecutar(ast *environment.AST, env interface{}) environment.Symbol {
	var result environment.Symbol
	//validar ciclo entorno
	if env.(environment.Environment).LoopValidation() {
		if p.Expresion == nil { //BREAK SIN VALOR
			result = environment.Symbol{
				Lin:     p.Lin,
				Col:     p.Col,
				Id:      "BREAK",
				Tipo:    environment.BREAK,
				Valor:   nil,
				Mutable: false,
			}
			return result
		} else { //BREAK CON VALOR
			result = environment.Symbol{
				Lin:     p.Lin,
				Col:     p.Col,
				Id:      "BREAK",
				Tipo:    environment.BREAK,
				Valor:   p.Expresion.Ejecutar(ast, env).Valor,
				Mutable: false,
			}
		}
	} else {
		fmt.Println("Error sentencia break")
	}
	return result
}
