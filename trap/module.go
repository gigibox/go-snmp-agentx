package trap

import "github.com/gosnmp/gosnmp"

type Module interface {
	Name() string
	Check() ([]gosnmp.SnmpPDU, error)
}

var moduleList = []Module{
	// &ExampleModule{},
	&SMSDeviceMonitorModule{},
	&CpuMonitorModule{},
	&WirelessModule{},
	&DiskModule{},
}
