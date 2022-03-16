package instructions

import (
	"OLC2/environment"
	"OLC2/interfaces"
	"fmt"
	arrayList "github.com/colegno/arraylist"
	"strings"
)

type If struct {
	Lin       int
	Col       int
	Expresion interfaces.Expression
	Bloque    *arrayList.List
}

func NewIf(lin int, col int, condition interfaces.Expression, bloque *arrayList.List) If {
	ifInstr := If{lin, col, condition, bloque}
	return ifInstr
}

func (p If) Ejecutar(ast *environment.AST, env interface{}) environment.Symbol {

	var condicion, result environment.Symbol
	condicion = p.Expresion.Ejecutar(ast, env)
	if condicion.Valor == true {
		var tmpEnv environment.Environment
		tmpEnv = environment.NewEnvironment(env.(environment.Environment), "IF")
		for _, s := range p.Bloque.ToArray() {
			if strings.Contains(fmt.Sprintf("%T", s), "instructions") {
				result = s.(interfaces.Instruction).Ejecutar(ast, tmpEnv)
			} else if strings.Contains(fmt.Sprintf("%T", s), "expressions") {
				result = s.(interfaces.Expression).Ejecutar(ast, tmpEnv)
			} else {
				fmt.Println("error en bloque")
			}
		}
	}
	return result
}
