package trap

import (
	"bytes"
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
		return nil, nil
	}

	pdu = append(pdu, gosnmp.SnmpPDU{
		Value: `{"msg": "Wireless device not found"}`,
		Name:  oids.TrapWiFiHardwareFailure,
		Type:  gosnmp.OctetString,
	})

	return pdu, nil
}
