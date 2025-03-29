package util

import (
	"encoding/json"
	"os/exec"
)

// Map2JSON 实现将 map 转换为 json 字符串
func Map2JSON(m map[string]interface{}) string {
	var b []byte
	b, err := json.Marshal(m)
	if err != nil {
		return ""
	}

	return string(b)
}

func RunUbusCommand(args ...string) ([]byte, error) {
	cmd := exec.Command("ubus", args...)
	return cmd.Output()
}
