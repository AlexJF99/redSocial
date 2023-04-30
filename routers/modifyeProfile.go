package routers

import (
	"encoding/json"
	"net/http"
	"redSocial/bd"
	"redSocial/models"
)

func ModifyProfile(w http.ResponseWriter, r *http.Request) {

	var t models.User

	err := json.NewDecoder(r.Body).Decode(&t)
	if err != nil {
		http.Error(w, "Datos Incorrectos "+err.Error(), 400)
		return
	}

	status, err := bd.ModifyRegist(t, IDUser)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar modificar el registro. Reintent again "+err.Error(), 400)
		return
	}
	if !status {
		http.Error(w, "No se ha logrado modificar el registro del usuario", 400)
		return
	}
	w.WriteHeader(http.StatusCreated)

}
