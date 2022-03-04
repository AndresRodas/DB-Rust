package environment

type Environment struct {
	Lin           int
	Col           int
	Anterior      interface{}
	Tabla         map[string]Symbol
	Identificador string
}
