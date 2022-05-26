package main

import (
	"cgr/server"
	"log"
)

func main() {
	app := server.NewApp()
	if err := app.Run("8080"); err != nil {
		log.Panic(err)
	}
}
