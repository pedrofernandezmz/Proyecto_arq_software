package app

import (
	cors "github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

var (
	router *gin.Engine
)

func init() {
	router = gin.Default()
	router.Use(cors.Default()) //RESUELVE PROBLEMA CORS-POLICY
}

func StartRoute() {
	mapUrls()
	log.Info("Starting...")
	router.Run(":8080") //CORRE EN EL PUERTO 8080

}
