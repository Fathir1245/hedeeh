package main

import (
	"github.com/gin-gonic/gin"

	"{{.ModuleName}}/internal/router"
)

func main() {
	r := gin.Default()

	router.RegisterUserRoutes(r, nil)

	log.Println("🚀 Server running on :{{.Port}}")
	r.Run("{.Port}")

}
