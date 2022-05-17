package controllers

import (
	"net/http"

	"github.com/pedrofernandezmz/arq-software/ej-auth/domain"
	"github.com/pedrofernandezmz/arq-software/ej-auth/services"
	"github.com/gin-gonic/gin"
)

func Login(c *gin.Context) {
	var cred domain.Credentials

	if err := c.BindJSON(&cred); err != nil {
		c.AbortWithError(http.StatusBadRequest, err)
		return
	}

	token, err := services.Login(cred)
	if err != nil {
		c.AbortWithError(http.StatusUnauthorized, err)
		return
	}

	c.JSON(http.StatusOK, token)
}
