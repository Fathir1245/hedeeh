package main

import (
	"log"
	"net/http"
	"github.com/gin-gonic/gin"

	"{{.ModuleName}}/internal/router"
)

func main() {
	r := gin.Default()
	router.RegisterProductRoutes(r, handler)
	r.Run(":{.Port}")


	log.Println("🚀 Server running on :{{.Port}}")
	if err := http.ListenAndServe(":{{.Port}}", r); err != nil {
		log.Fatal(err)
	}
}
