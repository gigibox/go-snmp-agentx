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

var socket, trapServer, community, version string
var trapPort, trapInterval, checkInterval int

func init() {
	flag.StringVar(&socket, "socket", "/var/run/agentx.sock", "snmpd agentx socket path")
	flag.StringVar(&trapServer, "trapServer", "", "trap server ip")
	flag.StringVar(&community, "community", "public", "snmp community")
	flag.IntVar(&trapPort, "trapPort", 162, "trap server port")
	flag.IntVar(&trapInterval, "trapInterval", 600, "trap interval in seconds")
	flag.IntVar(&checkInterval, "checkInterval", 60, "check interval in seconds")
	flag.StringVar(&version, "v", "1.6.4", "version")
}

func main() {
	flag.Parse()

	logger.Debug("snmp agentx service start.")

	if trapServer != "" {
		trap.Init(trapServer, community, trapPort)
		go trap.SystemMonitorLoop(checkInterval, trapInterval)
	}

	var client = new(agentx.Client)
	var err error

	for {
		client, err = agentx.Dial("unix", socket)
		if err != nil {
			time.Sleep(retryInterval * time.Second)
			continue
		}

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
