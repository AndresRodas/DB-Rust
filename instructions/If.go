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
	ElseIf    *arrayList.List
	ElseInst  *arrayList.List
}

func NewIf(lin int, col int, condition interfaces.Expression, bloque *arrayList.List, elif *arrayList.List, elinst *arrayList.List) If {
	ifInstr := If{lin, col, condition, bloque, elif, elinst}
	return ifInstr
}

func (p If) Ejecutar(ast *environment.AST, env interface{}) environment.Symbol {

	var condicion, result environment.Symbol
	condicion = p.Expresion.Ejecutar(ast, env)
	if condicion.Valor == true {
		//EJECUTANDO EL IF
		var ifEnv environment.Environment
		ifEnv = environment.NewEnvironment(env.(environment.Environment), "IF")
		for _, s := range p.Bloque.ToArray() {
			if strings.Contains(fmt.Sprintf("%T", s), "instructions") {
				result = s.(interfaces.Instruction).Ejecutar(ast, ifEnv)
				if result.Tipo == environment.BREAK || result.Tipo == environment.CONTINUE { //BREAK & CONTINUE
					return result
				}
			} else if strings.Contains(fmt.Sprintf("%T", s), "expressions") {
				result = s.(interfaces.Expression).Ejecutar(ast, ifEnv)
				if result.Tipo == environment.BREAK || result.Tipo == environment.CONTINUE { //BREAK & CONTINUE
					return result
				}
			} else {
				fmt.Println("error en bloque")
			}
		}
		return result
	} else if p.ElseIf.Len() > 0 {
		for _, s := range p.ElseIf.ToArray() {
			var condicion_elif environment.Symbol
			condicion_elif = s.(If).Expresion.Ejecutar(ast, env)
			if condicion_elif.Valor == true {
				//EJECUTANDO EL ELSE IF
				var elifEnv environment.Environment
				elifEnv = environment.NewEnvironment(env.(environment.Environment), "ELSE IF")
				for _, ns := range s.(If).Bloque.ToArray() {
					if strings.Contains(fmt.Sprintf("%T", ns), "instructions") {
						result = ns.(interfaces.Instruction).Ejecutar(ast, elifEnv)
						if result.Tipo == environment.BREAK || result.Tipo == environment.CONTINUE { //BREAK & CONTINUE
							return result
						}
					} else if strings.Contains(fmt.Sprintf("%T", ns), "expressions") {
						result = ns.(interfaces.Expression).Ejecutar(ast, elifEnv)
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
		if p.ElseInst.Len() > 0 {
			//EJECUTANDO ELSE
			var elseEnv environment.Environment
			elseEnv = environment.NewEnvironment(env.(environment.Environment), "ELSE")
			for _, ns := range p.ElseInst.ToArray() {
				if strings.Contains(fmt.Sprintf("%T", ns), "instructions") {
					result = ns.(interfaces.Instruction).Ejecutar(ast, elseEnv)
					if result.Tipo == environment.BREAK || result.Tipo == environment.CONTINUE { //BREAK & CONTINUE
						return result
					}
				} else if strings.Contains(fmt.Sprintf("%T", ns), "expressions") {
					result = ns.(interfaces.Expression).Ejecutar(ast, elseEnv)
					if result.Tipo == environment.BREAK || result.Tipo == environment.CONTINUE { //BREAK & CONTINUE
						return result
					}
				} else {
					fmt.Println("error en bloque")
				}
			}
			return result
		}

	} else if p.ElseInst.Len() > 0 {
		//EJECUTANDO ELSE
		var elseEnv environment.Environment
		elseEnv = environment.NewEnvironment(env.(environment.Environment), "ELSE")
		for _, ns := range p.ElseInst.ToArray() {
			if strings.Contains(fmt.Sprintf("%T", ns), "instructions") {
				result = ns.(interfaces.Instruction).Ejecutar(ast, elseEnv)
				if result.Tipo == environment.BREAK || result.Tipo == environment.CONTINUE { //BREAK & CONTINUE
					return result
				}
			} else if strings.Contains(fmt.Sprintf("%T", ns), "expressions") {
				result = ns.(interfaces.Expression).Ejecutar(ast, elseEnv)
				if result.Tipo == environment.BREAK || result.Tipo == environment.CONTINUE { //BREAK & CONTINUE
					return result
				}
			} else {
				fmt.Println("error en bloque")
			}
		}
		return result
	}

	return result
}
