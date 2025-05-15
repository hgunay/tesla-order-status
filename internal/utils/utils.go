// Author: Hakan Gunay
// Date: 2025-05-15

package utils

import (
	"fmt"
	"strings"
	"time"
)

func ColorText(text string, colorCode string) string {
	return fmt.Sprintf("\033[%sm%s\033[0m", colorCode, text)
}

func FormatDateTime(raw any) string {
	str, ok := raw.(string)
	if !ok || str == "" {
		return "-"
	}

	layouts := []string{
		"2006-01-02T15:04:05.000000",
		"2006-01-02T15:04:05.000",
		"2006-01-02T15:04:05",
		"2006-01-02",
		time.RFC3339,
	}

	for _, layout := range layouts {
		if t, err := time.Parse(layout, str); err == nil {
			return t.In(time.FixedZone("TRT", 3*60*60)).Format("02.01.2006 15:04")
		}
	}
	return str
}

func SafeVIN(vin string) string {
	if strings.TrimSpace(vin) == "" {
		return "N/A"
	}
	return vin
}

func GetMap(data map[string]interface{}, keys ...string) map[string]interface{} {
	current := data
	for _, key := range keys {
		if next, ok := current[key].(map[string]interface{}); ok {
			current = next
		} else {
			return map[string]interface{}{}
		}
	}
	return current
}

func CenterText(text string, width int) string {
	pad := (width - len(text)) / 2
	return strings.Repeat(" ", pad) + text
}

func GetSafeString(m map[string]interface{}, key string) string {
	if v, ok := m[key]; ok && v != nil {
		return fmt.Sprintf("%v", v)
	}
	return "-"
}

func GetReadableValue(v interface{}) string {
	if v == nil {
		return "(bilgi yok)"
	}
	return fmt.Sprintf("%v", v)
}

func SafeFloatFormat(val any) string {
	f, ok := val.(float64)
	if !ok {
		return "-"
	}
	return fmt.Sprintf("%.2f", f)
}
