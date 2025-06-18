package oids

import (
	"go-snmp-agentx/agentx"
	"go-snmp-agentx/agentx/pdu"
	"go-snmp-agentx/sysinfo"
)

func InitOidHandler() *agentx.ListHandler {
	lh := &agentx.ListHandler{}

	// 系统信息
	item := lh.Add(StatSystemInfo)
	item.Type = pdu.VariableTypeOctetString
	item.Value = sysinfo.GetBoardInfo

	// 系统启动时间
	item = lh.Add(StatSystemUptime)
	item.Type = pdu.VariableTypeOctetString
	item.Value = sysinfo.GetSysUpTime

	// 5G 信号强度
	item = lh.Add(Stat5GSignal)
	item.Type = pdu.VariableTypeOctetString
	item.Value = sysinfo.SMSSignal

	// 网络连接状态
	item = lh.Add(StatNetworkConnection)
	item.Type = pdu.VariableTypeOctetString
	item.Value = sysinfo.NetWorkDetect

	// 数据流量统计
	item = lh.Add(StatTrafficStatistics)
	item.Type = pdu.VariableTypeOctetString
	item.Value = sysinfo.TrafficStatistics

	// CPU 使用率
	item = lh.Add(StatCpuLoad)
	item.Type = pdu.VariableTypeOctetString
	item.Value = sysinfo.CPUStat

	// CPU/传感器温度
	item = lh.Add(StatSensorsTemperatures)
	item.Type = pdu.VariableTypeOctetString
	item.Value = sysinfo.SensorsTemperatures

	// 内存使用率
	item = lh.Add(StatMemoryUsage)
	item.Type = pdu.VariableTypeOctetString
	item.Value = sysinfo.MemoryStat

	// 无线网络状态
	item = lh.Add(StatWirelessStatus)
	item.Type = pdu.VariableTypeOctetString
	item.Value = sysinfo.WirelessStat

	// 设备状态
	item = lh.Add(StaDeviceStatus)
	item.Type = pdu.VariableTypeOctetString
	item.Value = sysinfo.GetDeviceStatus

	return lh
}
