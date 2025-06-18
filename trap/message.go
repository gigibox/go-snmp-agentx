package trap

import (
	"fmt"
	"time"

	"go-snmp-agentx/util"
)

func PackageTrapMessage(code int, level, msg string) string {
	return fmt.Sprintf(`
	"sn": %s,
	"module": "snmp",
	"code": %d,
	"level": %s,
	"message": %s,
	"timestamp": %d
	`,util.DeviceSN, code, level, msg, time.Now().Unix())
}
