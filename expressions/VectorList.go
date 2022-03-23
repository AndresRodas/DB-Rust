package expressions

import (
	"OLC2/environment"
	"OLC2/interfaces"
	"fmt"
	arrayList "github.com/colegno/arraylist"
)

type VectorList struct {
	Lin  int
	Col  int
	Exp1 interfaces.Expression
	Exp2 interfaces.Expression
}

func NewVectorList(lin, col int, e1, e2 interfaces.Expression) VectorList {
	exp := VectorList{lin, col, e1, e2}
	return exp
}

func (p VectorList) Ejecutar(ast *environment.AST, env interface{}) environment.Symbol {

	var templist *arrayList.List
	templist = arrayList.New()

	val1 := p.Exp1.Ejecutar(ast, env)
	val2 := p.Exp2.Ejecutar(ast, env)

	if val2.Tipo == environment.INTEGER {
		for i := 0; i < val2.Valor.(int); i++ {
			tmpSym := environment.Symbol{
				Lin:     p.Lin,
				Col:     p.Col,
				Id:      "",
				Tipo:    val1.Tipo,
				Valor:   val1.Valor,
				Mutable: true,
			}
			templist.Add(tmpSym)
		}
	} else {
		fmt.Println("El tipo de cantidad es incorrecto")
	}

	return environment.Symbol{
		Lin:      p.Lin,
		Col:      p.Col,
		Id:       "",
		Tipo:     environment.VECTOR,
		Valor:    templist,
		Mutable:  true,
		Capacity: templist.Len() * 2,
		TipoArr:  val1.Tipo,
	}

}
