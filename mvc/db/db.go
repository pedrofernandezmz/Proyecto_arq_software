package db

import (
	cartClient "mvc/clients/cart"
	categoryClient "mvc/clients/category"
	orderClient "mvc/clients/order"
	orderDetailClient "mvc/clients/order_detail"
	productClient "mvc/clients/product"
	userClient "mvc/clients/user"
	"mvc/model"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	log "github.com/sirupsen/logrus"
)

var (
	db  *gorm.DB
	err error
)

func init() {
	//BASE DE DATOS
	DBName := "Proyecto"
	DBUser := "root"
	DBPass := ""
	DBHost := "localhost"

	db, err = gorm.Open("mysql", DBUser+":"+DBPass+"@tcp("+DBHost+":3306)/"+DBName+"?charset=utf8&parseTime=True")

	if err != nil {
		log.Info("Connection Failed to Open")
		log.Fatal(err)
	} else {
		log.Info("Connection Established")
	}

	//SE AGREGAN LOS CLIENTS
	userClient.Db = db
	productClient.Db = db
	orderClient.Db = db
	orderDetailClient.Db = db
	categoryClient.Db = db
	cartClient.Db = db

}

func StartDbEngine() {
	//CLASES
	db.AutoMigrate(&model.User{}, &model.Product{}, &model.Order{}, &model.OrderDetail{}, &model.Category{}, &model.Cart{})

	log.Info("Finish Database Migration")
}
