package main

import (
	"os"

	"github.com/FlyingFeather/golang-data-demo/service"
	flag "github.com/spf13/pflag"
	//"github.com/pmlpml/golang-learning/web/cloudgo/service"
)

const (
	PORT string = "8080"
)

func main() {
	port := os.Getenv("PORT")
	if len(port) == 0 {
		port = PORT
	}

	pPort := flag.StringP("port", "p", PORT, "PORT for httpd listening")
	flag.Parse()
	if len(*pPort) != 0 {
		port = *pPort
	}

	server := service.NewServer()
	server.Run(":" + port)
}
