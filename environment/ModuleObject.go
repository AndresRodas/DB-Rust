package environment

type ModuleObject struct {
	Visibilidad TipoModulo //pub o priv
	Tipo        TipoModulo //mod o inst
	Instruction interface{}
}

func NewModuleObject(vis TipoModulo, tip TipoModulo, inst interface{}) ModuleObject {
	exp := ModuleObject{vis, tip, inst}
	return exp
}
