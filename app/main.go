package main

import (
	"quest/routes"

	"github.com/gin-gonic/gin"
)

// base_url: "169.51.207.94:31290/"

func main() {
	app := gin.Default()

	routes.RegisterRoutes(app)
	app.Run(":3000")
}
