package trap

import (
	"fmt"
	"strings"

	"github.com/gosnmp/gosnmp"
	"github.com/shirou/gopsutil/host"

	"go-snmp-agentx/logger"
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
				logWrite(logger.ErrorLevel, oids.TrapCpuHighTemp, fmt.Sprintf("50101 当前CPU温度过高,,将导致性能下降.当前温度: %f ℃", sensor.Temperature))

			} else if sensor.Temperature <= 18 {
				pdu = append(pdu, gosnmp.SnmpPDU{
					Value: fmt.Sprintf(`{"id":50102, "msg": "CPU temperature is too low, current temperature is %v"}`, sensor.Temperature),
					Name:  oids.TrapCpuLowTemp,
					Type:  gosnmp.OctetString,
				})

				logWrite(logger.ErrorLevel, oids.TrapCpuLowTemp, fmt.Sprintf("50102 当前CPU温度过低,,将导致性能下降.当前温度: %f ℃", sensor.Temperature))
			} else {
				logWrite(logger.DebugLevel, oids.TrapCpuHighTemp, fmt.Sprintf("50101 当前CPU温度正常,当前温度: %f ℃", sensor.Temperature))
				logWrite(logger.DebugLevel, oids.TrapCpuLowTemp, fmt.Sprintf("50102 当前CPU温度正常,当前温度: %f ℃", sensor.Temperature))
			}
		}
	}

	return pdu, nil
}
