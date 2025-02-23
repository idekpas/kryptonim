package server

import (
	"github.com/gin-gonic/gin"
	"github.com/idekpas/kryptonim/controllers"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())

	health := new(controllers.HealthController)
	rates := controllers.NewRatesController()
	exchange := controllers.NewExchangeController()

	router.GET("/status", health.Status)
	router.GET("/rates", rates.GetRates)
	router.GET("/exchange", exchange.GetExchange)

	return router
}
