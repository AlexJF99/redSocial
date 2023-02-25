package routers

import (
	"encoding/json"
	"net/http"
	"redSocial/bd"
	"redSocial/models"
)

func Regist(w http.ResponseWriter, r *http.Request) {
	var t models.User
	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Error en los datos recibidos "+err.Error(), http.StatusBadRequest)
		return
	}

	if len(t.Email) == 0 {
		http.Error(w, "El email de usuario es requerido", 400)
		return
	}

	if len(t.Password) < 6 {
		http.Error(w, "El password debe de ser al menos de 6 caracteres", 400)
		return
	}

	_, exist, _ := bd.CheckIfUserExist(t.Email)
	if exist {
		http.Error(w, "Ya existe un usuario con ese email", http.StatusBadRequest)
		return
	}

	_, status, err := bd.InsertRegist(t)
	if err != nil {
		http.Error(w, "Ocurrió un error al intentar realizar el registro de usuario "+err.Error(), http.StatusBadRequest)
		return
	}

	if !status {
		http.Error(w, "No se ha logrado insertar el registro del usuario", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
