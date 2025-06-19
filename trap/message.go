package trap

import (
	"time"

	"go-snmp-agentx/util"
)

func PackageTrapMessage(code int, level, msg string) string {
	var result = make(map[string]interface{})

	result["sn"] = util.DeviceSN
	result["module"] = "snmp"
	result["code"] = code
	result["level"] = level
	result["message"] = msg
	result["timestamp"] = time.Now().Unix()

	return util.Map2JSON(result)
}
