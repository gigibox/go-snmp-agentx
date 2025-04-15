package sysinfo

import (
	"github.com/shirou/gopsutil/host"

	"go-snmp-agentx/util"
)

func SensorsTemperatures() interface{} {
	var result = make(map[string]interface{})

	sensors, err := host.SensorsTemperatures()
	if err != nil {
		return ""
	}

	for _, sensor := range sensors {
		result[sensor.SensorKey] = sensor.Temperature
	}

	return util.Map2JSON(result)
}
