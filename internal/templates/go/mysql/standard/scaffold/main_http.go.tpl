package main

import (
	"log"
	"net/http"

	"{{.ModuleName}}/internal/router"
)

func main() {
	r := router.NewRouter()

	log.Println("🚀 Server running on :{{.Port}}")
	if err := http.ListenAndServe(":{{.Port}}", r); err != nil {
		log.Fatal(err)
	}
}
