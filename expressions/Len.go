package expressions

import (
	"OLC2/environment"
	"OLC2/interfaces"
	"fmt"
	arrayList "github.com/colegno/arraylist"
)

type Len struct {
	Lin int
	Col int
	Exp interface{}
}

func NewLen(lin int, col int, Exp interface{}) Len {
	expr := Len{lin, col, Exp}
	return expr
}

func (p Len) Ejecutar(ast *environment.AST, env interface{}) environment.Symbol {
	var tmpSymbol environment.Symbol
	tmpSymbol = p.Exp.(interfaces.Expression).Ejecutar(ast, env)

	//validar que id sea vector
	if tmpSymbol.Tipo == environment.VECTOR || tmpSymbol.Tipo == environment.ARRAY {
		//retornar tamaño
		return environment.Symbol{Lin: p.Lin, Col: p.Col, Tipo: environment.INTEGER, Valor: tmpSymbol.Valor.(*arrayList.List).Len(), Mutable: true}
	} else {
		fmt.Println("Solo se puede hacer Len sobre vectores", p.Lin, p.Col)
		return environment.Symbol{Lin: p.Lin, Col: p.Col, Tipo: environment.INTEGER, Valor: 0, Mutable: true}
	}
}
