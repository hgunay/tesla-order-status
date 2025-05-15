// Author: Hakan Gunay
// Date: 2025-05-15

package store

import (
	"encoding/json"
	"os"
)

var storeMap map[string]string

func LoadStoreData(path string) error {
	data, err := os.ReadFile(path)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &storeMap)
	if err != nil {
		return err
	}

	return nil
}

func GetStoreName(pickupStoreCode string) string {
	if name, ok := storeMap[pickupStoreCode]; ok {
		return name
	}
	return "-"
}
