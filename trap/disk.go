package trap

import (
	"fmt"

	"github.com/gosnmp/gosnmp"
	"github.com/shirou/gopsutil/disk"

	"go-snmp-agentx/oids"
)

type DiskModule struct {
}

func (d *DiskModule) Name() string {
	return "DiskModule"
}

func (d *DiskModule) Check() ([]gosnmp.SnmpPDU, error) {
	var pdu = make([]gosnmp.SnmpPDU, 0)
	// 获取磁盘使用情况
	usage, err := disk.Usage("/")
	if err != nil {
		return nil, err
	}

	if usage.UsedPercent >= 90 {
		pdu = append(pdu, gosnmp.SnmpPDU{
			Value: fmt.Sprintf(`{"id":40103, "msg": "Disk usage is too high, current usage is %.2f%%"}`, usage.UsedPercent),
			Name:  oids.TrapLowStorageSpace,
			Type:  gosnmp.OctetString,
		})
	}

	return pdu, nil
}
