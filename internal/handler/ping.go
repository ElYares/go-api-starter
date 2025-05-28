// Package handler contiene los controladores HTTP que corresponden a las rutas de la api.
package handler

import (
	"net/http"
)

// ping responde con pong  y se utiliza como endpoint de verificacion de salud
func Ping(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("pong"))
}
