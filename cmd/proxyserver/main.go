package main

import (
	"os"
	"os/signal"

	_ "github.com/joho/godotenv/autoload"

	"brsrker.com/diamond/proxyserver/internal/data"
	"brsrker.com/diamond/proxyserver/internal/logger"
	"brsrker.com/diamond/proxyserver/internal/server"
)

const TAG = "main"

func main() {
	port := os.Getenv("PORT")
	serv, err := server.New(port)
	if err != nil {
		logger.Error(TAG, err)
	}

	// connection to the database.
	d := data.New()
	if err := d.DB.Ping(); err != nil {
		logger.Error(TAG, err)
	}

	// start the server.
	go serv.Start()

	// Wait for an in interrupt.
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt)
	<-c

	// Attempt a graceful shutdown.
	serv.Close()
	data.Close()
}
