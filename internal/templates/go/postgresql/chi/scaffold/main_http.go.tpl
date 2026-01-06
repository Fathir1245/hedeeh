package main

import (
	"log"
	"net/http"
	"github.com/go-chi/chi"

	"{{.ModuleName}}/internal/router"
)

func main() {
	r := chi.NewRouter()

	router.RegisterUserRoutes(r, nil)

	log.Println("🚀 Server running on :{{.Port}}")
	if err := http.ListenAndServe(":{{.Port}}", r); err != nil {
		log.Fatal(err)
	}
}
