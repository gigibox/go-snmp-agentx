package trap

import (
	"time"

	"github.com/gosnmp/gosnmp"

	"go-snmp-agentx/logger"
)

func SystemMonitorLoop(interval int) {
	if interval == 0 {
		interval = 60
	}

	for {
		var sendList = make([]gosnmp.SnmpPDU, 0)

		for _, m := range moduleList {
			pdu, err := m.Check()
			if err != nil {
				logger.Warn("Check %s error: %s", m.Name(), err.Error())
				continue
			}

			if pdu != nil {
				sendList = append(sendList, pdu...)
			}
		}

		if len(sendList) > 0 {
			Send(sendList)
		}

		time.Sleep(time.Duration(interval) * time.Second)
	}
}
