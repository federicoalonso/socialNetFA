package controllers

import (
	"net/http"

	"github.com/federicoalonso/socialnetfa/app/dbconnection"

	"github.com/federicoalonso/socialnetfa/app/models"
)

// BorroRelacion controlador de la accion
func BorroRelacion(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "El parámetro ID es obligatorio", http.StatusBadRequest)
		return
	}

	var t models.Relacion
	t.UsuarioID = IDUsuario
	t.UsuarioRelacionID = ID

	status, err := dbconnection.BorroRelacion(t)
	if err != nil || !status {
		http.Error(w, "Error al eliminar la relacion "+err.Error(), http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)
}
