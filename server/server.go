package server

import (
	"github.com/idekpas/kryptonim/config"
)

func Init() {
	cfg := config.GetConfig()
	router := NewRouter()
	router.Run(cfg.GetString("server.port"))
}
