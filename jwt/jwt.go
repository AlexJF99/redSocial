package jwt

import (
	"redSocial/models"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
)

func GenerateJWT(t models.User) (string, error) {
	clave := []byte("MastersdelDesarrollo_GrupodeFacebook")

	payload := jwt.MapClaims{
		"email":    t.Email,
		"name":     t.Name,
		"lastName": t.LastName,
		"birthday": t.Birthday,
		"location": t.Location,
		"webSite":  t.WebSite,
		"_id":      t.ID.Hex(),
		"exp":      time.Now().Add(time.Hour * 24).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, payload)
	tokenStr, err := token.SignedString(clave)
	if err != nil {
		return tokenStr, err
	}
	return tokenStr, err
}
