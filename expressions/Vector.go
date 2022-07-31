package expressions

import (
	"OLC2/environment"
	"OLC2/interfaces"
	"fmt"
	arrayList "github.com/colegno/arraylist"
)

type Vector struct {
	Lin      int
	Col      int
	ListExp  *arrayList.List
	ListVec  interface{}
	Capacity interface{}
}

func NewVector(lin int, col int, list *arrayList.List, listv, capa interface{}) Vector {
	exp := Vector{lin, col, list, listv, capa}
	return exp
}

func (p Vector) Ejecutar(ast *environment.AST, env interface{}) environment.Symbol {
	var result environment.Symbol
	//lista temporal
	var tempExp *arrayList.List
	tempExp = arrayList.New()

	//viene un arreglo de la forma vec![1,2,3,4,5]
	if p.ListExp != nil {
		//tipo de elementos
		tempType := p.ListExp.GetValue(0).(interfaces.Expression).Ejecutar(ast, env).Tipo
		var tempArrType environment.TipoExpresion
		//validacion de tipo
		for _, s := range p.ListExp.ToArray() {
			tmpSym := s.(interfaces.Expression).Ejecutar(ast, env)
			if tmpSym.Tipo == tempType {
				tempExp.Add(tmpSym)
				tempArrType = tempType
				//result.TipoArr = tempType
			} else {
				fmt.Println("Error en el tipo de indices")
			}
		}
		//llenar vector
		result = environment.Symbol{
			Lin:      p.Lin,
			Col:      p.Col,
			Tipo:     environment.VECTOR,
			Valor:    tempExp,
			Mutable:  true,
			Capacity: tempExp.Len() * 2,
			TipoArr:  tempArrType,
		}
	} else if p.ListVec != nil { //viene un arrelgo de la forma vec![1;10]
		//simbolo temp vec
		tmpVecSym := p.ListVec.(interfaces.Expression).Ejecutar(ast, env)
		//guardar simbolo
		result = environment.Symbol{
			Lin:      tmpVecSym.Lin,
			Col:      tmpVecSym.Col,
			Tipo:     environment.VECTOR,
			Valor:    tmpVecSym.Valor,
			Mutable:  true,
			Capacity: tmpVecSym.Capacity,
			TipoArr:  tmpVecSym.TipoArr,
		}
	} else if p.Capacity != nil { //viene un arreglo de la forma Vec::with_capacity(10)
		var tempCap environment.Symbol
		capa := 0
		//seteando la capacidad
		tempCap = p.Capacity.(interfaces.Expression).Ejecutar(ast, env)
		if tempCap.Tipo == environment.INTEGER {
			if tempCap.Valor.(int) >= 0 {
				capa = tempCap.Valor.(int)
			} else {
				fmt.Println("El valor de la capacidad es incorrecto")
				capa = 0
			}
		} else {
			fmt.Println("El valor de la capacidad es incorrecto")
			capa = 0
		}
		//llenar vector
		result = environment.Symbol{
			Lin:      p.Lin,
			Col:      p.Col,
			Tipo:     environment.VECTOR,
			Valor:    tempExp,
			Mutable:  true,
			Capacity: capa,
		}
	} else { // viene un arreglo de la forma Vec::new()
		//llenar vector vacio
		result = environment.Symbol{
			Lin:     p.Lin,
			Col:     p.Col,
			Tipo:    environment.VECTOR,
			Valor:   tempExp,
			Mutable: true,
		}
	}
	return result
}
