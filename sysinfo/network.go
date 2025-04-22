package sysinfo

import (
	"encoding/json"
	"os/exec"
	"time"

	"go-snmp-agentx/util"
)

func NetWorkDetect() interface{} {
	var netStat = make(map[string]interface{})

	data, err := util.RunUbusCommand("call", "luci.internet-detector", "InetStatus")
	var result struct {
		Instances []struct {
			Inet        int64  `json:"inet"`
			ModPublicIp string `json:"mod_public_ip"`
		} `json:"instances"`
	}

	err = json.Unmarshal(data, &result)
	if err != nil && len(result.Instances) == 0 {
		return "{}"
	}

	netStat["connected"] = false
	if result.Instances[0].Inet == 0 {
		netStat["connected"] = true
		netStat["public_ip"] = result.Instances[0].ModPublicIp
		netStat["timestamp"] = time.Now().Unix()
	}

	return util.Map2JSON(netStat)
}

func TrafficStatistics() interface{} {
	data, err := util.RunUbusCommand("call", "luci-rpc", "getNetworkDevices")
	if err != nil {
		return ""
	}

	var devices = make(map[string]interface{})
	var result map[string]interface{}
	if err = json.Unmarshal(data, &result); err != nil {
		return ""
	}

	for name, dev := range result {
		// 过滤掉不需要的设备
		devInfo := dev.(map[string]interface{})
		if !devInfo["up"].(bool) ||
			devInfo["wireless"].(bool) ||
			devInfo["name"].(string) == "br-lan" ||
			devInfo["name"].(string) == "lo" {
			continue
		}

		devices[name] = map[string]interface{}{
			"rx_bytes":  devInfo["stats"].(map[string]interface{})["rx_bytes"],
			"tx_bytes":  devInfo["stats"].(map[string]interface{})["tx_bytes"],
			"timestamp": time.Now().Unix(),
		}
	}

	return util.Map2JSON(devices)
}

func SMSSignal() interface{} {
	cmd := exec.Command("sh", "/usr/share/3ginfo-lite/modeminfo.sh")
	output, _ := cmd.CombinedOutput()
	return string(output)
}
