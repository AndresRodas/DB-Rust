package instructions

import (
	"OLC2/environment"
	"OLC2/expressions"
	"OLC2/interfaces"
	"fmt"
	arrayList "github.com/colegno/arraylist"
)

type ForIn struct {
	Lin  int
	Col  int
	Id   string
	Exp  interfaces.Expression
	Inst *arrayList.List
}

func NewForIn(lin int, col int, id string, exp interfaces.Expression, ins *arrayList.List) ForIn {
	loopInstr := ForIn{lin, col, id, exp, ins}
	return loopInstr
}

func (p ForIn) Ejecutar(ast *environment.AST, env interface{}) environment.Symbol {
	breakFlag := false
	var result environment.Symbol
	var arr *arrayList.List

	if p.Exp.Ejecutar(ast, env).Tipo == environment.ARRAY {

		arr = p.Exp.Ejecutar(ast, env).Valor.(*arrayList.List)
		for _, s := range arr.ToArray() {
			var loopEnv environment.Environment
			loopEnv = environment.NewEnvironment(env.(environment.Environment), "FOR-IN")
			result = environment.Symbol{
				Lin:     p.Lin,
				Col:     p.Col,
				Id:      p.Id,
				Tipo:    p.Exp.(expressions.Range).Tipo,
				Valor:   s,
				Mutable: false,
			}
			loopEnv.SaveVariable(p.Id, result, environment.INTEGER, false)
			for _, b := range p.Inst.ToArray() {
				result = b.(interfaces.Instruction).Ejecutar(ast, loopEnv)
				if result.Tipo == environment.BREAK {
					breakFlag = true
					break
				}
				if result.Tipo == environment.CONTINUE {
					break
				}
			}
			if breakFlag == true {
				break
			}
		}
	} else {
		fmt.Println("El tipo de expresion tiene que ser arreglo o vector")
	}

	return result
}
