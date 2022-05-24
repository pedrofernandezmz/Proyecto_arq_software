package main

import (
	"github.com/pedrofernandezmz/Proyecto_arq_software/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	engine := gin.New()
	engine.POST("/login",
		controllers.Login)
	engine.Run(":8080")
}
