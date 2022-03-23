package expressions

import (
	"OLC2/environment"
	"OLC2/instructions"
	"OLC2/interfaces"
	"fmt"
	arrayList "github.com/colegno/arraylist"
	"strings"
)

type CallExp struct {
	Lin    int
	Col    int
	Id     string
	Params *arrayList.List
}

func NewCallExp(lin int, col int, id string, par *arrayList.List) CallExp {
	exp := CallExp{lin, col, id, par}
	return exp
}

func (p CallExp) Ejecutar(ast *environment.AST, env interface{}) environment.Symbol {
	var result environment.Symbol
	var funcSym environment.FunctionSymbol
	var listIdRef, listIdNew *arrayList.List
	listIdRef = arrayList.New()
	listIdNew = arrayList.New()
	//obtener la funcion anteriormente guardada
	funcSym = env.(environment.Environment).GetFunction(p.Id)

	//definiendo nuevo entorno de funcion
	var funcEnv environment.Environment

	//creando entorno          //llamar funcion obtenerbglobal
	funcEnv = environment.NewEnvironment(p.GetGlobal(env), "FUNCTION-"+p.Id)

	//validar tamaños de parametros
	if funcSym.ListDec.Len() == p.Params.Len() { // PARAMS: ByRef{ expression, bool(mut) }
		//recorrer las listas						ListDec: ParamsDeclaration { id, tipo, extratipo } funcSym
		for i := 0; i < p.Params.Len(); i++ {
			//valores a utilizar
			idNewVar := funcSym.ListDec.GetValue(i).(instructions.ParamsDeclaration).Id     //guardando id
			typeNewVar := funcSym.ListDec.GetValue(i).(instructions.ParamsDeclaration).Tipo //guardando tipo
			extraType := funcSym.ListDec.GetValue(i).(instructions.ParamsDeclaration).ExtraTipo
			symNewVar := p.Params.GetValue(i).(environment.ByReference).Exp.(interfaces.Expression).Ejecutar(ast, env) //guardando simbolo
			//Comparar los tipos de los parametros
			if typeNewVar == symNewVar.Tipo || extraType == symNewVar.Id {
				//setear muteabilidad y paso por referencia
				if p.Params.GetValue(i).(environment.ByReference).Ref && (typeNewVar == environment.ARRAY) || (typeNewVar == environment.VECTOR) {
					listIdRef.Add(symNewVar.Id)
					listIdNew.Add(idNewVar)
				}
				symNewVar.Mutable = true
				//guardar simbolo en entorno de funcSym
				symNewVar.ExtraTipo = funcSym.ListDec.GetValue(i).(instructions.ParamsDeclaration).ExtraTipo
				funcEnv.SaveVariable(idNewVar, symNewVar, typeNewVar, true)
			} else {
				fmt.Println("Los parametros son incorrectos")
				return result
			}
		}
	} else {
		fmt.Println("La cantidad de parametros no es la correcta")
		return result
	}

	//ejecutar bloque con entorno funcEnv
	for _, s := range funcSym.Block.ToArray() {
		if strings.Contains(fmt.Sprintf("%T", s), "instructions") {
			result = s.(interfaces.Instruction).Ejecutar(ast, funcEnv)
			p.SaveVarsReference(listIdRef, listIdNew, env.(environment.Environment), funcEnv)
			if result.Id == "BREAK" || result.Id == "CONTINUE" || result.Id == "RETURN" { //BREAK, CONTINUE & RETURN
				//comprobacion de tipo de retorno
				if result.Id == "RETURN" {
					if funcSym.TipoRetorno == result.Tipo || funcSym.IdTipo == result.ExtraTipo {
						return result
					} else {
						fmt.Println("El tipo del valor a retornar es incorrecto")
						return environment.Symbol{Lin: p.Lin, Col: p.Col, Id: "", Tipo: environment.NULL}
					}
				}
				return result
			}
		} else if strings.Contains(fmt.Sprintf("%T", s), "expressions") {
			result = s.(interfaces.Expression).Ejecutar(ast, funcEnv)
			p.SaveVarsReference(listIdRef, listIdNew, env.(environment.Environment), funcEnv)
			if result.Id == "BREAK" || result.Id == "CONTINUE" || result.Id == "RETURN" { //BREAK, CONTINUE & RETURN
				//comprobacion de tipo de retorno
				if result.Id == "RETURN" {
					if funcSym.TipoRetorno == result.Tipo || funcSym.IdTipo == result.ExtraTipo {
						return result
					} else {
						fmt.Println("El tipo del valor a retornar es incorrecto")
						return environment.Symbol{Lin: p.Lin, Col: p.Col, Id: "", Tipo: environment.NULL}
					}
				}
				return result
			}
		} else {
			fmt.Println("error en bloque")
		}
	}
	return result
}

func (p CallExp) GetGlobal(env interface{}) environment.Environment {
	//obteniendo entorno global
	var tmpEnvGbl environment.Environment
	tmpEnvGbl = env.(environment.Environment)
	for {
		if tmpEnvGbl.Id == "GLOBAL" {
			return tmpEnvGbl
		}
		if tmpEnvGbl.Anterior == nil {
			break
		} else {
			tmpEnvGbl = tmpEnvGbl.Anterior.(environment.Environment)
		}
	}
	return tmpEnvGbl
}

func (p CallExp) SaveVarsReference(listLocal, listFunc *arrayList.List, entLocal, entFunc environment.Environment) {
	//recorrer lista de ids
	for i := 0; i < listLocal.Len(); i++ {
		//obtener simbolo en func
		tmpSym := entFunc.GetVariable(listFunc.GetValue(i).(string))
		//setear simbolo entorno local
		entLocal.SetVariable(listLocal.GetValue(i).(string), tmpSym)
	}
}
