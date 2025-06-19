package oids

const (
	StatOidGroup            = "1.3.6.1.4.1.5688.0"
	StatSystemInfo          = "1.3.6.1.4.1.5688.0.1"  // 系统描述，涵盖设备厂商、型号、软件版本等信息
	StatSystemUptime        = "1.3.6.1.4.1.5688.0.2"  // 系统正常运行时间，即设备自上次启动后的运行时长
	Stat5GSignal            = "1.3.6.1.4.1.5688.0.3"  // 5G 信号强度，反映设备接收到的 5G 信号强弱
	StatNetworkConnection   = "1.3.6.1.4.1.5688.0.4"  // 网络连接状态，如已连接、断开连接等
	StatTrafficStatistics   = "1.3.6.1.4.1.5688.0.5"  // 数据流量统计，统计设备的上下行数据流量
	StatCpuLoad             = "1.3.6.1.4.1.5688.0.6"  // CPU 使用率
	StatSensorsTemperatures = "1.3.6.1.4.1.5688.0.7"  // CPU/传感器温度
	StatMemoryUsage         = "1.3.6.1.4.1.5688.0.8"  // 内存使用率
	StatWirelessStatus      = "1.3.6.1.4.1.5688.0.9"  // 无线网络状态
	StaDeviceStatus         = "1.3.6.1.4.1.5688.0.10" // 系统状态

	// Trap OIDs
	TrampExampleMessage        = "1.3.7.1.2.1.1.0.0"  // 测试程序
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
