package instructions

import (
	"OLC2/environment"
	"OLC2/interfaces"
	"fmt"
	arrayList "github.com/colegno/arraylist"
)

type ArrayDeclaration struct {
	Lin       int
	Col       int
	Id        string
	Tipo      *arrayList.List
	Expresion interfaces.Expression
	Mutable   bool
}

func NewArrayDeclaration(lin int, col int, id string, tipo *arrayList.List, val interfaces.Expression, mut bool) ArrayDeclaration {
	instr := ArrayDeclaration{lin, col, id, tipo, val, mut}
	return instr
}

func (p ArrayDeclaration) Ejecutar(ast *environment.AST, env interface{}) environment.Symbol {
	//Traer simbolo
	var result environment.Symbol
	result = p.Expresion.Ejecutar(ast, env)
	//validar array
	if p.ValidarArray(ast, env, result, p.Tipo) {
		result.Id = p.Id
		env.(environment.Environment).SaveVariable(p.Id, result, environment.ARRAY, p.Mutable)
	} else {
		fmt.Println("Los tipos no coinciden")
	}
	return result
}

func (p ArrayDeclaration) ValidarArray(ast *environment.AST, env interface{}, arr1 environment.Symbol, tip *arrayList.List) bool {
	flag := true
	arrCont := 0
	arrType := tip.GetValue(tip.Len() - 1)
	arrSize := arrType.(environment.ArrayType).Size.(interfaces.Expression).Ejecutar(ast, env).Valor.(int)
	tip.RemoveAtIndex(tip.Len() - 1)
	for _, arr := range arr1.Valor.(*arrayList.List).ToArray() {
		if arr.(environment.Symbol).Tipo != arrType.(environment.ArrayType).Tipo {
			return false
		}
		if arr.(environment.Symbol).Tipo == environment.ARRAY {
			flag = p.ValidarArray(ast, env, arr.(environment.Symbol), tip.Clone())
			if !flag {
				return flag
			}
		}
		arrCont++
	}

	if arrCont == arrSize && flag {
		return true
	}
	return false
}
