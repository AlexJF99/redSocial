package routers

import (
	"encoding/json"
	"net/http"
	"redSocial/bd"
	"redSocial/jwt"
	"redSocial/models"
	"time"
)

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("content-type", "application/json")

	var t models.User
	err := json.NewDecoder(r.Body).Decode(&t)

	if err != nil {
		http.Error(w, "Usuario y/o contraseña invalidos "+err.Error(), http.StatusBadRequest)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "El email del usuario es requerido "+err.Error(), http.StatusBadRequest)
		return
	}

	document, exist := bd.TryLogin(t.Email, t.Password)

	if !exist {
		http.Error(w, "Usuario y/o contraseñas inválidos ", http.StatusBadRequest)
		return
	}

	jwtKey, err := jwt.GenerateJWT(document)

	if err != nil {
		http.Error(w, "Ocurrio un error al intentar generar el Token correspondiente "+err.Error(), http.StatusBadRequest)
		return
	}

	resp := models.ResponseLogin{
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
