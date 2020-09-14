package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/federicoalonso/socialnetfa/app/dbconnection"
)

// VerPerfil es para ver el perfil
func VerPerfil(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parámetro Id", http.StatusBadRequest)
		return
	}
	perfil, err := dbconnection.BuscoPerfil(ID)

	if err != nil {
		http.Error(w, "Ocurrió un error al buscar el usuario "+err.Error(), 400)
		return
	}
	w.Header().Set("context-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(perfil)
}
