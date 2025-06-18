package trap

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"

	"github.com/gosnmp/gosnmp"

	"go-snmp-agentx/logger"
	"go-snmp-agentx/oids"
)

type WirelessModule struct {
}

func (w *WirelessModule) Name() string {
	return "WirelessModule"
}

func (w *WirelessModule) Check() ([]gosnmp.SnmpPDU, error) {
	var pdu = make([]gosnmp.SnmpPDU, 0)

	cmd := exec.Command("iw", "dev")
	var out bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &out
	cmd.Stderr = &stderr

	err := cmd.Run()
	if err != nil {
		logger.Warn("执行 'iw dev' 失败: %v, 错误信息: %s", err, stderr.String())
		return nil, nil
	}

	output := out.String()
	if strings.Contains(output, "Interface") {
		logWrite(logger.DebugLevel, oids.TrapWiFiHardwareFailure, "30103 WIFI模组硬件正常.")
		return nil, nil
	}

	pdu = append(pdu, gosnmp.SnmpPDU{
		Value: PackageTrapMessage(30103,"严重", fmt.Sprintf("WIFI模组硬件故障,请更换硬件!")),
		Name:  oids.TrapWiFiHardwareFailure,
		Type:  gosnmp.OctetString,
	})

	logWrite(logger.ErrorLevel, oids.TrapWiFiHardwareFailure, "30103 WIFI模组硬件故障,请更换硬件!")
	return pdu, nil
}
