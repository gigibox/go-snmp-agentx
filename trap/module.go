package trap

type Module interface {
	Name() string
	Check() (string, error)
	OID() string
}

var moduleList = []Module{
	&ExampleModule{},
}
