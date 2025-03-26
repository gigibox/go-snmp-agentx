package sysinfo

import (
	"github.com/shirou/gopsutil/cpu"
)

func CPUStat() float64 {
	// 获取 CPU 使用率
	percent, err := cpu.Percent(0, false)
	if err != nil {
		return 0
	}

	return  percent[0]
}