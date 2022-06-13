package cartController

import (
	"mvc/dto"
	"net/http"
	"github.com/gin-gonic/gin"
	log "github.com/sirupsen/logrus"
)

func GetCartById(c *gin.Context) {
	log.Debug("Cart id: " + c.Param("id"))
	var cartDto dto.CartDto
	c.JSON(http.StatusOK, cartDto)
}

func InsertCart(c *gin.Context) {

	return
}
