package db

import (
	//product "mvc/clients/product"
	addressClient "mvc/clients/address"
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
	// Autorizaciones para la Base de Datos
	DBName := "Proyecto"
	DBUser := "root"
	DBPass := ""
	DBHost := "127.0.0.1"

	db, err = gorm.Open("mysql", DBUser+":"+DBPass+"@tcp("+DBHost+":3306)/"+DBName+"?charset=utf8&parseTime=True") //ver puerto

	if err != nil {
		log.Info("Connection Failed to Open")
		log.Fatal(err)
	} else {
		log.Info("Connection Established")
	}

	// We need to add all CLients that we build
	userClient.Db = db
	productClient.Db = db
	addressClient.Db = db
	categoryClient.Db = db
	orderDetailClient.Db = db
	orderClient.Db = db
}

func StartDbEngine() {
	// We need to migrate all classes model.
	db.AutoMigrate(&model.User{}) //se debe agreagar por cada carpeta del model
	db.AutoMigrate(&model.Product{})
	db.AutoMigrate(&model.Address{})
	db.AutoMigrate(&model.Category{})
	db.AutoMigrate(&model.OrderDetail{})
	db.AutoMigrate(&model.Order{})
	log.Info("Finishing Migration Database Tables")
}
