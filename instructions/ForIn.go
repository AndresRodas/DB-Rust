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
	if p.Exp.Ejecutar(ast, env).Tipo == environment.ARRAY {
		//se extrae la lista
		arr = p.Exp.Ejecutar(ast, env).Valor.(*arrayList.List)
		//se recorre
		for _, s := range arr.ToArray() {
			//crendo entorno
			var loopEnv environment.Environment
			loopEnv = environment.NewEnvironment(env.(environment.Environment), "FOR-IN")
			//agregando variable al entorno
			loopEnv.SaveVariable(p.Id, s.(environment.Symbol), environment.INTEGER, false)
			//recorriendo bloque de instrucciones
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