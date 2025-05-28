// Package main es el punto de entrada del servidor HTTP.
package main

import (
	"fmt"
	"net/http"

	"github.com/ElYares/go-api-starter/internal/config"
	"github.com/ElYares/go-api-starter/internal/handler"
	"github.com/go-chi/chi/v5"
)

// main inicializa el enrutador y arranca el servidor HTTP en el puerto 8080
func main() {
	// Cargar variables de entorno
	config.Init()

	port := config.GetEnv("PORT", "8080")

	r := chi.NewRouter()
	r.Get("/ping", handler.Ping)

	fmt.Printf("Server running on http://localhost:%s\n", port)
	http.ListenAndServe(":"+port, r)
}
