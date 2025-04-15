package main

import (
	"time"

	"go-snmp-agentx/agentx"
	"go-snmp-agentx/agentx/value"
	"go-snmp-agentx/logger"
	"go-snmp-agentx/oids"
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

	for {
		session, err := client.Session()
		if err != nil {
			logger.Error(err.Error())
			time.Sleep(retryInterval * time.Second)

			continue
		}

		session.Handler = oids.InitOidHandler()
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
