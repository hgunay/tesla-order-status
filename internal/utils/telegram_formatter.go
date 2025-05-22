// Author: Hakan Gunay
// Date: 2025-05-15

package utils

import "fmt"

type TelegramFormatter struct{}

func (f TelegramFormatter) Removed(label string, oldVal interface{}) string {
	return fmt.Sprintf("🕓 %s: %v", label, oldVal)
}

func (f TelegramFormatter) Added(label string, newVal interface{}) string {
	return fmt.Sprintf("⚡ %s: %v", label, newVal)
}

func (f TelegramFormatter) Changed(label string, oldVal, newVal interface{}) []string {
	return []string{
		fmt.Sprintf("🕓 %s: %v", label, oldVal),
		fmt.Sprintf("⚡ %s: %v", label, newVal),
	}
}
