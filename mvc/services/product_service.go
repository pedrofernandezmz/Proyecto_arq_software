package services

//lugar donde yo defino los metodos que mi clase va a responder (Interfaz de objetos)
//Se puede reutilizar
import (
	productCliente "mvc/clients/product" //DAO
	"mvc/dto"
	"mvc/model"
	e "mvc/utils/errors"
	//"fmt"
)

type productService struct{}

type productServiceInterface interface {
	//siempre devuelve dto o error
	GetProductById(id int) (dto.ProductDto, e.ApiError)
	GetProducts() (dto.ProductsDto, e.ApiError)
	GetSearch(word string) (dto.ProductsDto, e.ApiError)
}

var (
	ProductService productServiceInterface
)

func init() {
	ProductService = &productService{}
}

func (s *productService) GetProductById(id int) (dto.ProductDto, e.ApiError) {

	var product model.Product = productCliente.GetProductById(id) //objeto de la DB, a traves del DAO
	var productDto dto.ProductDto

	if product.Id == 0 {
		return productDto, e.NewBadRequestApiError("product not found")
	}
	productDto.Name = product.Name
	productDto.Id = product.Id
	return productDto, nil
}

func (s *productService) GetProducts() (dto.ProductsDto, e.ApiError) {

	var products model.Products = productCliente.GetProducts()
	var productsDto dto.ProductsDto

	for _, product := range products {
		var productDto dto.ProductDto
		productDto.Name = product.Name
		productDto.Price = product.Price
		productDto.Picture = product.Picture
		productDto.IdCategory = product.IdCategory
		productDto.Id = product.Id

		productsDto = append(productsDto, productDto)
	}

	return productsDto, nil
}

func (s *productService) GetSearch(word string) (dto.ProductsDto, e.ApiError) {

	//fmt.Println(word)

	var products model.Products = productCliente.GetSearch(word)
	var productsDto dto.ProductsDto

	for _, product := range products {
		var productDto dto.ProductDto
		productDto.Name = product.Name
		productDto.Price = product.Price
		productDto.Picture = product.Picture
		productDto.IdCategory = product.IdCategory
		productDto.Id = product.Id

		productsDto = append(productsDto, productDto)
	}

	return productsDto, nil
}
