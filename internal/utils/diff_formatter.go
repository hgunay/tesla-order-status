// Author: Hakan Gunay
// Date: 2025-05-15

package utils

type DiffFormatter interface {
	Removed(label string, oldVal interface{}) string
	Added(label string, newVal interface{}) string
	Changed(label string, oldVal, newVal interface{}) []string
}
