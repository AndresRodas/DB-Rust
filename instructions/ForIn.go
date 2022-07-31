package instructions

import (
	"OLC2/environment"
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
	//validar si es tipo array
	if p.Exp.Ejecutar(ast, env).Tipo == environment.ARRAY || p.Exp.Ejecutar(ast, env).Tipo == environment.VECTOR {
		//se extrae la lista
		arr = p.Exp.Ejecutar(ast, env).Valor.(*arrayList.List)
		//se recorre
		for _, s := range arr.ToArray() {
			//crendo entorno
			var loopEnv environment.Environment
			loopEnv = environment.NewEnvironment(env.(environment.Environment), "FOR-IN")
			//agregando variable al entorno
			//newSym := environment.Symbol{Lin: p.Lin, Col: p.Col, Id: "", Tipo: environment.INTEGER, Valor: s, Mutable: true}

			loopEnv.SaveVariable(p.Id, s.(environment.Symbol))
			//recorriendo bloque de instrucciones
			for _, b := range p.Inst.ToArray() {
				result = b.(interfaces.Instruction).Ejecutar(ast, loopEnv)
				if result.BreakFlag {
					breakFlag = true
					break
				}
				if result.ContinueFlag {
					break
				}
				if result.ReturnFlag {
					result.ReturnFlag = false
					return result
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
