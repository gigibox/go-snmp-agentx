package trap

import (
	"encoding/json"
	"fmt"
	"os/exec"

	"github.com/gosnmp/gosnmp"

	"go-snmp-agentx/oids"
)

type SMSDeviceMonitorModule struct {
	Device   string  `json:"device"`
	ChipTemp float64 `json:"chip_temp"`
	Modem    string  `json:"modem"`
	Signal   float64 `json:"signal"`
	RSRP     float64 `json:"rsrp"`
}

func (s *SMSDeviceMonitorModule) Name() string {
	return "SMSDeviceMonitorModule"
}

func (s *SMSDeviceMonitorModule) Check() ([]gosnmp.SnmpPDU, error) {
	var pdu = make([]gosnmp.SnmpPDU, 0)

	cmd := exec.Command("sh", "/usr/share/3ginfo-lite/modeminfo.sh")
	output, _ := cmd.CombinedOutput()
	_ = json.Unmarshal(output, s)

	if s.Device == "" {
		pdu = append(pdu, gosnmp.SnmpPDU{
			Value: `{"id":30102, "msg": "5G module not found"}`,
			Name:  oids.Trap5GNotPresent,
			Type:  gosnmp.OctetString,
		})
	} else {
		if s.Modem == "" {
			pdu = append(pdu, gosnmp.SnmpPDU{
				Value: `{"id":30101, "msg": "5G module hardware failure"}`,
				Name:  oids.Trap5GHardwareFailure,
				Type:  gosnmp.OctetString,
			})
		}

		if s.RSRP < -95 {
			pdu = append(pdu, gosnmp.SnmpPDU{
				Value: output,
				Name:  fmt.Sprintf(`{"id":40101, "msg": "5G Signal too weak. rsrp %f dBm"}`, s.RSRP),
				Type:  gosnmp.OctetString,
			})
		}

		if s.ChipTemp >= 75 {
			pdu = append(pdu, gosnmp.SnmpPDU{
				Value: fmt.Sprintf(`{"id":50103, "msg": "5G module temperature too high. %f ℃"}`, s.ChipTemp),
				Name:  oids.Trap5GHighTemp,
				Type:  0,
			})
		}

		if s.ChipTemp <= 35 {
			pdu = append(pdu, gosnmp.SnmpPDU{
				Value: fmt.Sprintf(`{"id":50104, "msg": "5G module temperature too low. %f ℃"}`, s.ChipTemp),
				Name:  oids.Trap5GLowTemp,
				Type:  0,
			})
		}
	}

	return pdu, nil
}
