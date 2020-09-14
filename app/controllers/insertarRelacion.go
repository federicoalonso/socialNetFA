package controllers

import (
	"net/http"

	"github.com/federicoalonso/socialnetfa/app/dbconnection"

	"github.com/federicoalonso/socialnetfa/app/models"
)

// InsertarRelacion controlador de la accion
func InsertarRelacion(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "El parámetro ID es obligatorio", http.StatusBadRequest)
		return
	}

	var t models.Relacion
	t.UsuarioID = IDUsuario
	t.UsuarioRelacionID = ID

	status, err := dbconnection.InsertoRelacion(t)
	if err != nil || !status {
		http.Error(w, "Error al crear la relacion "+err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
