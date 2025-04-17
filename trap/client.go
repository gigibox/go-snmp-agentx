package trap

import (
	"log"
	"time"

	"github.com/gosnmp/gosnmp"
)

var serverAddr, community string // 接收Trap的服务器地址和SNMP共同体名称，根据实际情况修改
var serverPort int

func Init(targetIP, communityName string, port int) {
	serverAddr = targetIP
	community = communityName
	serverPort = port

	if serverPort == 0 {
		serverPort = 162
	}

	if community == "" {
		community = "public"
	}
}

func Send(msgMap map[string]string) {

	// 配置SNMP参数
	snmp := &gosnmp.GoSNMP{
		Target:    serverAddr,
		Port:      uint16(serverPort),
		Version:   gosnmp.Version2c,
		Community: community,
		Timeout:   time.Duration(3) * time.Second,
	}

	// 连接到SNMP代理
	err := snmp.Connect()
	if err != nil {
		log.Fatalf("连接失败: %v", err)
	}
	defer snmp.Conn.Close()

	// 构造Trap的PDU变量
	var pdus []gosnmp.SnmpPDU

	for k, v := range msgMap {
		pdus = append(pdus, gosnmp.SnmpPDU{
			Name:  k,
			Type:  gosnmp.OctetString,
			Value: v,
		})
	}

	// 创建SnmpTrap结构体
	trap := gosnmp.SnmpTrap{
		Variables: pdus,
	}

	_, err = snmp.SendTrap(trap)
	if err != nil {
		log.Fatalf("发送Trap失败: %s", err.Error())
	}
}
