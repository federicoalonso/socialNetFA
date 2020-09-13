package middlewares

import (
	"net/http"

	"github.com/federicoalonso/socialnetfa/config"
)

// ChequeoBD permite conocer el estado de la bd antes de ejecutar las funciones
func ChequeoBD(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if !config.CheckConnection() {
			http.Error(w, "Conexion perdida con la Base de Datos", 500)
			return
		}
		next.ServeHTTP(w, r)
	}
}
