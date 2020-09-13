package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/federicoalonso/socialnetfa/app/dbconnection"
	"github.com/federicoalonso/socialnetfa/app/models"
)

// Registro - crea en la BD el usuario
func Registro(w http.ResponseWriter, r *http.Request) {

	var t models.Usuario

	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Error en los datos recibidos "+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "El email de usuario es requerido", 400)
		return
	}

	if len(t.Email) < 6 {
		http.Error(w, "El largo de password debe ser mayor a 5", 400)
		return
	}

	_, encontrado, _ := dbconnection.ChequeoYaExisteUsuario(t.Email)
	if encontrado {
		http.Error(w, "El usuario ya existe", 400)
		return
	}

	_, status, err := dbconnection.InsertoRegistro(t)
	if err != nil {
		http.Error(w, "No se pudo guardar el usuario "+err.Error(), 500)
		return
	}

	if !status {
		http.Error(w, "No se pudo guardar el usuario", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
