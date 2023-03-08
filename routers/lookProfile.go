package routers

import (
	"encoding/json"
	"net/http"
	"redSocial/bd"
)

func LookProfile(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parámetro ID", http.StatusBadRequest)
		return
	}

	profile, err := bd.LookProfile(ID)

	if err != nil {
		http.Error(w, "Ocurrió un error al intentar buscar el registro "+err.Error(), http.StatusBadRequest)
	}

	w.Header().Set("context-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(profile)
}
