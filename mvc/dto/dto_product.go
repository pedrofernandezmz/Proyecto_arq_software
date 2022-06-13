package dto

type ProductDto struct {
	ProductId   int         `json:"product_id"`
	Category    CategoryDto `json:"category"`
	Name        string      `json:"name"`
	Price       float32     `json:"base_price"`
	Picture     string      `json:"picture_url"`
}

type ProductsDto []ProductDto
