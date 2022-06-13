package app

import (
	cartController "mvc/controllers/cart"
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
	router.GET("/user", userController.GetUsers)
	router.POST("/user", userController.UserInsert)
	// Parte del Login
	router.POST("/login", userController.Login)

	// Products
	router.GET("/product/:product_id", productController.GetProductById)
	router.GET("/products", productController.GetProducts)
	router.GET("/products/:category_id", productController.GetProductsByCategoryId)

	// Orders
	router.GET("/order/:id", orderController.GetOrderById)
	router.POST("/order", orderController.OrderInsert)

	// OrderDetails
	router.GET("/orderDetail/:id", orderDetailController.GetOrderDetailById)
	router.POST("/orderDetail", orderDetailController.OrderDetailInsert)

	// Category
	router.GET("/category/:id", categoryController.GetCategoryById)
	router.GET("/categories", categoryController.GetCategories)

	// Cart
	router.GET("/cart/:id", cartController.GetCartById)
	router.POST("/cart", cartController.InsertCart)

	log.Info("Finishing mappings configurations")
}
