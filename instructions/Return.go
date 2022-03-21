package instructions

import (
	"OLC2/environment"
	"OLC2/interfaces"
	"fmt"
)

type Return struct {
	Lin       int
	Col       int
	Expresion interfaces.Expression
}

func NewReturn(lin int, col int, exp interfaces.Expression) Return {
	breakInstr := Return{lin, col, exp}
	return breakInstr
}

func (p Return) Ejecutar(ast *environment.AST, env interface{}) environment.Symbol {
	var result environment.Symbol
	//validar que esté dentro de funcion
	if env.(environment.Environment).FuncValidation() {
		if p.Expresion == nil { //RETURN SIN VALOR
			result = environment.Symbol{
				Lin:     p.Lin,
				Col:     p.Col,
				Id:      "RETURN",
				Tipo:    environment.NULL,
				Valor:   nil,
				Mutable: true,
			}
			return result
		} else { //RETURN CON VALOR
			tmpVal := p.Expresion.Ejecutar(ast, env)
			result = environment.Symbol{
				Lin:     p.Lin,
				Col:     p.Col,
				Id:      "RETURN",
				Tipo:    tmpVal.Tipo,
				Valor:   tmpVal.Valor,
				Mutable: true,
			}
		}
	} else {
		fmt.Println("Error sentencia return")
	}
	return result
}
