package trap

import "go-snmp-agentx/oids"

type ExampleModule struct {
}

func (e *ExampleModule) Name() string {
	return "ExampleModule"
}

func (e *ExampleModule) OID() string {
	return oids.TrampExampleMessage
}

func (e *ExampleModule) Check() (string, error) {
	return "Test trap message.", nil
}
