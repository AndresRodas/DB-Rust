package environment

type Environment struct {
	Anterior interface{}
	Tabla    map[string]Symbol
	Id       string
}

func NewEnvironment(ant interface{}, id string) Environment {
	return Environment{ant, make(map[string]Symbol), id}
}
