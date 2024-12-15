package socket

import (
	"net"
	"os"

	"beloin.com/dataencoding/pkg/logger"
)

var CONN_TYPE = "tcp"

// TODO: Implement server

func Listen(host string, port string) {
	addr := host + ":" + port
	l, err := net.Listen(CONN_TYPE, addr)
	if err != nil {
		logger.Error.Println("Error listening:", err.Error())
		os.Exit(1)
	}
	
  defer l.Close()
  logger.Info.Println("Listening on " + addr)
}
