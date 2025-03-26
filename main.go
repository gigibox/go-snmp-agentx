package main

import (
	"log"
	"net"
	"time"

	"go-snmp-agentx/agentx"
	"go-snmp-agentx/agentx/pdu"
	"go-snmp-agentx/agentx/value"
)

func main() {
	client, err := agentx.Dial("tcp", "localhost:705")
	if err != nil {
		log.Fatalf(err.Error())
	}
	client.Timeout = 1 * time.Minute
	client.ReconnectInterval = 1 * time.Second

	session, err := client.Session()
	if err != nil {
		log.Fatalf(err.Error())
	}

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

	session.Handler = listHandler

	if err := session.Register(127, value.MustParseOID("1.3.6.1.4.1.45995.3")); err != nil {
		log.Fatalf(err.Error())
	}

	for {
		time.Sleep(100 * time.Millisecond)
	}
}
