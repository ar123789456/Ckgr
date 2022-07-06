package main

import (
	"cgr/server"
	"log"
)

func main() {
	app := server.NewApp()
	log.Println("Listen and Serv localhost:8080")
	if err := app.Run("8080"); err != nil {
		log.Panic(err)
	}
}
