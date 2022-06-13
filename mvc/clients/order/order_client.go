package order

import (
	"mvc/model"
	"github.com/jinzhu/gorm"
	log "github.com/sirupsen/logrus"
)

var Db *gorm.DB

func GetOrderById(id int) model.Order {
	var order model.Order
	Db.Where("order_id = ?", id).First(&order)
	log.Debug("Order: ", order)

	return order
}

func InsertOrder(order model.Order) model.Order {
	result := Db.Create(&order)

	if result.Error != nil {
		log.Debug(result.Error, order)
	}
	log.Debug("Order Created: ", order.OrderId)
	return order
}
