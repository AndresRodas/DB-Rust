package expressions

import (
	"OLC2/environment"
	"OLC2/interfaces"
	"fmt"
	"strconv"
)

type Cast struct {
	Lin       int
	Col       int
	Expresion interface{}
	Tipo      environment.TipoExpresion
}

func NewCast(lin int, col int, exp interface{}, tipo environment.TipoExpresion) Cast {
	expr := Cast{lin, col, exp, tipo}
	return expr
}

func (p Cast) Ejecutar(ast *environment.AST, env interface{}) environment.Symbol {

	var tmpExp, result environment.Symbol
	tmpExp = p.Expresion.(interfaces.Expression).Ejecutar(ast, env)
	if p.Tipo == environment.FLOAT { //si es float
		if tmpExp.Tipo == environment.INTEGER || tmpExp.Tipo == environment.FLOAT {
			valor, _ := strconv.ParseFloat(fmt.Sprintf("%v", tmpExp.Valor), 64)
			return environment.Symbol{Lin: p.Lin, Col: p.Col, Id: "", Tipo: p.Tipo, Valor: valor}
		} else {
			fmt.Println("No se pueden castear los tipos indicados")
		}
	} else if p.Tipo == environment.INTEGER { //si es entero
		if tmpExp.Tipo == environment.INTEGER || tmpExp.Tipo == environment.FLOAT {
			valor, _ := strconv.ParseFloat(fmt.Sprintf("%v", tmpExp.Valor), 64)
			return environment.Symbol{Lin: p.Lin, Col: p.Col, Id: "", Tipo: p.Tipo, Valor: int(valor)}
		} else if tmpExp.Tipo == environment.BOOLEAN {
			var tmpBool int
			if tmpExp.Valor == true {
				tmpBool = 1
			} else if tmpExp.Valor == false {
				tmpBool = 0
			}
			return environment.Symbol{Lin: p.Lin, Col: p.Col, Id: "", Tipo: p.Tipo, Valor: tmpBool}
		} else {
			fmt.Println("No se pueden castear los tipos indicados")
		}
	}
	return result
}
