package trap

import (
	"fmt"
	"log"
	"time"

	"github.com/gosnmp/gosnmp"
)

func Send() {
	target := "10.1.1.227" // 替换为接收Trap的服务器地址
	port := 162            // SNMP Trap默认端口162
	community := "public"  // SNMP共同体名称，根据实际情况修改

	// 配置SNMP参数
	snmp := &gosnmp.GoSNMP{
		Target:    target,
		Port:      uint16(port),
		Version:   gosnmp.Version2c,
		Community: community,
		Timeout:   time.Duration(2) * time.Second,
	}

	// 连接到SNMP代理
	err := snmp.Connect()
	if err != nil {
		log.Fatalf("连接失败: %v", err)
	}
	defer snmp.Conn.Close()

	// 构造Trap的PDU变量
	pdus := []gosnmp.SnmpPDU{
		// sysUpTime实例，类型为TimeTicks（单位：百分之一秒）
		{
			Name:  ".1.3.6.1.2.1.1.3.0",
			Type:  gosnmp.TimeTicks,
			Value: uint32(1000), // 示例值，表示10秒
		},
		// snmpTrapOID，标识Trap类型
		{
			Name:  ".1.3.6.1.6.3.1.1.4.1.0",
			Type:  gosnmp.ObjectIdentifier,
			Value: ".1.3.6.1.6.3.1.1.5.1", // coldStart Trap示例
		},
		// 添加自定义信息（可选）
		{
			Name:  ".1.3.6.1.2.1.1.5.0", // sysName
			Type:  gosnmp.OctetString,
			Value: "MyDevice",
		},
	}

	// 创建SnmpTrap结构体
	trap := gosnmp.SnmpTrap{
		Variables: pdus,
	}

	for {

		// 发送Trap
		_, err = snmp.SendTrap(trap)
		if err != nil {
			log.Fatalf("发送Trap失败: %v", err)
		}

		time.Sleep(10 * time.Second)
	}
	fmt.Println("SNMP Trap发送成功！")
}
