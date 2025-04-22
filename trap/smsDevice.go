package trap

import (
	"encoding/json"
	"os/exec"

	"github.com/gosnmp/gosnmp"

	"go-snmp-agentx/oids"
)

type SMSDeviceMonitorModule struct {
	Device   string  `json:"device"`
	ChipTemp float64 `json:"chip_temp"`
	Modem    string  `json:"modem"`
	Signal   float64 `json:"signal"`
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
			Value: `{"msg": "5G device not found"}`,
			Name:  oids.Trap5GNotPresent,
			Type:  gosnmp.OctetString,
		})
	} else {
		if s.Modem == "" {
			pdu = append(pdu, gosnmp.SnmpPDU{
				Value: output,
				Name:  oids.Trap5GHardwareFailure,
				Type:  gosnmp.OctetString,
			})
		}

		if s.Signal < 20 {
			pdu = append(pdu, gosnmp.SnmpPDU{
				Value: output,
				Name:  oids.Trap5GSignalTooWeak,
				Type:  gosnmp.OctetString,
			})
		}

		if s.ChipTemp > 70 {
			pdu = append(pdu, gosnmp.SnmpPDU{
				Value: output,
				Name:  oids.Trap5GHighTemp,
				Type:  0,
			})
		}

		if s.ChipTemp < 10 {
			pdu = append(pdu, gosnmp.SnmpPDU{
				Value: output,
				Name:  oids.Trap5GLowTemp,
				Type:  0,
			})
		}
	}

	return pdu, nil
}
