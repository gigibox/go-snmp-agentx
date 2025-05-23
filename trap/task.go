package trap

import (
	"time"

	"github.com/gosnmp/gosnmp"
	"github.com/patrickmn/go-cache"

	"go-snmp-agentx/logger"
)

var IdsTrapCache *cache.Cache
var IdsInTrapCache = make(map[string]bool)

func SystemMonitorLoop(checkInterval, trapInterval int) {
	if checkInterval == 0 {
		checkInterval = 60
	}

	IdsTrapCache = cache.New(time.Duration(trapInterval)*time.Second, 5*time.Minute)

	for {
		var pduList = make([]gosnmp.SnmpPDU, 0)
		var sendList = make([]gosnmp.SnmpPDU, 0)

		for _, m := range moduleList {
			pdu, err := m.Check()
			if err != nil {
				logger.Warn("Check %s error: %s", m.Name(), err.Error())
				continue
			}

			if pdu != nil {
				pduList = append(pduList, pdu...)
			}
		}

		for _, pdu := range pduList {
			IdsInTrapCache[pdu.Name] = true

			if pdu.Value == nil {
				continue
			}

			if _, found := IdsTrapCache.Get(pdu.Name); found {
				continue
			}

			IdsTrapCache.Set(pdu.Name, true, cache.DefaultExpiration)
			sendList = append(sendList, pdu)
		}

		if len(sendList) > 0 {
			Send(sendList)
		}

		time.Sleep(time.Duration(checkInterval) * time.Second)
	}
}

func logWrite(level int, oid, msg string) {
	if level == logger.DebugLevel {
		if _, ok := IdsInTrapCache[oid]; ok {
			logger.Debug(msg)
			delete(IdsInTrapCache, oid)
		}

		return
	}

	if _, found := IdsTrapCache.Get(oid); found {
		return
	}

	if level == logger.InfoLevel {
		logger.Info(msg)
	} else if level == logger.WarnLevel {
		logger.Warn(msg)
	} else if level == logger.ErrorLevel {
		logger.Error(msg)
	}
}
