package util

import (
	"encoding/json"
	"math"
	"os/exec"
	"time"
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

func FormatTimestamp(timestamp int64) string {
	// 将时间戳转换为时间对象
	t := time.Unix(timestamp, 0)
	// 格式化为日期字符串
	return t.Format("2006-01-02 15:04:05")
}

func RoundFloat(val float64, precision uint) float64 {
	ratio := math.Pow(10, float64(precision))
	return math.Round(val*ratio) / ratio
}
