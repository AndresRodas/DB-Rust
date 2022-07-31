package expressions

import (
	"OLC2/environment"
	"OLC2/interfaces"
	"fmt"
)

type StructAccess struct {
	Lin    int
	Col    int
	Struct interfaces.Expression
	Id     string
}

func NewStructAccess(lin int, col int, str interfaces.Expression, id string) StructAccess {
	exp := StructAccess{lin, col, str, id}
	return exp
}

func (p StructAccess) Ejecutar(ast *environment.AST, env interface{}) environment.Symbol {

	var result, tempStruct environment.Symbol
	tempStruct = p.Struct.Ejecutar(ast, env) //sym{[1,2,3,4]}

	if tempStruct.Tipo == environment.STRUCT {

		if variable, ok := tempStruct.Valor.(map[string]environment.Symbol)[p.Id]; ok {
			variable.Mutable = true
			return variable
		}
		fmt.Println("No existe el elemento", p.Id)
		return result
	}
	fmt.Println("No es un structsss ", p.Lin, p.Col)
	return result
}
