package server

import (
	"github.com/gin-gonic/gin"
	"github.com/idekpas/kryptonim/controllers"
)

func NewRouter() *gin.Engine {
	router := gin.New()
	router.Use(gin.Logger(), gin.Recovery())

	health := new(controllers.HealthController)

	router.GET("/status", health.Status)

	return router
}
