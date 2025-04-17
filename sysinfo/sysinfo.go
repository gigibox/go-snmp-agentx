package sysinfo

import (
	"encoding/json"
	"fmt"

	"go-snmp-agentx/util"
)

func GetBoardInfo() interface{} {
	data, err := util.RunUbusCommand("call", "system", "board")
	if err != nil {
		return "{}"
	}

	return string(data)
}

func GetSysUpTime() interface{} {
	var sysTime = make(map[string]interface{})

	data, err := util.RunUbusCommand("call", "system", "info")
	var result map[string]interface{}
	if err = json.Unmarshal(data, &result); err != nil {
		return "{}"
	}

	var localTime, upTime float64
	if v, ok := result["localtime"]; ok {
		localTime = v.(float64)
	}

	if v, ok := result["uptime"]; ok {
		upTime = v.(float64)
	}

	sysTime["localtime"] = util.FormatTimestamp(int64(localTime))
	sysTime["uptime"] = upTime

	if upTime > 0 {
		days := int(upTime / 86400)
		hours := int((upTime - float64(days*86400)) / 3600)
		minutes := int((upTime - float64(days*86400) - float64(hours*3600)) / 60)
		seconds := int(upTime - float64(days*86400) - float64(hours*3600) - float64(minutes*60))

		sysTime["desc"] = fmt.Sprintf("%d day, %d hour, %d minute, %d second", days, hours, minutes, seconds)
	}

	return util.Map2JSON(sysTime)
}
