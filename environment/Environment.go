package environment

import (
	"fmt"
	arrayList "github.com/colegno/arraylist"
)

type Environment struct {
	Anterior interface{}
	Tabla    map[string]Symbol
	Structs  map[string]Symbol
	Id       string
}

func NewEnvironment(ant interface{}, id string) Environment {
	return Environment{ant, make(map[string]Symbol), make(map[string]Symbol), id}
}

func (env Environment) SaveVariable(id string, value Symbol, tipo TipoExpresion, mut bool) {
	if variable, ok := env.Tabla[id]; ok {
		fmt.Println("La variable " + variable.Id + " ya existe")
		return
	}
	env.Tabla[id] = value
	//env.Tabla[id] = Symbol{Lin: 0, Col: 0, Id: id, Tipo: tipo, Valor: value, Mutable: mut}
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
				tmpEnv.Tabla[id] = value
				//tmpEnv.Tabla[id] = Symbol{Lin: 0, Col: 0, Id: id, Tipo: variable.Tipo, Valor: value, Mutable: true}
				return variable
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
	//fmt.Println(env.Structs[id].Valor)
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

	fmt.Println("El struct no existe")
	return Symbol{Lin: 0, Col: 0, Id: "", Tipo: NULL, Valor: 0, Mutable: false}
}
