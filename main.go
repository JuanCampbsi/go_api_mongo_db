package main

import (
	"github.com/JuanCampbsi/go-opportunities/config"
	"github.com/JuanCampbsi/go-opportunities/router"
)

var (
	logger *config.Logger
)

func main() {
	logger = config.GetLogger("main")
	//Initialize Configs
	err := config.Init()
	if err != nil {
		logger.ErrorF("config initialization error: %v", err)
		return
	}

	//Initialize Router
	router.Initialize()
}
