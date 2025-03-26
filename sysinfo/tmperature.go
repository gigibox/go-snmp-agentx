package sysinfo

import (
	"fmt"

	"github.com/shirou/gopsutil/host"
)

func SensorsTemperatures () {
	sensors, err := host.SensorsTemperatures()
	if err != nil {
		return
	}

	for _, sensor := range sensors {
		fmt.Printf("传感器: %s, 温度: %.2f°C\n", sensor.SensorKey, sensor.Temperature)
	}
}
