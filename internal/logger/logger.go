package logger

import (
	"fmt"
	"log"
	"os"
	"time"
)

var path = "logs"

func Debug(tag string, message string) {
	if os.Getenv("ENV") == "DEV" {
		pushLog(tag, message, "DEBUG")
	}
}

func Info(tag string, message string) {
	pushLog(tag, message, "INFO")
}

func Warning(tag string, message string) {
	pushLog(tag, message, "WARNING")
}

func Error(tag string, err error) {
	pushLog(tag, err.Error(), "ERROR")
}

func pushLog(tag string, message string, logType string) {
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, 0777)
	}

	if message == "" {
		Warning("Logger", "Log TAG empty")
	}

	if message == "" {
		Warning("Logger", "Log message empty")
	}

	file, err := os.OpenFile("logs/server.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0777)
	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	loggerPrefix := os.Getenv("LOGGER_PREFIX")
	if loggerPrefix != "" {
		log.SetPrefix(fmt.Sprintf("%s - ", loggerPrefix))
	}
	log.SetFlags(0)
	log.SetOutput(file)
	line := fmt.Sprintf("%s - %s - %s: %s", time.Now().Format(time.RFC3339), tag, logType, message)
	log.Print(line)
	fmt.Println(line)
}
