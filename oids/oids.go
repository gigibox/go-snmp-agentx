package oids

const (
	StatSystemInfo        = "1.3.6.1.2.1.1.1.0" // 系统描述，涵盖设备厂商、型号、软件版本等信息
	StatSystemUptime      = "1.3.6.1.2.1.1.2.0" // 系统正常运行时间，即设备自上次启动后的运行时长
	Stat5GSignal          = "1.3.6.1.2.1.1.3.0" // 5G 信号强度，反映设备接收到的 5G 信号强弱
	StatNetworkConnection = "1.3.6.1.2.1.1.4.0" // 网络连接状态，如已连接、断开连接等
	StatTrafficStatistics = "1.3.6.1.2.1.1.5.0" // 数据流量统计，统计设备的上下行数据流量

	// Trap OIDs
	Trap5GHardwareFailure      = "1.3.7.1.2.1.1.1.0"  // 5G模组硬件故障
	Trap5GNotPresent           = "1.3.7.1.2.1.1.2.0"  // 5G模组不在位告警
	TrapWiFiHardwareFailure    = "1.3.7.1.2.1.1.3.0"  // WIFI模组硬件故障告警
	TrapStorageHardwareFailure = "1.3.7.1.2.1.1.4.0"  // 存储器硬件故障告警
	Trap5GSignalTooWeak        = "1.3.7.1.2.1.1.5.0"  // 5G信号接收强度过低
	TrapAntennaImbalance       = "1.3.7.1.2.1.1.6.0"  // 天线接收性能不平衡告警
	TrapLowStorageSpace        = "1.3.7.1.2.1.1.7.0"  // 存储容量不足告警
	TrapGPSSignalTooWeak       = "1.3.7.1.2.1.1.8.0"  // BD/GPS收星不足
	TrapCpuHighTemp            = "1.3.7.1.2.1.1.9.0"  // CPU高温告警
	TrapCpuLowTemp             = "1.3.7.1.2.1.1.10.0" // CPU低温告警
	Trap5GHighTemp             = "1.3.7.1.2.1.1.11.0" // 5G模组高温告警
	Trap5GLowTemp              = "1.3.7.1.2.1.1.12.0" // 5G模组低温告警
	TrapPowerUnderVoltage      = "1.3.7.1.2.1.1.13.0" // 供电电压不足告警
	TrapPowerOverVoltage       = "1.3.7.1.2.1.1.14.0" // 供电电压过高告警
)
