package routers

import (
	"encoding/json"
	"net/http"
	"redSocial/bd"
	"redSocial/models"
)

func GetRelation(w http.ResponseWriter, r *http.Request) {
	ID := r.URL.Query().Get("id")
	if len(ID) < 1 {
		http.Error(w, "Debe enviar el parámetro ID", http.StatusBadRequest)
		return
	}

	var t models.Relation
	t.UserID = IDUser
	t.UserRelationID = ID

	var resp models.ResponseGetRelation

	status, err := bd.GetRelation(t)
	if err != nil || !status {
		resp.Status = false
	} else {
		resp.Status = true
	}
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(resp)
}
