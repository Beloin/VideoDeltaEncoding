package logger

import (
	"log"
	"os"
)

var (
	Warning *log.Logger
	Info    *log.Logger
	Error   *log.Logger
)

// ANSI escape codes for colors
const (
	Reset  = "\033[0m"
	Red    = "\033[31m"
	Green  = "\033[32m"
	Yellow = "\033[33m"
	Blue   = "\033[34m"
)

func ConfigureLogger() {
	Info = log.New(os.Stdout, Green+"I: "+Reset, log.Ldate|log.Ltime|log.Lshortfile)
	Warning = log.New(os.Stdout, Yellow+"W: "+Reset, log.Ldate|log.Ltime|log.Lshortfile)
	Error = log.New(os.Stderr, Red+"E: "+Reset, log.Ldate|log.Ltime|log.Lshortfile)
}
