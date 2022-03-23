package expressions

import (
	"OLC2/environment"
	"OLC2/interfaces"
	"fmt"
	arrayList "github.com/colegno/arraylist"
)

type Vector struct {
	Lin     int
	Col     int
	ListExp *arrayList.List
}

func NewVector(lin int, col int, list *arrayList.List) Vector {
	exp := Vector{lin, col, list}
	return exp
}

func (p Vector) Ejecutar(ast *environment.AST, env interface{}) environment.Symbol {

	var tempExp *arrayList.List
	tempExp = arrayList.New()

	tempType := p.ListExp.GetValue(0).(interfaces.Expression).Ejecutar(ast, env).Tipo

	for _, s := range p.ListExp.ToArray() {
		tmpSym := s.(interfaces.Expression).Ejecutar(ast, env)
		if tmpSym.Tipo == tempType {
			tempExp.Add(tmpSym)
		} else {
			fmt.Println("Error en el tipo del vector")
		}
	}

	return environment.Symbol{
		Lin:      p.Lin,
		Col:      p.Col,
		Id:       "",
		Tipo:     environment.VECTOR,
		Valor:    tempExp,
		Mutable:  true,
		Capacity: tempExp.Len() * 2,
	}
}
