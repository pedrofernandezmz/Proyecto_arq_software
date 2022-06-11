package productController

import (
	"mvc/dto"
	service "mvc/services"
	"net/http"
	"strconv"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
	//"fmt"
)

func GetProductById(c *gin.Context) {
	log.Debug("Product id to load: " + c.Param("id"))
	id, _ := strconv.Atoi(c.Param("id")) //se pasa el id de array a int
	var productDto dto.ProductDto
	productDto, err := service.ProductService.GetProductById(id) //delega el trabajo al service
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, productDto)
}

func GetProducts(c *gin.Context) {
	var productsDto dto.ProductsDto
	productsDto, err := service.ProductService.GetProducts() //delega
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, productsDto)
}


func GetSearch(c *gin.Context) {
	log.Debug("Palabra a buscar: " + c.Param("word"))
	//fmt.Println(c.Param("word"))
	word := c.Param("word") //se pasa el id de array a int
	//fmt.Println(word)
	var productsDto dto.ProductsDto
	productsDto, err := service.ProductService.GetSearch(word) //delega
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, productsDto)
}

func GetCategory(c *gin.Context) {
	log.Debug("Categoria a buscar: " + c.Param("category"))
	//fmt.Println(c.Param("word"))
	category := c.Param("category") //se pasa el id de array a int
	//fmt.Println(word)
	var productsDto dto.ProductsDto
	productsDto, err := service.ProductService.GetCategory(category) //delega
	if err != nil {
		c.JSON(err.Status(), err)
		return
	}
	c.JSON(http.StatusOK, productsDto)
}