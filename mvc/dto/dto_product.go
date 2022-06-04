package dto

type ProductDto struct {
	Name string `json:"name"`
	Picture     string  `json:"picture"`
	Price       int     `json:"price"`
	Id          int     `json:"id"`
	IdCategory  int     `json:"id_category"`
}

type ProductsDto []ProductDto

//cambio de prueba
