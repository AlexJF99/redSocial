package routers

import (
	"errors"
	"redSocial/bd"
	"redSocial/models"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
)

var Email string

var IDUser string

/*ProcessToken process token to extract its value */
func ProcessToken(tk string) (*models.Claim, bool, string, error) {
	miClave := []byte("MastersdelDesarrollo_GrupodeFacebook")
	claims := &models.Claim{}

	splitToken := strings.Split(tk, "Bearer")
	if len(splitToken) != 2 {
		return claims, false, string(""), errors.New("formato de token invalido")
	}

	tk = strings.TrimSpace(splitToken[1])

	tkn, err := jwt.ParseWithClaims(tk, claims, func(token *jwt.Token) (interface{}, error) {
		return miClave, nil
	})

	if err == nil {
		_, found, _ := bd.CheckIfUserExist(claims.Email)
		if found {
			Email = claims.Email
			IDUser = claims.ID.Hex()
		}
		return claims, found, IDUser, nil
	}
	if !tkn.Valid {
		return claims, false, string(""), errors.New("token invalid")
	}

	return claims, false, string(""), err

}
