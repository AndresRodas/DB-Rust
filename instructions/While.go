package instructions

import (
	"OLC2/environment"
	"OLC2/interfaces"
	"fmt"
	arrayList "github.com/colegno/arraylist"
	"strings"
)

type While struct {
	Lin       int
	Col       int
	Expresion interfaces.Expression
	Bloque    *arrayList.List
}

func NewWhile(lin int, col int, condition interfaces.Expression, bloque *arrayList.List) While {
	whileInstr := While{lin, col, condition, bloque}
	return whileInstr
}

func (p While) Ejecutar(ast *environment.AST, env interface{}) environment.Symbol {
	safeCont := 0
	breakFlag := false
	var result environment.Symbol

	for {
		safeCont++
		result = p.Expresion.Ejecutar(ast, env)

		if result.Valor == true {

			var whileEnv environment.Environment
			whileEnv = environment.NewEnvironment(env.(environment.Environment), "WHILE")

			for _, s := range p.Bloque.ToArray() {
				if strings.Contains(fmt.Sprintf("%T", s), "instructions") {
					result = s.(interfaces.Instruction).Ejecutar(ast, whileEnv) //BREAK, CONTINUE & RETURN
					if result.BreakFlag {
						result.BreakFlag = false
						breakFlag = true
						if result.Valor != nil {
							fmt.Println("No se permite retornar valor mediante un Break en un While")
						}
						break
					} else if result.ContinueFlag {
						result.ContinueFlag = false
						break
					} else if result.ReturnFlag {
						result.ReturnFlag = false
						return result
					}

				} else if strings.Contains(fmt.Sprintf("%T", s), "expressions") {
					result = s.(interfaces.Expression).Ejecutar(ast, whileEnv)
					if result.BreakFlag {
						result.BreakFlag = false
						breakFlag = true
						if result.Valor != nil {
							fmt.Println("No se permite retornar valor mediante un Break en un While")
						}
						break
					} else if result.ContinueFlag {
						result.ContinueFlag = false
						break
					} else if result.ReturnFlag {
						result.ReturnFlag = false
						return result
					}
				} else {
					fmt.Println("error en bloque del while")
					return result
				}
			}
			if breakFlag == true {
				break
			}
		} else {
			break
		}
		if safeCont >= 2500 {
			fmt.Println("StackOverflowError: se ha excedido el máximo de ciclos permitidos")
			break
		}
	}
	return result
}
