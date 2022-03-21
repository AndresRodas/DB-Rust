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
	//obtener la funcion anteriormente guardada
	funcSym = env.(environment.Environment).GetFunction(p.Id)

	//definiendo nuevo entorno de funcion
	var funcEnv environment.Environment

	//creando entorno          //llamar funcion obtenerbglobal
	funcEnv = environment.NewEnvironment(p.GetGlobal(env), "FUNCTION-"+p.Id)

	//validar tamaños de parametros
	if funcSym.ListDec.Len() == p.Params.Len() { // PARAMS: ByRef{ expression, bool(mut) }
		//recorrer las listas						ListDec: ParamsDeclaration { id, tipo }
		for i := 0; i < p.Params.Len(); i++ {
			//valores a utilizar
			idNewVar := funcSym.ListDec.GetValue(i).(instructions.ParamsDeclaration).Id                                //guardando id
			typeNewVar := funcSym.ListDec.GetValue(i).(instructions.ParamsDeclaration).Tipo                            //guardando tipo
			symNewVar := p.Params.GetValue(i).(environment.ByReference).Exp.(interfaces.Expression).Ejecutar(ast, env) //guardando simbolo
			//Comparar los tipos de los parametros
			if typeNewVar == symNewVar.Tipo {
				//setear muteabilidad
				symNewVar.Mutable = true
				//guardar simbolo en entorno de funcSym ******************************************
				//funcSym.Ent.(environment.Environment).SaveVariable(idNewVar, symNewVar, typeNewVar, true)
				funcEnv.SaveVariable(idNewVar, symNewVar, typeNewVar, true)
			} else {
				fmt.Println("Los parametros son incorrectos")
				return result
			}
		}
	} else {
		fmt.Println("La cantidad de parametros no es la indicada")
		return result
	}

	//sacar entorno
	//funcEnv = funcSym.Ent.(environment.Environment)

	//ejecutar bloque con entorno funcSym.ent
	for _, s := range funcSym.Block.ToArray() {
		if strings.Contains(fmt.Sprintf("%T", s), "instructions") {
			result = s.(interfaces.Instruction).Ejecutar(ast, funcEnv)
			if result.Tipo == environment.BREAK || result.Tipo == environment.CONTINUE || result.Tipo == environment.RETURN { //BREAK, CONTINUE & RETURN
				return result
			}
		} else if strings.Contains(fmt.Sprintf("%T", s), "expressions") {
			result = s.(interfaces.Expression).Ejecutar(ast, funcEnv)
			if result.Tipo == environment.BREAK || result.Tipo == environment.CONTINUE || result.Tipo == environment.RETURN { //BREAK, CONTINUE & RETURN
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
