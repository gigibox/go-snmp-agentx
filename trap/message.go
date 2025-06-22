package trap

import (
	"strconv"
	"time"

	"go-snmp-agentx/util"
)

func PackageTrapMessage(code int, level, msg string) string {
	var result = make(map[string]interface{})

	result["sn"] = util.DeviceSN
	result["module"] = "snmp"
	result["code"] = strconv.Itoa(code)
	result["level"] = level
	result["message"] = msg
	result["timestamp"] = time.Now().Unix()
	result["status"] = "alarm"

	if code == 0 {
		result["code"] = "00000"
		result["status"] = "normal"
	}
	return util.Map2JSON(result)
}
