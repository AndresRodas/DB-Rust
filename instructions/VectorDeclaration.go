package instructions

import (
	"OLC2/environment"
	"OLC2/interfaces"
	"fmt"
	arrayList "github.com/colegno/arraylist"
)

type VectorDeclaration struct {
	Lin       int
	Col       int
	Id        string
	Tipo      environment.TipoExpresion
	Mutable   bool
	Expresion interface{}
	VectorId  string
}

func NewVectorDeclaration(lin int, col int, id string, tipo environment.TipoExpresion, mut bool, exp interface{}, vecid string) VectorDeclaration {
	instr := VectorDeclaration{lin, col, id, tipo, mut, exp, vecid}
	return instr
}

func (p VectorDeclaration) Ejecutar(ast *environment.AST, env interface{}) environment.Symbol {

	var result environment.Symbol
	//simbolo expresion vector
	tmpExp := p.Expresion.(interfaces.Expression).Ejecutar(ast, env)
	//validando tipo
	if tmpExp.Tipo == environment.VECTOR {
		//validando tipos en indices
		for _, s := range tmpExp.Valor.(*arrayList.List).ToArray() {
			//cuando no trae tipo
			if p.Tipo == environment.WILDCARD {
				//comparar ids
				if p.VectorId == s.(environment.Symbol).Id || s.(environment.Symbol).Valor.(*arrayList.List).Len() == 0 {
					continue
				}
			}
			//tipos diferentes
			if s.(environment.Symbol).Tipo != p.Tipo {
				fmt.Println("Tipo de indice incorrecto")
				return result
			}
		}
		//result = tmpExp
		result = environment.Symbol{
			Lin:      tmpExp.Lin,
			Col:      tmpExp.Col,
			Id:       p.VectorId,
			Tipo:     environment.VECTOR,
			Valor:    tmpExp.Valor,
			Mutable:  p.Mutable,
			Capacity: tmpExp.Capacity,
			TipoArr:  p.Tipo,
		}
	} else {
		fmt.Println("Tipo de dato incorrecto")
		return result
	}
	//guardando variable
	env.(environment.Environment).SaveVariable(p.Id, result)

	return result
}
