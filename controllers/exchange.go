package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/idekpas/kryptonim/services"
)

type ExchangeController struct {
	ExchangeService services.ExchangeService
}

func NewExchangeController() *ExchangeController {
	return &ExchangeController{
		ExchangeService: services.NewDefaultExchangeService(),
	}
}

func (e ExchangeController) GetExchange(c *gin.Context) {
	from := c.Query("from")
	to := c.Query("to")
	amount := c.Query("amount")

	if from == "" || to == "" || amount == "" {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	result, err := e.ExchangeService.Exchange(from, to, amount)
	if err != nil || result == nil {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	c.JSON(http.StatusOK, result)
}
