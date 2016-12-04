package main

import (
	"comentarismo-age/server"
	"os"
)

var Port = os.Getenv("PORT")

func main() {
	if Port == "" {
		Port = "3005"
	}
	server.StartServer(Port)
}
