package response

type OrderResponse struct {
	OrderId       int    `json:"id"`
	ProductName   string `json:"product_name"`
	OrderType     string `json:"order_type"`
	OrderPrice    int    `json:"order_price"`
	OrderQuantity int    `json:"order_quantity"`
}
