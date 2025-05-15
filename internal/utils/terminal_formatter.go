// Author: Hakan Gunay
// Date: 2025-05-15

package utils

import "fmt"

type TerminalFormatter struct{}

func (f TerminalFormatter) Removed(label string, oldVal interface{}) string {
	return fmt.Sprintf("\033[91m- %s: %v\033[0m", label, oldVal)
}

func (f TerminalFormatter) Added(label string, newVal interface{}) string {
	return fmt.Sprintf("\033[92m+ %s: %v\033[0m", label, newVal)
}

func (f TerminalFormatter) Changed(label string, oldVal, newVal interface{}) []string {
	return []string{
		fmt.Sprintf("\033[91m- %s: %v\033[0m", label, oldVal),
		fmt.Sprintf("\033[92m+ %s: %v\033[0m", label, newVal),
	}
}
