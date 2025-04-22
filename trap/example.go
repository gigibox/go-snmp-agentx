package trap

import (
	"github.com/gosnmp/gosnmp"

	"go-snmp-agentx/oids"
)

type ExampleModule struct {
}

func (e *ExampleModule) Name() string {
	return "ExampleModule"
}

func (e *ExampleModule) Check() ([]gosnmp.SnmpPDU, error) {
	pdu := []gosnmp.SnmpPDU{
		{
			Name:  oids.TrampExampleMessage,
			Type:  gosnmp.OctetString,
			Value: "Test trap message",
		},
	}

	return pdu, nil
}
