package expressions

import (
	"OLC2/environment"
	"OLC2/interfaces"
	"fmt"
	"strconv"
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

	suma_resta_dominante := [12][12]environment.TipoExpresion{
		//		INTEGER			FLOAT				STRING				STR				CHAR			BOOLEAN
		//INTEGER
		{environment.INTEGER, environment.NULL, environment.NULL, environment.NULL, environment.NULL, environment.NULL},
		//FLOAT
		{environment.NULL, environment.FLOAT, environment.NULL, environment.NULL, environment.NULL, environment.NULL},
		//STRING
		{environment.NULL, environment.NULL, environment.NULL, environment.STRING, environment.NULL, environment.NULL},
		//STR
		{environment.NULL, environment.NULL, environment.NULL, environment.NULL, environment.NULL, environment.NULL},
		//CHAR
		{environment.NULL, environment.NULL, environment.NULL, environment.NULL, environment.NULL, environment.NULL},
		//BOOLEAN
		{environment.NULL, environment.NULL, environment.NULL, environment.NULL, environment.NULL, environment.NULL},
	}
	multi_division_dominante := [12][12]environment.TipoExpresion{
		//		INTEGER			FLOAT				STRING				STR				CHAR			BOOLEAN
		//INTEGER
		{environment.INTEGER, environment.FLOAT, environment.NULL, environment.NULL, environment.NULL, environment.NULL},
		//FLOAT
		{environment.FLOAT, environment.FLOAT, environment.NULL, environment.NULL, environment.NULL, environment.NULL},
		//STRING
		{environment.NULL, environment.NULL, environment.NULL, environment.NULL, environment.NULL, environment.NULL},
		//STR
		{environment.NULL, environment.NULL, environment.NULL, environment.NULL, environment.NULL, environment.NULL},
		//CHAR
		{environment.NULL, environment.NULL, environment.NULL, environment.NULL, environment.NULL, environment.NULL},
		//BOOLEAN
		{environment.NULL, environment.NULL, environment.NULL, environment.NULL, environment.NULL, environment.NULL},
	}
	relacional_dominante := [12][12]environment.TipoExpresion{
		//		INTEGER			FLOAT				STRING				STR				CHAR			BOOLEAN
		//INTEGER
		{environment.INTEGER, environment.FLOAT, environment.NULL, environment.NULL, environment.NULL, environment.NULL},
		//FLOAT
		{environment.FLOAT, environment.FLOAT, environment.NULL, environment.NULL, environment.NULL, environment.NULL},
		//STRING
		{environment.NULL, environment.NULL, environment.NULL, environment.NULL, environment.NULL, environment.NULL},
		//STR
		{environment.NULL, environment.NULL, environment.NULL, environment.NULL, environment.NULL, environment.NULL},
		//CHAR
		{environment.NULL, environment.NULL, environment.NULL, environment.NULL, environment.NULL, environment.NULL},
		//BOOLEAN
		{environment.NULL, environment.NULL, environment.NULL, environment.NULL, environment.NULL, environment.NULL},
	}
	var op1, op2 environment.Symbol
	op1 = o.Op_izq.Ejecutar(ast, env)
	if o.Op_der != nil {
		op2 = o.Op_der.Ejecutar(ast, env)
	}
	switch o.Operador {
	case "+":
		{
			dominante = suma_resta_dominante[op1.Tipo][op2.Tipo]
			if dominante == environment.INTEGER {
				return environment.Symbol{Lin: o.Lin, Col: o.Col, Id: "", Tipo: dominante, Valor: op1.Valor.(int) + op2.Valor.(int), Mutable: true}
			} else if dominante == environment.FLOAT {
				val1, _ := strconv.ParseFloat(fmt.Sprintf("%v", op1.Valor), 64)
				val2, _ := strconv.ParseFloat(fmt.Sprintf("%v", op2.Valor), 64)
				return environment.Symbol{Lin: o.Lin, Col: o.Col, Id: "", Tipo: dominante, Valor: val1 + val2, Mutable: true}
			} else if dominante == environment.STRING {
				r1 := fmt.Sprintf("%v", op1.Valor)
				r2 := fmt.Sprintf("%v", op2.Valor)
				return environment.Symbol{Lin: o.Lin, Col: o.Col, Id: "", Tipo: dominante, Valor: r1 + r2, Mutable: true}
			} else {
				fmt.Println("ERROR: No es posible sumar")
			}
		}
	case "-":
		{
			dominante = suma_resta_dominante[op1.Tipo][op2.Tipo]

			if dominante == environment.INTEGER {

				return environment.Symbol{Lin: o.Lin, Col: o.Col, Id: "", Tipo: dominante, Valor: op1.Valor.(int) - op2.Valor.(int), Mutable: true}

			} else if dominante == environment.FLOAT {
				val1, _ := strconv.ParseFloat(fmt.Sprintf("%v", op1.Valor), 64)
				val2, _ := strconv.ParseFloat(fmt.Sprintf("%v", op2.Valor), 64)
				return environment.Symbol{Lin: o.Lin, Col: o.Col, Id: "", Tipo: dominante, Valor: val1 - val2, Mutable: true}

			} else {
				fmt.Println("ERROR: No es posible restar")
			}
		}
	case "*":
		{
			dominante = multi_division_dominante[op1.Tipo][op2.Tipo]

			if dominante == environment.INTEGER {
				return environment.Symbol{Lin: o.Lin, Col: o.Col, Id: "", Tipo: dominante, Valor: op1.Valor.(int) * op2.Valor.(int), Mutable: true}

			} else if dominante == environment.FLOAT {
				val1, _ := strconv.ParseFloat(fmt.Sprintf("%v", op1.Valor), 64)
				val2, _ := strconv.ParseFloat(fmt.Sprintf("%v", op2.Valor), 64)
				return environment.Symbol{Lin: o.Lin, Col: o.Col, Id: "", Tipo: dominante, Valor: val1 * val2, Mutable: true}

			} else {
				fmt.Println("ERROR: No es posible Multiplicar")
			}

		}
	case "/":
		{
			dominante = multi_division_dominante[op1.Tipo][op2.Tipo]

			if dominante == environment.INTEGER {
				if op2.Valor.(int) != 0 {
					return environment.Symbol{Lin: o.Lin, Col: o.Col, Id: "", Tipo: dominante, Valor: op1.Valor.(int) / op2.Valor.(int), Mutable: true}
				} else {
					fmt.Println("ERROR: No es posible dividir en cero")
				}

			} else if dominante == environment.FLOAT {
				val1, _ := strconv.ParseFloat(fmt.Sprintf("%v", op1.Valor), 64)
				val2, _ := strconv.ParseFloat(fmt.Sprintf("%v", op2.Valor), 64)
				if val2 != 0 {
					return environment.Symbol{Lin: o.Lin, Col: o.Col, Id: "", Tipo: dominante, Valor: val1 / val2, Mutable: true}
				} else {
					fmt.Println("ERROR: No es posible dividir en cero")
				}

			} else {
				fmt.Println("ERROR: No es posible Dividir")
			}

		}
	case "%":
		{
			dominante = multi_division_dominante[op1.Tipo][op2.Tipo]

			if dominante == environment.INTEGER {
				return environment.Symbol{Lin: o.Lin, Col: o.Col, Id: "", Tipo: dominante, Valor: op1.Valor.(int) % op2.Valor.(int), Mutable: true}

			} else if dominante == environment.FLOAT {
				val1, _ := strconv.ParseFloat(fmt.Sprintf("%v", op1.Valor), 64)
				val2, _ := strconv.ParseFloat(fmt.Sprintf("%v", op2.Valor), 64)
				return environment.Symbol{Lin: o.Lin, Col: o.Col, Id: "", Tipo: dominante, Valor: int(val1) % int(val2), Mutable: true}

			} else {
				fmt.Println("ERROR: No es posible Dividir")
			}

		}
	case "<":
		{
			dominante = relacional_dominante[op1.Tipo][op2.Tipo]
			if dominante == environment.INTEGER {
				return environment.Symbol{Lin: o.Lin, Col: o.Col, Id: "", Tipo: environment.BOOLEAN, Valor: op1.Valor.(int) < op2.Valor.(int), Mutable: true}
			} else if dominante == environment.FLOAT {
				val1, _ := strconv.ParseFloat(fmt.Sprintf("%v", op1.Valor), 64)
				val2, _ := strconv.ParseFloat(fmt.Sprintf("%v", op2.Valor), 64)
				return environment.Symbol{Lin: o.Lin, Col: o.Col, Id: "", Tipo: environment.BOOLEAN, Valor: val1 < val2, Mutable: true}
			} else {
				fmt.Println("ERROR: No es posible comparar <")
			}
		}
	case ">":
		{
			dominante = relacional_dominante[op1.Tipo][op2.Tipo]
			if dominante == environment.INTEGER {
				return environment.Symbol{Lin: o.Lin, Col: o.Col, Id: "", Tipo: environment.BOOLEAN, Valor: op1.Valor.(int) > op2.Valor.(int), Mutable: true}
			} else if dominante == environment.FLOAT {
				val1, _ := strconv.ParseFloat(fmt.Sprintf("%v", op1.Valor), 64)
				val2, _ := strconv.ParseFloat(fmt.Sprintf("%v", op2.Valor), 64)
				return environment.Symbol{Lin: o.Lin, Col: o.Col, Id: "", Tipo: environment.BOOLEAN, Valor: val1 > val2, Mutable: true}
			} else {
				fmt.Println("ERROR: No es posible comparar <")
			}
		}
	case "<=":
		{
			dominante = relacional_dominante[op1.Tipo][op2.Tipo]
			if dominante == environment.INTEGER {
				return environment.Symbol{Lin: o.Lin, Col: o.Col, Id: "", Tipo: environment.BOOLEAN, Valor: op1.Valor.(int) <= op2.Valor.(int), Mutable: true}
			} else if dominante == environment.FLOAT {
				val1, _ := strconv.ParseFloat(fmt.Sprintf("%v", op1.Valor), 64)
				val2, _ := strconv.ParseFloat(fmt.Sprintf("%v", op2.Valor), 64)
				return environment.Symbol{Lin: o.Lin, Col: o.Col, Id: "", Tipo: environment.BOOLEAN, Valor: val1 <= val2, Mutable: true}
			} else {
				fmt.Println("ERROR: No es posible comparar <")
			}
		}
	case ">=":
		{
			dominante = relacional_dominante[op1.Tipo][op2.Tipo]
			if dominante == environment.INTEGER {
				return environment.Symbol{Lin: o.Lin, Col: o.Col, Id: "", Tipo: environment.BOOLEAN, Valor: op1.Valor.(int) >= op2.Valor.(int), Mutable: true}
			} else if dominante == environment.FLOAT {
				val1, _ := strconv.ParseFloat(fmt.Sprintf("%v", op1.Valor), 64)
				val2, _ := strconv.ParseFloat(fmt.Sprintf("%v", op2.Valor), 64)
				return environment.Symbol{Lin: o.Lin, Col: o.Col, Id: "", Tipo: environment.BOOLEAN, Valor: val1 >= val2, Mutable: true}
			} else {
				fmt.Println("ERROR: No es posible comparar <")
			}
		}
	case "==":
		{
			if op1.Tipo == op2.Tipo {
				return environment.Symbol{Lin: o.Lin, Col: o.Col, Id: "", Tipo: environment.BOOLEAN, Valor: op1.Valor == op2.Valor, Mutable: true}
			} else {
				fmt.Println("ERROR: No es posible comparar == ", o.Lin, o.Col)
			}
		}
	case "!=":
		{
			if op1.Tipo == op2.Tipo {
				return environment.Symbol{Lin: o.Lin, Col: o.Col, Id: "", Tipo: environment.BOOLEAN, Valor: op1.Valor != op2.Valor, Mutable: true}
			} else {
				fmt.Println("ERROR: No es posible comparar !=")
			}
		}
	case "&&":
		{
			if (op1.Tipo == environment.BOOLEAN) && (op2.Tipo == environment.BOOLEAN) {
				return environment.Symbol{Lin: o.Lin, Col: o.Col, Id: "", Tipo: environment.BOOLEAN, Valor: op1.Valor.(bool) && op2.Valor.(bool), Mutable: true}
			} else {
				fmt.Print("ERROR: tipo no compatible &&")
			}
		}
	case "||":
		{
			if (op1.Tipo == environment.BOOLEAN) && (op2.Tipo == environment.BOOLEAN) {
				return environment.Symbol{Lin: o.Lin, Col: o.Col, Id: "", Tipo: environment.BOOLEAN, Valor: op1.Valor.(bool) || op2.Valor.(bool), Mutable: true}
			} else {
				fmt.Println("ERROR: tipo no compatible ||")
			}
		}
	case "!":
		{
			if op1.Tipo == environment.BOOLEAN {
				return environment.Symbol{Lin: o.Lin, Col: o.Col, Id: "", Tipo: environment.BOOLEAN, Valor: !op1.Valor.(bool), Mutable: true}
			} else {
				fmt.Println("ERROR: tipo no compatible !")
			}
		}
	case "MENOS_UNARIO":
		{
			if op1.Tipo == environment.INTEGER {
				return environment.Symbol{Lin: o.Lin, Col: o.Col, Id: "", Tipo: op1.Tipo, Valor: 0 - op1.Valor.(int), Mutable: true}
			} else if op1.Tipo == environment.FLOAT {
				val1, _ := strconv.ParseFloat(fmt.Sprintf("%v", op1.Valor), 64)
				return environment.Symbol{Lin: o.Lin, Col: o.Col, Id: "", Tipo: op1.Tipo, Valor: 0 - val1, Mutable: true}
			} else {
				fmt.Println("ERROR: tipo no compatible -")
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
