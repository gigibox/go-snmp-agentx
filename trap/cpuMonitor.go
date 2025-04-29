package trap

import (
	"fmt"
	"strings"

	"github.com/gosnmp/gosnmp"
	"github.com/shirou/gopsutil/host"

	"go-snmp-agentx/oids"
)

type CpuMonitorModule struct {
}

func (c *CpuMonitorModule) Name() string {
	return "CpuMonitorModule"
}

func (c *CpuMonitorModule) Check() ([]gosnmp.SnmpPDU, error) {
	sensors, err := host.SensorsTemperatures()
	if err != nil {
		return nil, err
	}

	var pdu = make([]gosnmp.SnmpPDU, 0)

	for _, sensor := range sensors {
		if strings.Contains(sensor.SensorKey, "cpu") {
			if sensor.Temperature >= 60 {
				pdu = append(pdu, gosnmp.SnmpPDU{
					Value: fmt.Sprintf(`{"id":50101, "msg": "CPU temperature is too high, current temperature is %v"}`, sensor.Temperature),
					Name:  oids.TrapCpuHighTemp,
					Type:  gosnmp.OctetString,
				})
			} else if sensor.Temperature <= 18 {
				pdu = append(pdu, gosnmp.SnmpPDU{
					Value: fmt.Sprintf(`{"id":50102, "msg": "CPU temperature is too low, current temperature is %v"}`, sensor.Temperature),
					Name:  oids.TrapCpuLowTemp,
					Type:  gosnmp.OctetString,
				})
			}
		}
	}

	return pdu, nil
}
