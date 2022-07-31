package expressions

import (
	"OLC2/environment"
	"OLC2/instructions"
	"OLC2/interfaces"
	"fmt"
	arrayList "github.com/colegno/arraylist"
	"reflect"
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

	//creando entorno
	if strings.Contains(env.(environment.Environment).Id, "MODULO") {
		//si es modulo se usa el anterior como anterior
		funcEnv = environment.NewEnvironment(env.(environment.Environment), p.Id+"(FUNCTION)")
	} else {
		//si no es modulo se usa el global como anterior
		funcEnv = environment.NewEnvironment(p.GetGlobal(env), p.Id+"(FUNCTION)")
	}

	//si la funcion existe..
	if funcSym.TipoRetorno != environment.NULL {
		//validar tamaños de parametros
		//fmt.Println("LOS TAMAÑOS DEBEN SER IGUALES... GUARDADA: ", funcSym.ListDec.Len(), ", ENTRA: ", p.Params.Len())
		//fmt.Println(funcSym.Id, funcSym.ListDec.ToArray())
		if funcSym.ListDec.Len() == p.Params.Len() { // PARAMS: ByRef{ expression, bool(mut) }  "entrante"
			//recorrer las listas						ListDec: ParamsDeclaration { id, tipo, extratipo, Vectipo } funcSym "guardada"
			for i := 0; i < p.Params.Len(); i++ {
				//valores a utilizar
				idNewVar := funcSym.ListDec.GetValue(i).(instructions.ParamsDeclaration).Id                                //guardando id func guardada
				typeNewVar := funcSym.ListDec.GetValue(i).(instructions.ParamsDeclaration).Tipo                            //guardando tipo func guardada
				extraType := funcSym.ListDec.GetValue(i).(instructions.ParamsDeclaration).ExtraTipo                        //extra tipo func guardada
				vecTipo := funcSym.ListDec.GetValue(i).(instructions.ParamsDeclaration).Vectipo                            //guardando tipo de vector
				symNewVar := p.Params.GetValue(i).(environment.ByReference).Exp.(interfaces.Expression).Ejecutar(ast, env) //guardando simbolo que viene

				//valor por referencia
				var idnewer string
				if reflect.TypeOf(p.Params.GetValue(i).(environment.ByReference).Exp) == reflect.TypeOf(CallVar{}) {
					identificador := p.Params.GetValue(i).(environment.ByReference).Exp.(CallVar)
					idnewer = identificador.Id
				}
				//Comparar los tipos de los parametros
				if typeNewVar == environment.VECTOR {
					//es un vector con tipo adentro
					//fmt.Println("tipo func alma: ", vecTipo)
					//fmt.Println("tipo func exp: ", symNewVar.TipoArr)
					if vecTipo == symNewVar.TipoArr {
						//setear muteabilidad y paso por referencia
						if p.Params.GetValue(i).(environment.ByReference).Ref {
							listIdRef.Add(idnewer)
							listIdNew.Add(idNewVar)
						}
						symNewVar.Mutable = true
						//guardar simbolo en entorno de funcSym
						symNewVar.ExtraTipo = funcSym.ListDec.GetValue(i).(instructions.ParamsDeclaration).ExtraTipo
						//symNewVar.Id = idNewVar
						funcEnv.SaveVariable(idNewVar, symNewVar)
					} else {
						fmt.Println("Los parametros son incorrectos nooooo ", p.Lin, p.Col)
						return result
					}

				} else if typeNewVar == symNewVar.Tipo || (extraType == symNewVar.Id && extraType != "") { //se cambio id por exratipo
					fmt.Println("saved type: ", typeNewVar)
					fmt.Println("new type: ", symNewVar.Tipo)
					fmt.Println("extra saved type: ", extraType)
					fmt.Println("new extra type: ", symNewVar.Id)
					//setear muteabilidad y paso por referencia
					if p.Params.GetValue(i).(environment.ByReference).Ref && (typeNewVar == environment.ARRAY) || (typeNewVar == environment.VECTOR) || (typeNewVar == environment.STRUCT) {
						listIdRef.Add(idnewer)
						listIdNew.Add(idNewVar)
					}
					symNewVar.Mutable = true
					//guardar simbolo en entorno de funcSym
					symNewVar.ExtraTipo = funcSym.ListDec.GetValue(i).(instructions.ParamsDeclaration).ExtraTipo
					//symNewVar.Id = idNewVar
					fmt.Println("GUARDANDO EL SIMBOLO: ", idNewVar)
					funcEnv.SaveVariable(idNewVar, symNewVar)
				} else {
					fmt.Println("Los parametros son incorrectos", p.Lin, p.Col)
					return result
				}
			}
		} else {
			fmt.Println("La cantidad de parametros no es la correcta, Lin: ", p.Lin, ", Col: ", p.Col)
			return result
		}
	} else {
		return result
	}

	//ejecutar bloque con entorno funcEnv
	for _, s := range funcSym.Block.ToArray() {
		if strings.Contains(fmt.Sprintf("%T", s), "instructions") {
			result = s.(interfaces.Instruction).Ejecutar(ast, funcEnv)
			p.SaveVarsReference(listIdRef, listIdNew, env.(environment.Environment), funcEnv)
			//BREAK, CONTINUE & RETURN
			if result.BreakFlag {
				result.BreakFlag = false
				return result
			} else if result.ContinueFlag {
				result.ContinueFlag = false
				return result
			} else if result.ReturnFlag {
				result.ReturnFlag = false

				if funcSym.TipoRetorno == result.Tipo || funcSym.IdTipo == result.ExtraTipo {
					return result
				} else {
					fmt.Println("El tipo del valor a retornar es incorrecto")
					return environment.Symbol{Lin: p.Lin, Col: p.Col, Id: "", Tipo: environment.NULL}
				}
			}

		} else if strings.Contains(fmt.Sprintf("%T", s), "expressions") {
			result = s.(interfaces.Expression).Ejecutar(ast, funcEnv)
			p.SaveVarsReference(listIdRef, listIdNew, env.(environment.Environment), funcEnv)
			//BREAK, CONTINUE & RETURN
			if result.BreakFlag {
				result.BreakFlag = false
				return result
			} else if result.ContinueFlag {
				result.ContinueFlag = false
				return result
			} else if result.ReturnFlag {
				result.ReturnFlag = false
				if funcSym.TipoRetorno == result.Tipo || funcSym.IdTipo == result.ExtraTipo {
					return result
				} else {
					fmt.Println("El tipo del valor a retornar es incorrecto")
					return environment.Symbol{Lin: p.Lin, Col: p.Col, Id: "", Tipo: environment.NULL}
				}
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
