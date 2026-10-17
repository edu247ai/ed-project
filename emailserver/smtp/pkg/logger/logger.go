package logger

import (
	"log"
	"os"
)

func Info(msg string, args ...interface{}) {
	log.Printf("[SMTP][INFO] "+msg, args...)
}

func Error(msg string, args ...interface{}) {
	log.Printf("[SMTP][ERROR] "+msg, args...)
}

func Debug(msg string, args ...interface{}) {
	log.Printf("[SMTP][DEBUG] "+msg, args...)
}

func init() {
	log.SetOutput(os.Stdout)
	log.SetFlags(log.Ldate | log.Ltime | log.LUTC)
}
