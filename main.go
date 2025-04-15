package main

import (
	"github.com/franciellyl/golang-minha-biblioteca/biblioteca/routes"
)

func main() {
	r := routes.SetupRoutes()
	r.Run(":8080")
}
