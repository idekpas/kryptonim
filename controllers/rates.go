package controllers

import (
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/idekpas/kryptonim/services"
)

type RatesController struct{}

func (h RatesController) GetRates(c *gin.Context) {
	currencies := c.Query("currencies")
	currencyList := strings.Split(currencies, ",")
	if currencies == "" || len(currencyList) < 2 {
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	rates, err := services.GetExchangeRates(currencyList)
	if err != nil {
		log.Print(err.Error())
		c.JSON(http.StatusBadRequest, gin.H{})
		return
	}

	c.JSON(http.StatusOK, rates)
}
