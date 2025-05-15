// Author: Hakan Gunay
// Date: 2025-05-15

package client

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	"tesla-order-status/internal/order"
)

const appVersion = "4.43.0-3212"

func RetrieveOrders(accessToken string) []order.Order {
	client := &http.Client{}
	req, _ := http.NewRequest("GET", "https://owner-api.teslamotors.com/api/1/users/orders", nil)
	req.Header.Add("Authorization", "Bearer "+accessToken)
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := io.ReadAll(resp.Body)
	var result struct {
		Response []order.Order `json:"response"`
	}

	_ = json.Unmarshal(body, &result)
	return result.Response
}

func GetOrderDetails(orderID string, accessToken string) (map[string]interface{}, error) {
	apiURL := fmt.Sprintf("https://akamai-apigateway-vfx.tesla.com/tasks?deviceLanguage=en&deviceCountry=DE&referenceNumber=%s&appVersion=%s", orderID, appVersion)
	req, _ := http.NewRequest("GET", apiURL, nil)
	req.Header.Add("Authorization", "Bearer "+accessToken)
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var details map[string]interface{}
	err = json.NewDecoder(resp.Body).Decode(&details)
	return details, err
}
