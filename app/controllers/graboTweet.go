package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/federicoalonso/socialnetfa/app/dbconnection"
	"github.com/federicoalonso/socialnetfa/app/models"
)

// GraboTweet hace lo que se llama
func GraboTweet(w http.ResponseWriter, r *http.Request) {
	var mensaje models.Tweet
	err := json.NewDecoder(r.Body).Decode(&mensaje)
	if err != nil {
		http.Error(w, "Error en los datos recibidos "+err.Error(), 400)
		return
	}

	registro := models.GraboTweet{
		UserID:  IDUsuario,
		Mensaje: mensaje.Mensaje,
		Fecha:   time.Now(),
	}

	_, status, err2 := dbconnection.InsertoTweet(registro)

	if err2 != nil {
		http.Error(w, "Ocurrió un error al intentar insertar el registro "+err2.Error(), 400)
		return
	}

	if !status {
		http.Error(w, "Ocurrió un error al intentar insertar el registro", 400)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
