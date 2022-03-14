package expressions

import (
	"OLC2/environment"
	"OLC2/interfaces"
	"fmt"
)

type Operation struct {
	Lin      int
	Col      int
	Op_izq   interfaces.Expression
	Operador string
	Op_der   interfaces.Expression
}

func (o Operation) Ejecutar(ast *environment.AST, env interface{}) environment.Symbol {
	var dominante environment.TipoExpresion
	suma_resta_dominante := [5][5]environment.TipoExpresion{
		{environment.INTEGER, environment.FLOAT, environment.STRING, environment.NULL, environment.NULL},
		{environment.FLOAT, environment.FLOAT, environment.STRING, environment.NULL, environment.NULL},
		{environment.STRING, environment.STRING, environment.STRING, environment.STRING, environment.NULL},
		{environment.NULL, environment.NULL, environment.STRING, environment.NULL, environment.NULL},
		{environment.NULL, environment.NULL, environment.NULL, environment.NULL, environment.NULL},
	}
	multi_division_dominante := [5][5]environment.TipoExpresion{
		{environment.INTEGER, environment.FLOAT, environment.NULL, environment.NULL, environment.NULL},
		{environment.FLOAT, environment.FLOAT, environment.NULL, environment.NULL, environment.NULL},
		{environment.NULL, environment.NULL, environment.NULL, environment.NULL, environment.NULL},
		{environment.NULL, environment.NULL, environment.NULL, environment.NULL, environment.NULL},
		{environment.NULL, environment.NULL, environment.NULL, environment.NULL, environment.NULL},
	}
	relacional_dominante := [5][5]environment.TipoExpresion{
		{environment.INTEGER, environment.FLOAT, environment.NULL, environment.NULL, environment.NULL},
		{environment.FLOAT, environment.FLOAT, environment.NULL, environment.NULL, environment.NULL},
		{environment.NULL, environment.NULL, environment.NULL, environment.NULL, environment.NULL},
		{environment.NULL, environment.NULL, environment.NULL, environment.NULL, environment.NULL},
		{environment.NULL, environment.NULL, environment.NULL, environment.NULL, environment.NULL},
	}

	if o.Operador != "MENOS_UNARIO" && o.Operador != "NOT" {
		op1 := o.Op_izq.Ejecutar(ast, env)
		op2 := o.Op_der.Ejecutar(ast, env)
		switch o.Operador {
		case "+":
			{
				dominante = suma_resta_dominante[op1.Tipo][op2.Tipo]
				if dominante == environment.INTEGER {
					return environment.Symbol{Lin: o.Lin, Col: o.Col, Id: "", Tipo: dominante, Valor: op1.Valor.(int) + op2.Valor.(int), Mutable: true}
				} else if dominante == environment.FLOAT {
					return environment.Symbol{Lin: o.Lin, Col: o.Col, Id: "", Tipo: dominante, Valor: op1.Valor.(float64) + op2.Valor.(float64), Mutable: true}
				} else if dominante == environment.STRING {
					r1 := fmt.Sprintf("%v", op1.Valor)
					r2 := fmt.Sprintf("%v", op2.Valor)
					return environment.Symbol{Lin: o.Lin, Col: o.Col, Id: "", Tipo: dominante, Valor: r1 + r2, Mutable: true}
				} else {
					fmt.Print("ERROR: No es posible sumar")
				}
			}
		case "-":
			{
				dominante = suma_resta_dominante[op1.Tipo][op2.Tipo]

				if dominante == environment.INTEGER {

					return environment.Symbol{Lin: o.Lin, Col: o.Col, Id: "", Tipo: dominante, Valor: op1.Valor.(int) - op2.Valor.(int), Mutable: true}

				} else if dominante == environment.FLOAT {
					return environment.Symbol{Lin: o.Lin, Col: o.Col, Id: "", Tipo: dominante, Valor: op1.Valor.(float64) - op2.Valor.(float64), Mutable: true}

				} else {
					fmt.Print("ERROR: No es posible restar")
				}
			}
		case "*":
			{
				dominante = multi_division_dominante[op1.Tipo][op2.Tipo]

				if dominante == environment.INTEGER {
					return environment.Symbol{Lin: o.Lin, Col: o.Col, Id: "", Tipo: dominante, Valor: op1.Valor.(int) * op2.Valor.(int), Mutable: true}

				} else if dominante == environment.FLOAT {
					return environment.Symbol{Lin: o.Lin, Col: o.Col, Id: "", Tipo: dominante, Valor: op1.Valor.(float64) * op2.Valor.(float64), Mutable: true}

				} else {
					fmt.Print("ERROR: No es posible Multiplicar")
				}

			}
		case "/":
			{
				dominante = multi_division_dominante[op1.Tipo][op2.Tipo]

				if dominante == environment.INTEGER {
					return environment.Symbol{Lin: o.Lin, Col: o.Col, Id: "", Tipo: dominante, Valor: op1.Valor.(int) / op2.Valor.(int), Mutable: true}

				} else if dominante == environment.FLOAT {
					return environment.Symbol{Lin: o.Lin, Col: o.Col, Id: "", Tipo: dominante, Valor: op1.Valor.(float64) / op2.Valor.(float64), Mutable: true}

				} else {
					fmt.Print("ERROR: No es posible Dividir")
				}

			}
		case "<":
			{
				dominante = relacional_dominante[op1.Tipo][op2.Tipo]

				if dominante == environment.INTEGER {

					return environment.Symbol{Lin: o.Lin, Col: o.Col, Id: "", Tipo: environment.BOOLEAN, Valor: op1.Valor.(int) < op2.Valor.(int), Mutable: true}

				} else if dominante == environment.FLOAT {
					return environment.Symbol{Lin: o.Lin, Col: o.Col, Id: "", Tipo: environment.BOOLEAN, Valor: op1.Valor.(float64) < op2.Valor.(float64), Mutable: true}

				} else {
					fmt.Print("ERROR: No es posible comparar <")
				}
			}
		case ">":
			{
				dominante = relacional_dominante[op1.Tipo][op2.Tipo]

				if dominante == environment.INTEGER {

					return environment.Symbol{Lin: o.Lin, Col: o.Col, Id: "", Tipo: environment.BOOLEAN, Valor: op1.Valor.(int) > op2.Valor.(int), Mutable: true}

				} else if dominante == environment.FLOAT {
					return environment.Symbol{Lin: o.Lin, Col: o.Col, Id: "", Tipo: environment.BOOLEAN, Valor: op1.Valor.(float64) > op2.Valor.(float64), Mutable: true}

				} else {
					fmt.Print("ERROR: No es posible comparar <")
				}
			}
		case "<=":
			{
				dominante = relacional_dominante[op1.Tipo][op2.Tipo]

				if dominante == environment.INTEGER {

					return environment.Symbol{Lin: o.Lin, Col: o.Col, Id: "", Tipo: environment.BOOLEAN, Valor: op1.Valor.(int) <= op2.Valor.(int), Mutable: true}

				} else if dominante == environment.FLOAT {
					return environment.Symbol{Lin: o.Lin, Col: o.Col, Id: "", Tipo: environment.BOOLEAN, Valor: op1.Valor.(float64) <= op2.Valor.(float64), Mutable: true}

				} else {
					fmt.Print("ERROR: No es posible comparar <")
				}
			}
		case ">=":
			{
				dominante = relacional_dominante[op1.Tipo][op2.Tipo]

				if dominante == environment.INTEGER {

					return environment.Symbol{Lin: o.Lin, Col: o.Col, Id: "", Tipo: environment.BOOLEAN, Valor: op1.Valor.(int) >= op2.Valor.(int), Mutable: true}

				} else if dominante == environment.FLOAT {
					return environment.Symbol{Lin: o.Lin, Col: o.Col, Id: "", Tipo: environment.BOOLEAN, Valor: op1.Valor.(float64) >= op2.Valor.(float64), Mutable: true}

				} else {
					fmt.Print("ERROR: No es posible comparar <")
				}
			}
		}
	}
	var result interface{}
	return environment.Symbol{Lin: o.Lin, Col: o.Col, Id: "", Tipo: environment.NULL, Valor: result, Mutable: true}
}

func NewOperation(lin, col int, Op1 interfaces.Expression, Operador string, Op2 interfaces.Expression) Operation {
	exp := Operation{Lin: lin, Col: col, Op_izq: Op1, Operador: Operador, Op_der: Op2}
	return exp
}
