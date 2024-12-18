package main

import (
	"beloin.com/dataencoding/internal/encoding"
	"beloin.com/dataencoding/pkg/logger"
)

var encoder encoding.Encoder

func main() {
  encoder = &encoding.HardEncode{}
	logger.ConfigureLogger()
	logger.Info.Println("Started Application")
}
