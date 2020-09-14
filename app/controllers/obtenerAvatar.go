package controllers

import (
	"io"
	"net/http"
	"os"

	"github.com/federicoalonso/socialnetfa/app/dbconnection"
)

// ObtenerAvatar obtiene avatares
func ObtenerAvatar(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")

	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parámetro id", http.StatusBadRequest)
		return
	}

	perfil, err := dbconnection.BuscoPerfil(ID)
	if err != nil {
		http.Error(w, "Usuario no encontrado "+err.Error(), http.StatusBadRequest)
		return
	}

	OpenFile, err := os.Open("storage/app/public/avatar/" + perfil.Avatar)
	if err != nil {
		http.Error(w, "Imágen no encontrada "+err.Error(), http.StatusBadRequest)
		return
	}

	_, err = io.Copy(w, OpenFile)
	if err != nil {
		http.Error(w, "Imágen no se pudo enviar "+err.Error(), http.StatusBadRequest)
	}
}
