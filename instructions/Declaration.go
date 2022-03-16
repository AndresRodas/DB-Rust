package instructions

import (
	"OLC2/environment"
	"OLC2/interfaces"
	"fmt"
)

type Declaration struct {
	Lin       int
	Col       int
	Id        string
	Tipo      environment.TipoExpresion
	Expresion interfaces.Expression
	Mutable   bool
}

func NewDeclaration(lin int, col int, id string, tipo environment.TipoExpresion, val interfaces.Expression, mut bool) Declaration {
	instr := Declaration{lin, col, id, tipo, val, mut}
	return instr
}

func (p Declaration) Ejecutar(ast *environment.AST, env interface{}) environment.Symbol {

	var result environment.Symbol

	result = p.Expresion.Ejecutar(ast, env)
	result.Mutable = p.Mutable

	if result.Tipo == p.Tipo {
		env.(environment.Environment).SaveVariable(p.Id, result, p.Tipo, p.Mutable)
	} else if p.Tipo == environment.WILDCARD { //aqui no se define tipo
		env.(environment.Environment).SaveVariable(p.Id, result, result.Tipo, p.Mutable)
	} else {
		fmt.Println("Los tipos no coinciden")
	}
	return result
}
