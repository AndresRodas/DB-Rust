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
	Capacity  interface{}
	Mutable   bool
	Expresion interface{}
	VectorId  string
}

func NewVectorDeclaration(lin int, col int, id string, tipo environment.TipoExpresion, cap interface{}, mut bool, exp interface{}, vecid string) VectorDeclaration {
	instr := VectorDeclaration{lin, col, id, tipo, cap, mut, exp, vecid}
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
			if tempCap.Valor.(int) >= 0 {
				cap = tempCap.Valor.(int)
			} else {
				fmt.Println("El valor de la capacidad es incorrecto")
				cap = 0
			}
		} else {
			fmt.Println("El valor de la capacidad es incorrecto")
			cap = 0
		}
	} else {
		cap = 0
	}
	//si viene expression
	if p.Expresion != nil {
		tmpExp := p.Expresion.(interfaces.Expression).Ejecutar(ast, env)
		//validando tipo
		if tmpExp.Tipo == environment.VECTOR {
			//validando tipos en indices
			for _, s := range tmpExp.Valor.(*arrayList.List).ToArray() {
				//cuando no trae tipo
				if p.Tipo == environment.WILDCARD {
					//comparar ids
					if p.VectorId == s.(environment.Symbol).Id {
						continue
					}
				}
				//tipos diferentes
				if s.(environment.Symbol).Tipo != p.Tipo {
					fmt.Println("Tipo de dato incorrecto")
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
				Capacity: cap,
				TipoArr:  p.Tipo,
			}
		} else {
			fmt.Println("Tipo de dato incorrecto")
			return result
		}
	} else {
		//generando simbolo
		result = environment.Symbol{
			Lin:      p.Lin,
			Col:      p.Col,
			Id:       p.VectorId,
			Tipo:     environment.VECTOR,
			Valor:    templist,
			Mutable:  p.Mutable,
			Capacity: cap,
			TipoArr:  p.Tipo,
		}
	}

	//guardando variable
	env.(environment.Environment).SaveVariable(p.Id, result, environment.VECTOR, p.Mutable)

	return result
}
