package expressions

import (
	"OLC2/environment"
	"OLC2/interfaces"
	"fmt"
	arrayList "github.com/colegno/arraylist"
)

type Contains struct {
	Lin       int
	Col       int
	Id        string
	Expresion interfaces.Expression
}

func NewContains(lin int, col int, id string, exp interfaces.Expression) Contains {
	expr := Contains{lin, col, id, exp}
	return expr
}

func (p Contains) Ejecutar(ast *environment.AST, env interface{}) environment.Symbol {
	var result, tmpSymbol, tmpExp environment.Symbol
	tmpSymbol = env.(environment.Environment).GetVariable(p.Id)
	tmpExp = p.Expresion.Ejecutar(ast, env)
	result = environment.Symbol{Lin: p.Lin, Col: p.Col, Tipo: environment.BOOLEAN, Valor: false, Mutable: true}

	//validar que id sea vector
	if tmpSymbol.Tipo == environment.VECTOR {
		//validar tipo de vector con expression
		if tmpExp.Tipo == tmpSymbol.TipoArr {
			//recorrer expresion
			for _, s := range tmpSymbol.Valor.(*arrayList.List).ToArray() {
				if s.(environment.Symbol).Valor == tmpExp.Valor {
					return environment.Symbol{Lin: p.Lin, Col: p.Col, Tipo: environment.BOOLEAN, Valor: true, Mutable: true}
				}
			}
			return environment.Symbol{Lin: p.Lin, Col: p.Col, Tipo: environment.BOOLEAN, Valor: false, Mutable: true}
		} else {
			fmt.Println("Tipo de expression incorrecta")
			return result
		}
	} else {
		fmt.Println("Solo se puede hacer Contains sobre vectores")
		return result
	}
}
