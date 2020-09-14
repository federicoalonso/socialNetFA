package dbconnection

import (
	"github.com/federicoalonso/socialnetfa/app/models"
	"golang.org/x/crypto/bcrypt"
)

// IntentoLogin me verifica si se loguea
func IntentoLogin(email string, password string) (models.Usuario, bool) {
	usu, existe, _ := ChequeoYaExisteUsuario(email)

	if !existe {
		return usu, false
	}

	passwordBytes := []byte(password)
	passwordBD := []byte(usu.Password)

	err := bcrypt.CompareHashAndPassword(passwordBD, passwordBytes)

	if err != nil {
		return usu, false
	}

	return usu, true
}
