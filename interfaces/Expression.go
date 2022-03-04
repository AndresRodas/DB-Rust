package interfaces

import "OLC2/environment"

type Expression interface {
	Ejecutar(ast, env interface{}) environment.Symbol
}
