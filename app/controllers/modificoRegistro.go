package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/federicoalonso/socialnetfa/app/dbconnection"
	"github.com/federicoalonso/socialnetfa/app/models"
)

// ModificoRegistro funcion para modificar los datos del usuario
func ModificoRegistro(w http.ResponseWriter, r *http.Request) {
	var t models.Usuario

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Datos incorrectos "+err.Error(), 400)
		return
	}

	var status bool
	status, err = dbconnection.ModificoRegistro(t, IDUsuario)

	if err != nil {
		http.Error(w, "Ocurrio un error al modificar "+err.Error(), 400)
		return
	}

	if !status {
		http.Error(w, "Ocurrio un error al modificar", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
