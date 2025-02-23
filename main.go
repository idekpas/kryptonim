package main

import (
	"flag"

	"github.com/idekpas/kryptonim/config"
	"github.com/idekpas/kryptonim/server"
)

func main() {
	env := flag.String("e", "dev", "")
	flag.Parse()
	config.Init(*env)
	server.Init()
}
