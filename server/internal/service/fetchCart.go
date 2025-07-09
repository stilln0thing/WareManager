package service

import (
	"encoding/json"
    "fmt"
    "net/http"
    "time"

)

type CartItem struct {
    ID       uint   `json:"id"`
    CartID   uint   `json:"cartId"`
    SKU      string `json:"sku"`
    Name     string `json:"name"`
    Quantity int    `json:"quantity"`
}

type Cart struct {
    OrderID string     `json:"orderId"`
    Items   []CartItem `json:"items"`
}

func FetchCart() (*Cart, error){
	url := fmt.Sprintf("http://127.0.0.1:8080/cart")
	client := &http.Client{Timeout : 5*time.Second}

	resp, err := client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to get the cart")
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("status Code not OK")
	}

	var cart Cart
	if err := json.NewDecoder(resp.Body).Decode(&cart); err != nil {
		return nil, fmt.Errorf("cannot decode the cart JSON")
	}
	return &cart, nil
}