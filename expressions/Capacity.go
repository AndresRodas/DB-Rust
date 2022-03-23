package expressions

import (
	"OLC2/environment"
	"fmt"
)

type Capacity struct {
	Lin int
	Col int
	Id  string
}

func NewCapacity(lin int, col int, id string) Capacity {
	expr := Capacity{lin, col, id}
	return expr
}

func (p Capacity) Ejecutar(ast *environment.AST, env interface{}) environment.Symbol {
	var tmpSymbol environment.Symbol
	tmpSymbol = env.(environment.Environment).GetVariable(p.Id)

	//validar que id sea vector
	if tmpSymbol.Tipo == environment.VECTOR {
		//retornar tamaño
		return environment.Symbol{Lin: p.Lin, Col: p.Col, Tipo: environment.INTEGER, Valor: tmpSymbol.Capacity, Mutable: true}
	} else {
		fmt.Println("Solo se puede hacer Len sobre vectores")
		return environment.Symbol{Lin: p.Lin, Col: p.Col, Tipo: environment.INTEGER, Valor: 0, Mutable: true}
	}
}
