package cart

import (
	"mvc/model"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

func GetCartById(id int) model.Cart {
	var cart model.Cart
	Db.Where("cart_id = ?", id).First(&cart)
	log.Debug("Cart: ", cart)

	return cart
}

func InsertCart(cart model.Cart) model.Cart {
	result := Db.Create(&cart)

	if result.Error != nil {
		log.Error("")
	}
	log.Debug("Cart Created: ", cart.CartId)
	return cart
}
