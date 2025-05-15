// Author: Hakan Gunay
// Date: 2025-05-15

package order

import (
	"encoding/json"
	"os"
	"reflect"

	"tesla-order-status/internal/utils"
)

var fieldLabelMap = map[string]string{
	"vehicleOdometer":           "Kilometre",
	"vehicleOdometerType":       "Kilometre Tipi",
	"vin":                       "Araç VIN",
	"modelCode":                 "Model Kodu",
	"orderStatus":               "Sipariş Durumu",
	"orderSubstatus":            "Sipariş Alt Durumu",
	"series":                    "Seri",
	"trimCode":                  "Trim Kodu",
	"vehicleRoutingLocation":    "Araç Konum Kodu",
	"vehicleMapId":              "Konfigürasyon ID",
	"vehicleMktConfiguration":   "Pazarlama Adı",
	"vehicleModelYear":          "Üretim Yılı",
	"marketingLexiconDate":      "Sipariş Tarihi",
	"amountDue":                 "Toplam Tutar",
	"amountSent":                "Ödenen Tutar",
	"reservationAmountReceived": "Ön Ödeme Tutarı",
	"registrationStatus":        "Kayıt Durumu",
	"lastUpdateDatetime":        "Kayıt Tarihi",
	"deliveryWindowDisplay":     "Tahmini Teslimat Aralığı",
	"apptDateTimeAddressStr":    "Teslimat Randevusu",
	"status":                    "Sigorta Durumu",
	"insuranceCompanyName":      "Sigorta Şirketi",
	"etaToDeliveryCenter":       "Tahmini Varış Tarihi",
}

func SaveOrdersToFile(orders []DetailedOrder, filePath string) error {
	file, err := os.Create(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	return json.NewEncoder(file).Encode(orders)
}

func LoadOrdersFromFile(filePath string) ([]DetailedOrder, error) {
	if _, err := os.Stat(filePath); os.IsNotExist(err) {
		return nil, nil
	}
	file, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	var orders []DetailedOrder
	err = json.NewDecoder(file).Decode(&orders)
	return orders, err
}

func CompareDicts(oldMap, newMap map[string]interface{}, path string, formatter utils.DiffFormatter) []string {
	differences := []string{}
	allKeys := map[string]struct{}{}

	for k := range oldMap {
		allKeys[k] = struct{}{}
	}
	for k := range newMap {
		allKeys[k] = struct{}{}
	}

	for key := range allKeys {
		oldVal, oldOk := oldMap[key]
		newVal, newOk := newMap[key]

		labelPath := path + key
		label := labelPath
		if labelStr, ok := fieldLabelMap[key]; ok {
			label = labelStr
		}

		switch {
		case oldOk && newOk:
			switch oldTyped := oldVal.(type) {
			case map[string]interface{}:
				if newTyped, ok := newVal.(map[string]interface{}); ok {
					differences = append(differences, CompareDicts(oldTyped, newTyped, labelPath+".", formatter)...)
				} else if !reflect.DeepEqual(oldVal, newVal) {
					differences = append(differences, formatter.Changed(label, oldVal, newVal)...)
				}
			default:
				if !reflect.DeepEqual(oldVal, newVal) {
					differences = append(differences, formatter.Changed(label, oldVal, newVal)...)
				}
			}

		case oldOk && !newOk:
			differences = append(differences, formatter.Removed(label, oldVal))

		case !oldOk && newOk:
			differences = append(differences, formatter.Added(label, newVal))
		}
	}

	return differences
}
