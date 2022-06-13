package dto

type OrderDetailInsertDto struct {
	ProductId  int     `json:"product_id"`
	Quantity   int     `json:"quantity"`
	Price      float32 `json:"price"`
	Name       string  `json:"name"`
}

type OrderDetailsInsertDto []OrderDetailInsertDto
