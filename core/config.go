package core

import (
	"log"

	"github.com/joho/godotenv"
)

// LoadConfig carga las configuraciones desde el archivo .env y retorna las variables requeridas.
func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Println("Error cargando el archivo .env (usando variables de entorno del sistema)")
	}
}
