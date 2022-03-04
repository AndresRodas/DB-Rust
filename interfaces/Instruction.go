package interfaces

type Instruction interface {
	Ejecutar(ast, env interface{}) interface{}
}
