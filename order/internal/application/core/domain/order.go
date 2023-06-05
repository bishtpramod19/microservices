package domain

type OrderItem struct {
	ProductCode string `json:"product_code"`
	UnitPrice float32 `json:"unit_price"`
	Quantity int32 `json:"quantity"`
}

type Order struct{
	Id int64 `json:"id"`
	CustomerId int64 `json:"customer_id"`
	Status string `json:"status"`
	OrderItems []OrderItem `json:"order_items"`
	CreatedAt int64 `json:"created_at"`
}


func NewOrder(customerId int64, orderItems []OrderItem) Order{
return Order{
	
}
}