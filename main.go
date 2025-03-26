package main

import (
	"net"
	"time"

	"go-snmp-agentx/agentx"
	"go-snmp-agentx/agentx/pdu"
	"go-snmp-agentx/agentx/value"
	"go-snmp-agentx/logger"
)

const retryInterval = 30

func main() {
	logger.Info("snmp agentx service start.")

	var client  = new (agentx.Client)
	var err error

	for {
		client, err = agentx.Dial("unix", "/var/run/agentx.sock")
		if err != nil {
			time.Sleep( retryInterval * time.Second)
			continue
		}

		logger.Info("snmpd connection successful.")
		break
	}

	client.Timeout = 1 * time.Minute
	client.ReconnectInterval = 1 * time.Second

	listHandler := &agentx.ListHandler{}

	item := listHandler.Add("1.3.6.1.4.1.45995.3.1")
	item.Type = pdu.VariableTypeInteger
	item.Value = int32(-123)

	item = listHandler.Add("1.3.6.1.4.1.45995.3.2")
	item.Type = pdu.VariableTypeOctetString
	item.Value = "echo test"

	item = listHandler.Add("1.3.6.1.4.1.45995.3.3")
	item.Type = pdu.VariableTypeNull
	item.Value = nil

	item = listHandler.Add("1.3.6.1.4.1.45995.3.4")
	item.Type = pdu.VariableTypeObjectIdentifier
	item.Value = "1.3.6.1.4.1.45995.1.5"

	item = listHandler.Add("1.3.6.1.4.1.45995.3.5")
	item.Type = pdu.VariableTypeIPAddress
	item.Value = net.IP{10, 10, 10, 10}

	item = listHandler.Add("1.3.6.1.4.1.45995.3.6")
	item.Type = pdu.VariableTypeCounter32
	item.Value = uint32(123)

	item = listHandler.Add("1.3.6.1.4.1.45995.3.7")
	item.Type = pdu.VariableTypeGauge32
	item.Value = uint32(123)

	item = listHandler.Add("1.3.6.1.4.1.45995.3.8")
	item.Type = pdu.VariableTypeTimeTicks
	item.Value = 123 * time.Second

	item = listHandler.Add("1.3.6.1.4.1.45995.3.9")
	item.Type = pdu.VariableTypeOpaque
	item.Value = []byte{1, 2, 3}

	item = listHandler.Add("1.3.6.1.4.1.45995.3.10")
	item.Type = pdu.VariableTypeCounter64
	item.Value = uint64(12345678901234567890)

	for {
		session, err := client.Session()
		if err != nil {
			logger.Error(err.Error())
			time.Sleep( retryInterval * time.Second)

			continue
		}

		session.Handler = listHandler

		if err := session.Register(127, value.MustParseOID("1.3.6.1.4.1.45995.3")); err != nil {
			logger.Error(err.Error())
			session.Close()
			time.Sleep( retryInterval * time.Second)

			continue
		}

		break
	}

	select {}
}
