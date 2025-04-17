package trap

import (
	"time"

	"go-snmp-agentx/logger"
)

func SystemMonitorLoop(interval int) {
	if interval == 0 {
		interval = 60
	}

	for {
		sendList := make(map[string]string)
		for _, m := range moduleList {
			msg, err := m.Check()
			if err != nil {
				logger.Warn("Check %s error: %s", m.Name(), err.Error())
				continue
			}

			if msg != "" {
				logger.Warn("Check %s: %s", m.Name(), msg)

				sendList[m.OID()] = msg
			}
		}

		if len(sendList) > 0 {
			Send(sendList)
		}

		time.Sleep(time.Duration(interval) * time.Second)
	}
}
