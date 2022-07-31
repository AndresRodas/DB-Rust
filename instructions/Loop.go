package instructions

import (
	"OLC2/environment"
	"OLC2/interfaces"
	"fmt"
	arrayList "github.com/colegno/arraylist"
	"strings"
)

type Loop struct {
	Lin    int
	Col    int
	Bloque *arrayList.List
}

func NewLoop(lin int, col int, bloque *arrayList.List) Loop {
	loopInstr := Loop{lin, col, bloque}
	return loopInstr
}

func (p Loop) Ejecutar(ast *environment.AST, env interface{}) environment.Symbol {
	safeCont := 0
	breakFlag := false
	var result environment.Symbol

	for {
		safeCont++
		var loopEnv environment.Environment
		loopEnv = environment.NewEnvironment(env.(environment.Environment), "LOOP")

		for _, s := range p.Bloque.ToArray() {
			if strings.Contains(fmt.Sprintf("%T", s), "instructions") {
				result = s.(interfaces.Instruction).Ejecutar(ast, loopEnv)
				//BREAK, CONTINUE & RETURN
				if result.BreakFlag {
					breakFlag = true
					result.BreakFlag = false
					break
				} else if result.ContinueFlag {
					result.ContinueFlag = false
					break
				} else if result.ReturnFlag {
					result.ReturnFlag = false
					return result
				}
			} else if strings.Contains(fmt.Sprintf("%T", s), "expressions") {
				result = s.(interfaces.Expression).Ejecutar(ast, loopEnv)
				//BREAK, CONTINUE & RETURN
				if result.BreakFlag {
					breakFlag = true
					result.BreakFlag = false
					break
				} else if result.ContinueFlag {
					result.ContinueFlag = false
					break
				} else if result.ReturnFlag {
					result.ReturnFlag = false
					return result
				}
			} else {
				fmt.Println("error en bloque")
			}
		}
		if breakFlag == true {
			break
		}
		if safeCont >= 150 {
			fmt.Println("StackOverflowError: se ha excedido el máximo de ciclos permitidos")
			break
		}
	}
	return result
}
