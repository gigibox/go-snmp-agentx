package sysinfo

import (
	"encoding/json"

	"go-snmp-agentx/util"
)

func WirelessStat() interface{} {
	data, err := util.RunUbusCommand("call", "network.wireless", "status")
	if err != nil {
		return ""
	}

	var wireless = make(map[string]interface{})
	var result map[string]interface{}
	if err = json.Unmarshal(data, &result); err != nil {
		return ""
	}

	for radio, config := range result {
		radioInfo := config.(map[string]interface{})
		wireless[radio] = map[string]interface{}{
			"up":      radioInfo["up"],
			"band":    radioInfo["config"].(map[string]interface{})["band"],
			"channel": radioInfo["config"].(map[string]interface{})["channel"],
			"htmode":  radioInfo["config"].(map[string]interface{})["htmode"],
		}

		interfaces := radioInfo["interfaces"].([]interface{})
		for _, iface := range interfaces {
			ifaceInfo := iface.(map[string]interface{})
			wireless[radio] = map[string]interface{}{
				"ifname":     ifaceInfo["ifname"],
				"mode":       ifaceInfo["config"].(map[string]interface{})["mode"],
				"ssid":       ifaceInfo["config"].(map[string]interface{})["ssid"],
				"encryption": ifaceInfo["config"].(map[string]interface{})["encryption"],
				"stations":   WirelessClientCount(ifaceInfo["ifname"].(string)),
			}
		}
	}

	return util.Map2JSON(wireless)
}

func WirelessClientCount(ifname string) string {
	data, err := util.RunUbusCommand("iw", "dev", ifname, "station", "dump")
	if err != nil {
		return ""
	}

	return string(data)
}
