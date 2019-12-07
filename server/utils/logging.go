package utils

import (
	"io"
	"log"
	"os"
)

func LogSettings(logFile string) {
	logfile, err := os.OpenFile(logFile, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		log.Fatalf("Failed to open log file: %s", error.Error)
	}
	multiLogFile := io.MultiWriter(os.Stdout, logfile)
	log.SetFlags(log.LstdFlags | log.Llongfile)
	log.SetOutput(multiLogFile)
}
