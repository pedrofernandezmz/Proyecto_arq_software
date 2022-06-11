package app

import (
	addressController "mvc/controllers/address"
	categoryController "mvc/controllers/category"
	orderController "mvc/controllers/order"
	orderDetailController "mvc/controllers/order_detail"
	productController "mvc/controllers/product"
	userController "mvc/controllers/user"

	log "github.com/sirupsen/logrus"
)

func mapUrls() {
	// Users
	router.GET("/user/:id", userController.GetUserById)
	//router.OPTIONS("/user/:id", userController.GetUserById)
	router.GET("/user", userController.GetUsers)
	router.POST("/login", userController.UserLogin)

	//Products
	router.GET("/product/:id", productController.GetProductById)
	router.GET("/product", productController.GetProducts)
	router.GET("/search/:word", productController.GetSearch) //busca por palabra en el name
	router.GET("/searchcategory/:category", productController.GetCategory) //busca por categoria

	//Details
	router.GET("/orderDetail/:id", orderDetailController.GetOrderDetailById)
	router.GET("/orderDetail", orderDetailController.GetOrderDetails)
	router.POST("/orderDetail", orderDetailController.OrderDetailInsert)

	//Orders
	router.GET("/order/:id", orderController.GetOrderById)
	router.GET("/order", orderController.GetOrders)
	router.POST("/order", orderController.OrderInsert)
	router.GET("/orderUser/:idUser", orderController.GetOrdersByIdUser)

	//Categorys
	router.GET("/category/:id", categoryController.GetCategoryById)
	router.GET("/category", categoryController.GetCategories)

	//Addresses
	router.GET("/address/:id", addressController.GetAddressById)
	router.GET("/address", addressController.GetAddresses)

	log.Info("Mappings finalizados!")
}
