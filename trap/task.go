package trap

import (
	"fmt"
	"time"

	"go-snmp-agentx/oids"
	"go-snmp-agentx/util"

	"github.com/gosnmp/gosnmp"
	"github.com/patrickmn/go-cache"

	"go-snmp-agentx/logger"
)

var IdsTrapCache *cache.Cache
var IdsLogDebugCache = make(map[string]bool)

func SystemMonitorLoop(checkInterval, trapInterval int) {
	if checkInterval == 0 {
		checkInterval = 60
	}

	util.DeviceSN, _ = util.GetBrLanMAC()

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
			IdsLogDebugCache[pdu.Name] = true

			if pdu.Value == nil {
				continue
			}

			if _, found := IdsTrapCache.Get(pdu.Name); found {
				continue
			}

			IdsTrapCache.Set(pdu.Name, true, cache.DefaultExpiration)
			sendList = append(sendList, pdu)
		}

		util.DeviceStatus = "normal"
		if len(sendList) > 0 {
			util.DeviceStatus = "alarm"
		} else {
			sendList = append(sendList, gosnmp.SnmpPDU{
				Value: PackageTrapMessage(0, "正常", "设备状态正常."),
				Name:  oids.TrapDeviceStatus,
				Type:  gosnmp.OctetString,
			})
		}

		Send(sendList)

		time.Sleep(time.Duration(checkInterval) * time.Second)
	}
}

func logWrite(level int, oid, msg string) {
	fmt.Println(level, oid, msg)
	if level == logger.DebugLevel {
		if _, ok := IdsLogDebugCache[oid]; ok {
			logger.Debug(msg)
			fmt.Println("delete cache ", oid)
			delete(IdsLogDebugCache, oid)
			IdsTrapCache.Delete(oid)
		}

		return
	}

	if _, found := IdsTrapCache.Get(oid); found {
		fmt.Println("oid found in cache.", oid)
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
