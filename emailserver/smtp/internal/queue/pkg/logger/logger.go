package logger

import (
	"log"
	"os"
)

func init() {
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Ldate | log.Ltime | log.LUTC)
}

func Info(msg string, args ...interface{}) {
	log.Printf("[INFO] "+msg, args...)
}

func Error(msg string, args ...interface{}) {
	log.Printf("[ERROR] "+msg, args...)
}
