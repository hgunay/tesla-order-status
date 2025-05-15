// Author: Hakan Gunay
// Date: 2025-05-15

package utils

func TranslateOrderStatus(v string) string {
	switch v {
	case "BOOKED":
		return "Rezerve Edildi"
	case "DELIVERED":
		return "Teslim Edildi"
	case "CANCELLED":
		return "İptal Edildi"
	default:
		return v
	}
}

func TranslateVehicleStatus(v string) string {
	switch v {
	case "NEW":
		return "Yeni"
	case "USED":
		return "İkinci El"
	default:
		return v
	}
}

func TranslateGateCode(code string) string {
	switch code {
	case "FINANCING_TASK":
		return "Finansman"
	case "FACTORY_GATE":
		return "Fabrika Teslim"
	case "SERVICE_VISIT":
		return "Servis Ziyareti"
	case "FINAL_PAYMENT":
		return "Son Ödeme"
	case "TRADE_IN_TASK":
		return "Takas"
	case "INSURANCE_TASK":
		return "Sigorta"
	case "SCHEDULING_TASK":
		return "Zamanlama"
	case "REGISTRATION_TASK":
		return "Kayıt"
	case "DELIVERY_TASK":
		return "Teslimat"
	case "CONTAINMENT_HOLD":
		return "Gecikme Engeli"
	case "FINAL_INVOICE":
		return "Son Fatura"
	case "FINISHED_GOODS":
		return "Hazır Ürün"
	case "STAGING":
		return "Hazırlık"
	case "ORDER_ACKNOWLEDGEMENT":
		return "Sipariş Onayı"
	default:
		return code
	}
}
