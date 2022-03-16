package interfaces

import "OLC2/environment"

type Instruction interface {
	Ejecutar(ast *environment.AST, env interface{}) environment.Symbol
}
