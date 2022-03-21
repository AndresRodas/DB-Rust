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
	BREAK
	CONTINUE
	RETURN
	WILDCARD
)
