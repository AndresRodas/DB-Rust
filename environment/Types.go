package environment

type TipoExpresion int

const (
	INTEGER TipoExpresion = iota
	FLOAT
	STRING
	STR
	CHAR
	BOOLEAN
	ARRAY
	VECTOR
	STRUCT
	NULL
	WILDCARD
)

type TipoModulo int

const (
	PUBLIC TipoModulo = iota
	PRIVATE
	MOD
	INST
)
