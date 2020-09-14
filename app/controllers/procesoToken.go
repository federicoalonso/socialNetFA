package controllers

import (
	"errors"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/federicoalonso/socialnetfa/app/dbconnection"
	"github.com/federicoalonso/socialnetfa/app/models"
)

// Email es el email de usuario
var Email string

// IDUsuario es el id, los guardo para no hacerlo siempre el proceso
var IDUsuario string

// ProcesoToken me dice si esta autorizado
func ProcesoToken(tk string) (*models.Claim, bool, string, error) {
	miClave := []byte("AcaHgoUnaClavePropiaPAraQeuNoLaDecifren")
	claims := &models.Claim{}

	/*splitToken := strings.Split(tk, "Bearer")

	if len(splitToken) != 2 {
		// los mensajes de error en minuscula sin caracteres raros
		return claims, false, string(""), errors.New("formato de token invalido")
	}*/

	//tk = strings.TrimSpace(splitToken[1])
	tk = strings.TrimSpace(tk)

	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return miClave, nil
	})
	if err == nil {
		_, existe, _ := dbconnection.ChequeoYaExisteUsuario(claims.Email)
		if existe {
			Email = claims.Email
			IDUsuario = claims.ID.Hex()
		}
		return claims, existe, IDUsuario, nil
	}
	if !tkn.Valid {
		return claims, false, string(""), errors.New("token invalido")
	}

	return claims, false, string(""), err
}
