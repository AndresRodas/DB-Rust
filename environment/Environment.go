package environment

import (
	"fmt"
	arrayList "github.com/colegno/arraylist"
	"strings"
)

type Environment struct {
	Anterior  interface{}
	Tabla     map[string]Symbol
	Structs   map[string]Symbol
	Functions map[string]FunctionSymbol
	Modules   map[string]ModuleSymbol
	Id        string
}

func NewEnvironment(ant interface{}, ide string) Environment {
	return Environment{
		Anterior:  ant,
		Tabla:     make(map[string]Symbol),
		Structs:   make(map[string]Symbol),
		Functions: make(map[string]FunctionSymbol),
		Modules:   make(map[string]ModuleSymbol),
		Id:        ide,
	}
}

func (env Environment) SaveVariable(id string, value Symbol, tipo TipoExpresion, mut bool) {
	if variable, ok := env.Tabla[id]; ok {
		fmt.Println("La variable "+id+" ya existe", variable)
		return
	}
	env.Tabla[id] = value
}

func (env Environment) GetVariable(id string) Symbol {
	var tmpEnv Environment
	tmpEnv = env
	for {
		if variable, ok := tmpEnv.Tabla[id]; ok {
			return variable
		}
		if tmpEnv.Anterior == nil {
			break
		} else {
			tmpEnv = tmpEnv.Anterior.(Environment)
		}
	}
	fmt.Println("La variable no existe")
	return Symbol{Lin: 0, Col: 0, Id: "", Tipo: NULL, Valor: 0, Mutable: false}
}

func (env Environment) SetVariable(id string, value Symbol) Symbol {
	var tmpEnv Environment
	tmpEnv = env
	for {
		if variable, ok := tmpEnv.Tabla[id]; ok {
			if tmpEnv.Tabla[id].Mutable {
				if tmpEnv.Tabla[id].Tipo == value.Tipo {
					tmpEnv.Tabla[id] = value
					return variable
				} else {
					fmt.Println("El tipo es incorrecto")
					return Symbol{Lin: 0, Col: 0, Id: "", Tipo: NULL, Valor: 0, Mutable: false}
				}
			} else {
				fmt.Println("La variable no es mutable")
				return Symbol{Lin: 0, Col: 0, Id: "", Tipo: NULL, Valor: 0, Mutable: false}
			}
		}
		if tmpEnv.Anterior == nil {
			break
		} else {
			tmpEnv = tmpEnv.Anterior.(Environment)
		}
	}
	fmt.Println("La variable no existe")
	return Symbol{Lin: 0, Col: 0, Id: "", Tipo: NULL, Valor: 0, Mutable: false}
}

func (env Environment) SaveStruct(id string, list *arrayList.List, mut bool) {
	if structs, ok := env.Structs[id]; ok {
		fmt.Println("El struct " + structs.Id + " ya existe")
		return
	}
	env.Structs[id] = Symbol{Lin: 0, Col: 0, Id: id, Tipo: STRUCT, Valor: list, Mutable: mut}
}

func (env Environment) GetStruct(id string) Symbol {

	var tmpEnv Environment
	tmpEnv = env

	for {
		if tmpStruct, ok := tmpEnv.Structs[id]; ok {
			return tmpStruct
		}
		if tmpEnv.Anterior == nil {
			break
		} else {
			tmpEnv = tmpEnv.Anterior.(Environment)
		}
	}

	fmt.Println("El struct ", id, " no existe")
	return Symbol{Lin: 0, Col: 0, Id: "", Tipo: NULL, Valor: 0, Mutable: false}
}

func (env Environment) SetStruct(id *arrayList.List, value Symbol) Symbol {
	var tmpEnv Environment
	tmpEnv = env
	for {
		//creacion de diccionario temporal
		var tmpDic map[string]Symbol
		tmpDic = make(map[string]Symbol)
		//asignacion de diccionario
		tmpDic = env.Tabla
		//recorro la lista de id
		for _, s := range id.ToArray() { //recorremos lista
			//validando variable
			if variable, ok := tmpDic[s.(string)]; ok {
				if variable.Tipo == STRUCT {
					if variable.Mutable == true {
						tmpDic = variable.Valor.(map[string]Symbol)
					} else {
						fmt.Println("La variable no es mutable")
						return variable
					}
				} else if tmpDic[s.(string)].Mutable {
					tmpDic[s.(string)] = value
					return variable
				} else {
					fmt.Println("La variable no es mutable")
					return Symbol{Lin: 0, Col: 0, Id: "", Tipo: NULL, Valor: 0, Mutable: false}
				}
			}

		}
		if tmpEnv.Anterior == nil {
			break
		} else {
			tmpEnv = tmpEnv.Anterior.(Environment)
		}
	}
	fmt.Println("La variable no existe")
	return Symbol{Lin: 0, Col: 0, Id: "", Tipo: NULL, Valor: 0, Mutable: false}
}

func (env Environment) SaveFunction(id string, value FunctionSymbol) {
	if variable, ok := env.Functions[id]; ok {
		fmt.Println("La funcion " + variable.Id + " ya existe")
		return
	}
	env.Functions[id] = value
}

func (env Environment) GetFunction(id string) FunctionSymbol {
	fmt.Println("EN ENV:", id)
	var tmpEnv Environment
	tmpEnv = env
	for {
		if variable, ok := tmpEnv.Functions[id]; ok {
			return variable
		}
		if tmpEnv.Anterior == nil {
			break
		} else {
			tmpEnv = tmpEnv.Anterior.(Environment)
		}
	}
	fmt.Println("La funcion no existe")
	return FunctionSymbol{Ent: nil}
}

func (env Environment) LoopValidation() bool {
	var tmpEnv Environment
	tmpEnv = env
	for {
		if tmpEnv.Id == "WHILE" || tmpEnv.Id == "FOR-IN" || tmpEnv.Id == "LOOP" {
			return true
		}
		if tmpEnv.Anterior == nil {
			break
		} else {
			tmpEnv = tmpEnv.Anterior.(Environment)
		}
	}
	fmt.Println("la sentencia tiene que estar dentro de un ciclo")
	return false
}

func (env Environment) FuncValidation() bool {
	var tmpEnv Environment
	tmpEnv = env
	for {
		if strings.Contains(tmpEnv.Id, "FUNCTION") {
			return true
		}
		if tmpEnv.Anterior == nil {
			break
		} else {
			tmpEnv = tmpEnv.Anterior.(Environment)
		}
	}
	fmt.Println("la sentencia tiene que estar dentro de un ciclo")
	return false
}

func (env Environment) ReplaceVars(vars map[string]Symbol) {
	for key, element := range vars {
		fmt.Println(key, element)
	}
}

func (env Environment) SaveModule(id string, value ModuleSymbol) {
	if variable, ok := env.Modules[id]; ok {
		fmt.Println("El Modulo "+id+" ya existe", variable)
		return
	}
	env.Modules[id] = value
}

func (env Environment) GetModule(id string) ModuleSymbol {
	var tmpEnv Environment
	tmpEnv = env
	for {
		if variable, ok := tmpEnv.Modules[id]; ok {
			return variable
		}
		if tmpEnv.Anterior == nil {
			break
		} else {
			tmpEnv = tmpEnv.Anterior.(Environment)
		}
	}
	fmt.Println("El modulo ", id, " no existe")
	return ModuleSymbol{}
}
