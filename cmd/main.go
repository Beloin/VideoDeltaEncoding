package main

import (
	"beloin.com/dataencoding/pkg/logger"
)

func main() {
	logger.ConfigureLogger()
	logger.Info.Println("Started Application")
}
