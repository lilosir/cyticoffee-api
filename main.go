package main

import (
	"github.com/lilosir/cyticoffee-api/routes"
)

func main() {
	r := routes.SetupRoutes()
	r.Run(":8090") // listen and serve on 0.0.0.0:8080
}
