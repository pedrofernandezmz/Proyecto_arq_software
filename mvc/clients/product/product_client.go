package product

import (
	"mvc/model"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
	//"fmt"
	//"strconv"
)

var Db *gorm.DB //ORM para la base de datos

func GetProductById(id int) model.Product {
	var product model.Product
	Db.Where("id = ?", id).First(&product) //traduccion y seteo en product
	log.Debug("Product: ", product)

	return product
}

func GetProducts() model.Products {
	var products model.Products
	Db.Find(&products) //busca y guarda todo en products
	log.Debug("Products: ", products)

	return products
}

func GetSearch(word string) model.Products {
	var products model.Products
	//fmt.Println("%"+word+"%")
	//fmt.Println(word)
	// LIKE
    Db.Where("name LIKE ?", "%"+word+"%").Find(&products)
    // SELECT * FROM users WHERE name LIKE '%jin%';
	//Db.Find(&products) //busca y guarda todo en products
	log.Debug("Products: ", products)

	return products
}