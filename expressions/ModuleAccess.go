package expressions

import (
	"OLC2/environment"
	"OLC2/instructions"
	"OLC2/interfaces"
	"fmt"
	arrayList "github.com/colegno/arraylist"
	"strings"
)

type ModuleAccess struct {
	Lin    int
	Col    int
	ListId *arrayList.List
	Exp    interface{}
}

func NewModuleAccess(lin int, col int, list *arrayList.List, ex interface{}) ModuleAccess {
	exp := ModuleAccess{lin, col, list, ex}
	return exp
}

func (p ModuleAccess) Ejecutar(ast *environment.AST, env interface{}) environment.Symbol {
	var result environment.Symbol
	//clonando lista por seguridad jeje
	listClone := p.ListId.Clone()
	//obteniendo primer id y limpiando
	id := listClone.GetValue(0).(string)
	listClone.RemoveAtIndex(0)
	//obteniendo primer modulo y seteandolo como temporal!!!!
	var tmpMod environment.ModuleSymbol
	tmpMod = env.(environment.Environment).GetModule(id)
	//creando entorno y seteandole el global como anterior
	var modEnv environment.Environment
	modEnv = environment.NewEnvironment(p.GetGlobal(env), tmpMod.Id+"(MODULO)")
	//clonando temporales
	var extraMod environment.ModuleSymbol
	var tmpEnv environment.Environment

	//validar tamaño del array
	if listClone.Len() > 0 {
		//recorrer resto de ids
		for _, md := range listClone.ToArray() { //itera lista de id de llamada --------LUGAR
			//recorrer hijos del modulo temporal
			for _, is := range tmpMod.Childs.ToArray() { //itera lista de ModuleObject{ visibilidad, tipo, instr }
				//validando visibilidad
				if is.(environment.ModuleObject).Visibilidad == environment.PUBLIC {
					//ejecutando funcion o guardando modulo en entorno
					if is.(environment.ModuleObject).Tipo == environment.INST {
						//ejecutar instrucciones con entorno de modulo
						result = is.(environment.ModuleObject).Instruction.(interfaces.Instruction).Ejecutar(ast, modEnv)
					} else if is.(environment.ModuleObject).Tipo == environment.MOD {
						//validar id con modulo
						if is.(environment.ModuleObject).Instruction.(instructions.Module).Id == md.(string) {
							//ejecuta modulo y lo guarda en entorno actual
							is.(environment.ModuleObject).Instruction.(interfaces.Instruction).Ejecutar(ast, modEnv)
							//seteando el modulo actual
							extraMod = modEnv.GetModule(md.(string))
							//seteando el entorno actual
							tmpEnv = environment.NewEnvironment(modEnv, extraMod.Id+"(MODULO)")
						}
					}
				}
			}
			tmpMod = extraMod
			modEnv = tmpEnv
		}
	}

	//agregando las variables al entorno a ejecutar********
	modEnv = p.CloneVars(env, modEnv)

	//ejecucion mod final
	for _, is := range tmpMod.Childs.ToArray() { //itera lista de ModuleObject{ visibilidad, tipo, instr }
		//validando visibilidad
		if is.(environment.ModuleObject).Visibilidad == environment.PUBLIC {
			//ejecutando funcion o guardando modulo en entorno
			if is.(environment.ModuleObject).Tipo == environment.INST {
				//ejecutar instrucciones con entorno de modulo
				result = is.(environment.ModuleObject).Instruction.(interfaces.Instruction).Ejecutar(ast, modEnv)
			}
		}
	}

	//EJECUCION DE EXPRESSION CON ENTORNO NUEVO
	if strings.Contains(fmt.Sprintf("%T", p.Exp), "instructions") {
		result = p.Exp.(interfaces.Instruction).Ejecutar(ast, modEnv)
		//BREAK, CONTINUE & RETURN
		if result.Id == "BREAK" || result.Id == "CONTINUE" || result.Id == "RETURN" {
			return result
		}
	} else if strings.Contains(fmt.Sprintf("%T", p.Exp), "expressions") {
		result = p.Exp.(interfaces.Expression).Ejecutar(ast, modEnv)
		if result.Id == "BREAK" || result.Id == "CONTINUE" || result.Id == "RETURN" { //BREAK & CONTINUE
			return result
		}
	} else {
		fmt.Println("error en bloque")
	}
	return result
}

func (p ModuleAccess) GetGlobal(env interface{}) environment.Environment {
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

func (p ModuleAccess) CloneVars(env1, env2 interface{}) environment.Environment {
	var result environment.Environment
	//recorriendo mapa del env1
	for key, element := range env1.(environment.Environment).Tabla {
		//setenado variable nueva
		env2.(environment.Environment).SaveVariable(key, element)
	}
	result = env2.(environment.Environment)
	return result
}
