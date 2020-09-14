package controllers

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/federicoalonso/socialnetfa/app/dbconnection"
	"github.com/federicoalonso/socialnetfa/app/models"
	"github.com/federicoalonso/socialnetfa/config"
)

// Login login de usuarios
func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")

	var t models.Usuario

	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Usuario y/o Contraseña invaálidos "+err.Error(), 400)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "El email de usuario es requerido", 400)
		return
	}

	documento, existe := dbconnection.IntentoLogin(t.Email, t.Password)

	if !existe {
		http.Error(w, "Usuario y/o Contraseña invaálidos", 400)
		return
	}

	jwtKey, err := config.GeneroJWT(documento)
	if err != nil {
		http.Error(w, "Ocurrió un error "+err.Error(), 400)
		return
	}

	resp := models.RespuestaLogin{
		Token: jwtKey,
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)

	expirationTime := time.Now().Add(24 * time.Hour)
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   jwtKey,
		Expires: expirationTime,
	})
}
