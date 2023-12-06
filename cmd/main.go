package main

import (
	"as-capital-crawler-fi-ms/internal/api/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	app := gin.Default()
	routes.AppRoutes(app)
	app.Run(":3001")
}
