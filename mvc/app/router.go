package app

import (
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

var (
	router *gin.Engine
)

func init() {
	router = gin.Default()
}

func StartRoute() {
	mapUrls()

	log.Info("Iniciando Servidor...")
	router.Run(":8080")

}
