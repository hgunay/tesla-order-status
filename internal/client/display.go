// Author: Hakan Gunay
// Date: 2025-05-15

package client

import (
	"fmt"
	"strings"

	"tesla-order-status/internal/store"
	"tesla-order-status/internal/utils"
)

func displayTaskStatus(tasks map[string]interface{}) {
	fmt.Println("\n" + strings.Repeat("-", 45))
	fmt.Println(utils.CenterText("GÖREV DURUMU", 45))
	fmt.Println(strings.Repeat("-", 45))
	for key, val := range tasks {
		status := "✗"
		m, ok := val.(map[string]interface{})
		if ok {
			if complete, ok := m["isBlocker"].(bool); ok && !complete {
				status = "✓"
			}
		}
		fmt.Printf("%s %s\n", status, utils.TranslateGateCode(key))
	}
}

func FormatReadable(v interface{}) string {
	if str, ok := v.(string); ok {
		return strings.ReplaceAll(strings.Title(strings.ToLower(strings.ReplaceAll(str, "_", " "))), "Pro ", "PRO ")
	}
	return utils.GetReadableValue(v)
}

func DisplayAdditionalOrderDetails(details map[string]interface{}) {
	detailsMap := utils.GetMap(details, "state")
	detailsMap = utils.GetMap(details, "tasks", "registration", "orderDetails")
	paymentMap := utils.GetMap(details, "tasks", "finalPayment", "data")
	regDetails := utils.GetMap(details, "tasks", "registration", "regData", "regDetails")
	owner := utils.GetMap(regDetails, "owner")
	address := utils.GetMap(details, "tasks", "registration", "regData", "registrationAddress")
	scheduling := utils.GetMap(details, "tasks", "scheduling")
	taskMap := utils.GetMap(details, "tasks", "deliveryAcceptance", "gates")
	insurance := utils.GetMap(details, "tasks", "insurance")
	orderMap := utils.GetMap(details, "tasks", "order")

	fmt.Println("\n" + strings.Repeat("-", 45))
	fmt.Println(utils.CenterText("ARAÇ BİLGİLERİ", 45))
	fmt.Println(strings.Repeat("-", 45))
	fmt.Printf("%s %v\n", utils.ColorText("- Model Kodu:", "94"), strings.ToUpper(utils.GetReadableValue(detailsMap["modelCode"])))
	fmt.Printf("%s %v\n", utils.ColorText("- Araç Tipi:", "94"), utils.TranslateVehicleStatus(utils.GetReadableValue(detailsMap["vehicleTitleStatus"])))
	fmt.Printf("%s %v\n", utils.ColorText("- VIN:", "94"), utils.SafeVIN(utils.GetReadableValue(detailsMap["vin"])))
	fmt.Printf("%s %v\n", utils.ColorText("- Sipariş Durumu:", "94"), utils.TranslateOrderStatus(utils.GetReadableValue(detailsMap["orderStatus"])))
	fmt.Printf("%s %v\n", utils.ColorText("- Sipariş Alt Durumu:", "94"), utils.GetReadableValue(detailsMap["orderSubstatus"]))
	fmt.Printf("%s %v\n", utils.ColorText("- Seri:", "94"), FormatReadable(detailsMap["series"]))
	fmt.Printf("%s %v\n", utils.ColorText("- Trim Kodu:", "94"), utils.GetReadableValue(detailsMap["trimCode"]))
	fmt.Printf("%s %v\n", utils.ColorText("- Araç Konum Kodu:", "94"), utils.GetReadableValue(detailsMap["vehicleRoutingLocation"]))

	if f, ok := detailsMap["vehicleMapId"].(float64); ok {
		fmt.Printf("%s %d\n", utils.ColorText("- Konfigürasyon ID:", "94"), int64(f))
	} else {
		fmt.Printf("%s %v\n", utils.ColorText("- Konfigürasyon ID:", "94"), utils.GetReadableValue(detailsMap["vehicleMapId"]))
	}

	fmt.Printf("%s %v\n", utils.ColorText("- Üretim Yılı:", "94"), utils.GetReadableValue(detailsMap["vehicleModelYear"]))
	fmt.Printf("%s %v\n", utils.ColorText("- Sipariş Tarihi:", "94"), utils.FormatDateTime(detailsMap["marketingLexiconDate"]))

	odom := utils.SafeFloatFormat(detailsMap["vehicleOdometer"])
	typ := utils.GetReadableValue(detailsMap["vehicleOdometerType"])
	fmt.Printf("%s %s %s\n", utils.ColorText("- Kilometre:", "94"), odom, strings.ToUpper(typ))

	fmt.Println("\n" + strings.Repeat("-", 45))
	fmt.Println(utils.CenterText("MÜŞTERİ BİLGİLERİ", 45))
	fmt.Println(strings.Repeat("-", 45))
	fmt.Printf("%s %v\n", utils.ColorText("- İsim:", "94"), utils.GetSafeString(utils.GetMap(owner, "user"), "firstName"))
	fmt.Printf("%s %v\n", utils.ColorText("- Soyisim:", "94"), utils.GetSafeString(utils.GetMap(owner, "user"), "lastName"))
	fmt.Printf("%s %v\n", utils.ColorText("- E-posta:", "94"), utils.GetSafeString(owner, "email"))
	fmt.Printf("%s %v\n", utils.ColorText("- Telefon:", "94"), utils.GetSafeString(owner, "phoneNumber"))

	fmt.Println("\n" + utils.ColorText("Adres Bilgileri:", "94"))
	fmt.Printf("%s %v\n", utils.ColorText("- Adres:", "94"), utils.GetSafeString(address, "address1"))
	fmt.Printf("%s %v\n", utils.ColorText("- Şehir:", "94"), utils.GetSafeString(address, "city"))
	fmt.Printf("%s %v\n", utils.ColorText("- İl:", "94"), utils.GetSafeString(address, "stateProvince"))
	fmt.Printf("%s %v\n", utils.ColorText("- Posta Kodu:", "94"), utils.GetSafeString(address, "zipCode"))

	fmt.Println("\n" + strings.Repeat("-", 45))
	fmt.Println(utils.CenterText("ÖDEME BİLGİLERİ", 45))
	fmt.Println(strings.Repeat("-", 45))
	fmt.Printf("%s %s TL\n", utils.ColorText("- Toplam Tutar:", "94"), utils.SafeFloatFormat(paymentMap["amountDue"]))
	fmt.Printf("%s %s TL\n", utils.ColorText("- Ödenen Tutar:", "94"), utils.SafeFloatFormat(paymentMap["amountSent"]))
	fmt.Printf("%s %s TL\n", utils.ColorText("- Ön Ödeme Tutarı:", "94"), utils.SafeFloatFormat(detailsMap["reservationAmountReceived"]))

	fmt.Println("\n" + strings.Repeat("-", 45))
	fmt.Println(utils.CenterText("KAYIT DURUMU", 45))
	fmt.Println(strings.Repeat("-", 45))
	fmt.Printf("%s %v\n", utils.ColorText("- Kayıt Durumu:", "94"), utils.GetReadableValue(regDetails["registrationStatus"]))
	fmt.Printf("%s %v\n", utils.ColorText("- Kayıt Tarihi:", "94"), utils.FormatDateTime(regDetails["lastUpdateDatetime"]))

	fmt.Println("\n" + strings.Repeat("-", 45))
	fmt.Println(utils.CenterText("TESLİMAT", 45))
	fmt.Println(strings.Repeat("-", 45))
	fmt.Printf("%s %v\n", utils.ColorText("- Tahmini Teslimat Aralığı:", "94"), utils.GetReadableValue(scheduling["deliveryWindowDisplay"]))
	fmt.Printf("%s %v\n", utils.ColorText("- Tahmini Varış Tarihi:", "94"), utils.FormatDateTime(paymentMap["etaToDeliveryCenter"]))
	fmt.Printf("%s %v\n", utils.ColorText("- Teslimat Randevusu:", "94"), utils.FormatDateTime(scheduling["apptDateTimeAddressStr"]))
	fmt.Printf("%s %v\n", utils.ColorText("- Teslimat Yeri:", "94"), store.GetStoreName(utils.GetReadableValue(orderMap["pickupStoreCode"])))

	displayTaskStatus(taskMap)

	fmt.Println("\n" + strings.Repeat("-", 45))
	fmt.Println(utils.CenterText("SİGORTA BİLGİSİ", 45))
	fmt.Println(strings.Repeat("-", 45))
	fmt.Printf("%s %v\n", utils.ColorText("- Sigorta Durumu:", "94"), utils.GetReadableValue(insurance["status"]))
	fmt.Printf("%s %v\n", utils.ColorText("- Sigorta Şirketi:", "94"), utils.GetReadableValue(insurance["insuranceCompanyName"]))
}
