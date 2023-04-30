package routers

import (
	"encoding/json"
	"net/http"
	"redSocial/bd"
	"strconv"
)

func ReadTweetFollowers(w http.ResponseWriter, r *http.Request) {
	if len(r.URL.Query().Get("page")) < 1 {
		http.Error(w, "Debe enviar el parámetro page", http.StatusBadRequest)
		return
	}

	page, err := strconv.Atoi(r.URL.Query().Get("page"))
	if err != nil {
		http.Error(w, "Debe enviar el parámetro page como entero mayor a 0", http.StatusBadRequest)
		return
	}

	response, correct := bd.ReadTweetFollowers(IDUser, int64(page))
	if !correct {
		http.Error(w, "Error al leer los tweets", http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(response)
}
