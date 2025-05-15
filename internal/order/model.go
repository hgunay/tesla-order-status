// Author: Hakan Gunay
// Date: 2025-05-15

package order

type Tokens struct {
	AccessToken  string `json:"access_token"`
	RefreshToken string `json:"refresh_token"`
	ExpiresIn    int64  `json:"expires_in"`
	CreatedAt    int64  `json:"created_at"`
}

type Order struct {
	ReferenceNumber string `json:"referenceNumber"`
	OrderStatus     string `json:"orderStatus"`
	ModelCode       string `json:"modelCode"`
	VIN             string `json:"vin"`
	PickupStoreCode string `json:"pickupStoreCode"`
}

type DetailedOrder struct {
	Order   Order                  `json:"order"`
	Details map[string]interface{} `json:"details"`
}
