package routers

import (
	"net/http"
	"redSocial/bd"
)

func DeleteTweet(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parámetro id", http.StatusBadRequest)
		return
	}

	err := bd.BorroTweet(ID, IDUser)
	if err != nil {
		http.Error(w, "Ocurrio un error al intentar borrar el tweet "+err.Error(), http.StatusBadRequest)
		return
	}
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
}
