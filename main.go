package main

import (
	"github.com/Fawaidac/siapel_backend/configs"
	"github.com/Fawaidac/siapel_backend/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	
	configs.InitializeJWT()

	r := gin.Default()

	routes.SetupRoutes(r)

	r.Run(":8080")
}
