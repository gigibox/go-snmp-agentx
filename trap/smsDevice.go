package trap

import (
	"encoding/json"
	"fmt"
	"os/exec"

	"github.com/gosnmp/gosnmp"

	"go-snmp-agentx/logger"
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

func (s *SMSDeviceMonitorModule) Clean() {
	s.Device = ""
	s.ChipTemp = 0
	s.Modem = ""
	s.Signal = 0
	s.RSRP = 0
}

func (s *SMSDeviceMonitorModule) Check() ([]gosnmp.SnmpPDU, error) {
	var pdu = make([]gosnmp.SnmpPDU, 0)
	if _, err := exec.LookPath("/usr/share/3ginfo-lite/modeminfo.sh"); err != nil {
		return pdu, nil
	}

	fmt.Println("SMSDeviceMonitorModule Check")
	cmd := exec.Command("sh", "/usr/share/3ginfo-lite/modeminfo.sh")

	output, _ := cmd.CombinedOutput()

	s.Clean()

	err := json.Unmarshal(output, s)
	if err != nil {
		s.Device = ""
		fmt.Println("unmarshal modem info error.", err.Error())
	}

	fmt.Printf("SMSDeviceMonitorModule info %+v\n", s)

	if s.Device == "" {
		pdu = append(pdu, gosnmp.SnmpPDU{
			Value: PackageTrapMessage(30102,"严重", fmt.Sprintf("5G模组不在位,请重新安装!")),
			Name:  oids.Trap5GNotPresent,
			Type:  gosnmp.OctetString,
		})
		logWrite(logger.ErrorLevel, oids.Trap5GNotPresent, "30102 5G模组不在位,请重新安装!")
	} else {
		logWrite(logger.DebugLevel, oids.Trap5GNotPresent, "30102 5G模组已在位.")

		if s.Modem == "" {
			pdu = append(pdu, gosnmp.SnmpPDU{
				Value: PackageTrapMessage(30101,"严重", fmt.Sprintf("5G模组硬件故障,请更换硬件!")),
				Name:  oids.Trap5GHardwareFailure,
				Type:  gosnmp.OctetString,
			})

			logWrite(logger.ErrorLevel, oids.Trap5GHardwareFailure, "30101 5G模组硬件故障,请更换硬件!")
			return pdu, nil
		} else {
			logWrite(logger.DebugLevel, oids.Trap5GHardwareFailure, "30101 5G模组硬件正常.")
		}

		if s.RSRP < -95 {
			pdu = append(pdu, gosnmp.SnmpPDU{
				Value: PackageTrapMessage(40101,"一般", fmt.Sprintf("当前5G信号过低,请检查天线,信号强度: %f dBm", s.RSRP)),
				Name:  oids.Trap5GSignalTooWeak,
				Type:  gosnmp.OctetString,
			})

			logWrite(logger.WarnLevel, oids.Trap5GSignalTooWeak, fmt.Sprintf("40101 当前5G信号过低,请检查天线,信号强度: %f dBm", s.RSRP))
		} else {
			logWrite(logger.DebugLevel, oids.Trap5GSignalTooWeak, fmt.Sprintf("40101 5G模组信号恢复正常.信号强度: %f dBm", s.RSRP))
		}

		if s.ChipTemp >= 75 {
			pdu = append(pdu, gosnmp.SnmpPDU{
				Value: PackageTrapMessage(50103,"严重", fmt.Sprintf("当前5G通信模块温度过高,将导致通信性能下降.当前温度: %.2f ℃", s.ChipTemp)),
				Name:  oids.Trap5GHighTemp,
				Type:  0,
			})

			logWrite(logger.ErrorLevel, oids.Trap5GHighTemp, fmt.Sprintf("50103 当前5G通信模块温度过高,将导致通信性能下降.当前温度: %.2f ℃", s.ChipTemp))
		} else if s.ChipTemp <= 35 {
			pdu = append(pdu, gosnmp.SnmpPDU{
				Value: PackageTrapMessage(50104,"严重", fmt.Sprintf("当前5G通信模块温度过低,将导致通信性能下降.当前温度: %.2f ℃", s.ChipTemp)),
				Name:  oids.Trap5GLowTemp,
				Type:  0,
			})

			logWrite(logger.ErrorLevel, oids.Trap5GLowTemp, fmt.Sprintf("50104 当前5G通信模块温度过低,将导致通信性能下降.当前温度: %.2f ℃", s.ChipTemp))
		} else {
			logWrite(logger.DebugLevel, oids.Trap5GHighTemp, fmt.Sprintf("50103 5G模组温度恢复正常.当前温度: %.2f ℃", s.ChipTemp))
			logWrite(logger.DebugLevel, oids.Trap5GLowTemp, fmt.Sprintf("50104 5G模组温度恢复正常.当前温度: %.2f ℃", s.ChipTemp))
		}
	}

	return pdu, nil
}
