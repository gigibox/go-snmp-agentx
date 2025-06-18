package util

import (
	"encoding/json"
	"fmt"
	"math"
	"os"
	"os/exec"
	"strings"
	"time"
)

var DeviceSN, DeviceStatus string

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

func GetBrLanMAC() (string, error) {
	// 读取 br-lan 接口的 MAC 地址
	data, err := os.ReadFile("/sys/class/net/br-lan/address")
	if err != nil {
		return "", fmt.Errorf("读取 MAC 地址失败: %w", err)
	}

	mac := strings.TrimSpace(string(data))           // 原始: "b8:27:eb:12:34:56"
	mac = strings.ReplaceAll(mac, ":", "")           // 去除冒号: "b827eb123456"
	mac = strings.ToUpper(mac)                       // 转为大写: "B827EB123456"
	return mac, nil
}