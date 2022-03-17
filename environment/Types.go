package environment

type TipoExpresion int

const (
	INTEGER TipoExpresion = iota
	FLOAT
	STRING
	BOOLEAN
	ARRAY
	NULL
	STRUCT
	CHAR
	VECTOR
	WILDCARD
	BREAK
	CONTINUE
)
