// Package config proporciona funciones para cargar y acceder a las variables de entorno
package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

// Init carga las variables de entorno desde un archivo .env
func Init() {
	err := godotenv.Load()
	if err != nil {
		log.Println("No .env file found.")
	}
}

// GetEnv devuelve el valor de una variable de entorno o un valor por defecto si no existe.
func GetEnv(key string, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}

	return defaultValue

}
