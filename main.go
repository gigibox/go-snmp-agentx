package main

import (
	"flag"
	"time"

	"go-snmp-agentx/agentx"
	"go-snmp-agentx/agentx/value"
	"go-snmp-agentx/logger"
	"go-snmp-agentx/oids"
	"go-snmp-agentx/trap"
)

const retryInterval = 30

var socket, trapServer, community string
var trapPort, trapInterval int

func init() {
	flag.StringVar(&socket, "socket", "/var/run/agentx.sock", "snmpd agentx socket path")
	flag.StringVar(&trapServer, "trapServer", "", "trap server ip")
	flag.StringVar(&community, "community", "public", "snmp community")
	flag.IntVar(&trapPort, "trapPort", 162, "trap server port")
	flag.IntVar(&trapInterval, "trapInterval", 600, "trap interval in seconds")
}

func main() {
	flag.Parse()

	logger.Info("snmp agentx service start.")

	if trapServer != "" {
		trap.Init(trapServer, community, trapPort)
		go trap.SystemMonitorLoop(trapInterval)
		logger.Info("trap server: %s, port: %d, community:%s, interval: %d", trapServer, trapPort, community, trapInterval)

	}

	var client = new(agentx.Client)
	var err error

	for {
		client, err = agentx.Dial("unix", socket)
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
		if err := session.Register(127, value.MustParseOID(oids.StatOidGroup)); err != nil {
			logger.Error(err.Error())
			session.Close()

			time.Sleep(retryInterval * time.Second)

			continue
		}

		break
	}

	select {}
}
