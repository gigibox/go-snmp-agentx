package oids

import (
	"go-snmp-agentx/agentx"
	"go-snmp-agentx/agentx/pdu"
	"go-snmp-agentx/sysinfo"
)

func InitOidHandler() *agentx.ListHandler {
	lh := &agentx.ListHandler{}

	// CPU 使用率
	item := lh.Add("1.3.6.1.4.1.45995.3.1")
	item.Type = pdu.VariableTypeOctetString
	item.Value = sysinfo.CPUStat

	// CPU/传感器温度
	item = lh.Add("1.3.6.1.4.1.45995.3.2")
	item.Type = pdu.VariableTypeOctetString
	item.Value = sysinfo.SensorsTemperatures

	// 内存使用率
	item = lh.Add("1.3.6.1.4.1.45995.3.3")
	item.Type = pdu.VariableTypeOctetString
	item.Value = sysinfo.MemoryStat

	// 无线网络状态
	item = lh.Add("1.3.6.1.4.1.45995.3.4")
	item.Type = pdu.VariableTypeOctetString
	item.Value = sysinfo.WirelessStat

	return lh
}
