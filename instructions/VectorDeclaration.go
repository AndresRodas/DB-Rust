package instructions

import (
	"OLC2/environment"
	"OLC2/interfaces"
	"fmt"
	arrayList "github.com/colegno/arraylist"
)

type VectorDeclaration struct {
	Lin      int
	Col      int
	Id       string
	Tipo     environment.TipoExpresion
	Capacity interface{}
	Mutable  bool
}

func NewVectorDeclaration(lin int, col int, id string, tipo environment.TipoExpresion, cap interface{}, mut bool) VectorDeclaration {
	instr := VectorDeclaration{lin, col, id, tipo, cap, mut}
	return instr
}

func (p VectorDeclaration) Ejecutar(ast *environment.AST, env interface{}) environment.Symbol {

	var templist *arrayList.List
	var tempCap, result environment.Symbol
	var cap int
	templist = arrayList.New()
	//seteando la capacidad
	if p.Capacity != nil {
		tempCap = p.Capacity.(interfaces.Expression).Ejecutar(ast, env)
		if tempCap.Tipo == environment.INTEGER {
			cap = tempCap.Valor.(int)
		} else {
			fmt.Println("El valor de la capacidad es incorrecto")
			cap = 0
		}
	} else {
		cap = 0
	}
	//generando simbolo
	result = environment.Symbol{
		Lin:      p.Lin,
		Col:      p.Col,
		Id:       "",
		Tipo:     environment.VECTOR,
		Valor:    templist,
		Mutable:  p.Mutable,
		Capacity: cap,
		TipoArr:  p.Tipo,
	}
	//guardando variable
	env.(environment.Environment).SaveVariable(p.Id, result, environment.VECTOR, p.Mutable)

	return result
}
