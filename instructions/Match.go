package instructions

import (
	"OLC2/environment"
	"OLC2/interfaces"
	"fmt"
	arrayList "github.com/colegno/arraylist"
	"strings"
)

type Match struct {
	Lin    int
	Col    int
	Exp    interfaces.Expression
	Arms   *arrayList.List
	DefArm *arrayList.List
}

func NewMatch(lin int, col int, exp interfaces.Expression, arms *arrayList.List, def *arrayList.List) Match {
	mtch := Match{lin, col, exp, arms, def}
	return mtch
}

func (p Match) Ejecutar(ast *environment.AST, env interface{}) environment.Symbol {

	var exp, result environment.Symbol
	exp = p.Exp.Ejecutar(ast, env)
	//creando entorno
	var matchEnv environment.Environment
	matchEnv = environment.NewEnvironment(env.(environment.Environment), "MATCH")
	//validando tipos de matches
	for _, arms := range p.Arms.ToArray() {
		for _, matchs := range arms.(Arm).Matches.ToArray() {
			if matchs.(interfaces.Expression).Ejecutar(ast, env).Tipo != exp.Tipo {
				fmt.Println("Los tipos de expresiones en los brazos deben ser iguales al valor match")
				return result
			}
		}
	}
	//validando si es bool
	if exp.Tipo == environment.BOOLEAN {
		var flagTrue, flagFalse = false, false
		for _, arms := range p.Arms.ToArray() {
			for _, matchs := range arms.(Arm).Matches.ToArray() {
				if matchs.(interfaces.Expression).Ejecutar(ast, env).Valor == true {
					flagTrue = true
				} else if matchs.(interfaces.Expression).Ejecutar(ast, env).Valor == false {
					flagFalse = true
				}
			}
		}
		if (flagTrue == false && p.DefArm.Len() == 0) || (flagFalse == false && p.DefArm.Len() == 0) {
			fmt.Println("Todas las coincidencias no han sido cubiertas")
			return result
		}
	}
	//recorrindo brazos
	for _, arms := range p.Arms.ToArray() {

		//recorriendo matches
		for _, matchs := range arms.(Arm).Matches.ToArray() {
			//si un elemento hace match
			if matchs.(interfaces.Expression).Ejecutar(ast, env).Valor == exp.Valor {
				for _, s := range arms.(Arm).Block.ToArray() {
					if strings.Contains(fmt.Sprintf("%T", s), "instructions") {
						result = s.(interfaces.Instruction).Ejecutar(ast, matchEnv)
						if result.Tipo == environment.BREAK || result.Tipo == environment.CONTINUE { //BREAK & CONTINUE
							return result
						}
					} else if strings.Contains(fmt.Sprintf("%T", s), "expressions") {
						result = s.(interfaces.Expression).Ejecutar(ast, matchEnv)
						if result.Tipo == environment.BREAK || result.Tipo == environment.CONTINUE { //BREAK & CONTINUE
							return result
						}
					} else {
						fmt.Println("error en bloque")
					}
				}
				return result
			}
		}
	}
	//ejecutando brazo default
	for _, d := range p.DefArm.ToArray() {
		if strings.Contains(fmt.Sprintf("%T", d), "instructions") {
			result = d.(interfaces.Instruction).Ejecutar(ast, matchEnv)
			if result.Tipo == environment.BREAK || result.Tipo == environment.CONTINUE { //BREAK & CONTINUE
				return result
			}
		} else if strings.Contains(fmt.Sprintf("%T", d), "expressions") {
			result = d.(interfaces.Expression).Ejecutar(ast, matchEnv)
			if result.Tipo == environment.BREAK || result.Tipo == environment.CONTINUE { //BREAK & CONTINUE
				return result
			}
		} else {
			fmt.Println("error en bloque")
			return result
		}
	}
	return result
}
