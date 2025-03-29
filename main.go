package main

import (
	"time"

	"go-snmp-agentx/agentx"
	"go-snmp-agentx/agentx/pdu"
	"go-snmp-agentx/agentx/value"
	"go-snmp-agentx/logger"
	"go-snmp-agentx/sysinfo"
)

const retryInterval = 30

func main() {
	logger.Info("snmp agentx service start.")

	var client = new(agentx.Client)
	var err error

	for {
		client, err = agentx.Dial("unix", "/var/run/agentx.sock")
		if err != nil {
			time.Sleep(retryInterval * time.Second)
			continue
		}

		logger.Info("snmpd connection successful.")
		break
	}

	client.Timeout = 1 * time.Minute
	client.ReconnectInterval = 1 * time.Second

	listHandler := &agentx.ListHandler{}

	// CPU 使用率
	item := listHandler.Add("1.3.6.1.4.1.45995.3.1")
	item.Type = pdu.VariableTypeOctetString
	item.Value = sysinfo.CPUStat()

	// CPU/传感器温度
	item = listHandler.Add("1.3.6.1.4.1.45995.3.2")
	item.Type = pdu.VariableTypeOctetString
	item.Value = sysinfo.SensorsTemperatures()

	// 内存使用率
	item = listHandler.Add("1.3.6.1.4.1.45995.3.3")
	item.Type = pdu.VariableTypeOctetString
	item.Value = sysinfo.MemoryStat()

	// 无线网络状态
	item = listHandler.Add("1.3.6.1.4.1.45995.3.4")
	item.Type = pdu.VariableTypeOctetString
	item.Value = sysinfo.WirelessStat()

	for {
		session, err := client.Session()
		if err != nil {
			logger.Error(err.Error())
			time.Sleep(retryInterval * time.Second)

			continue
		}

		session.Handler = listHandler

		if err := session.Register(127, value.MustParseOID("1.3.6.1.4.1.45995.3")); err != nil {
			logger.Error(err.Error())
			session.Close()
			time.Sleep(retryInterval * time.Second)

			continue
		}

		break
	}

	select {}
}
