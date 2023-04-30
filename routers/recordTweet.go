package routers

import (
	"encoding/json"
	"net/http"
	"redSocial/bd"
	"redSocial/models"
	"time"
)

func RecordTweet(w http.ResponseWriter, r *http.Request) {
	var message models.Tweet
	err := json.NewDecoder(r.Body).Decode(&message)

	regist := models.RecordTweet{
		UserID:  IDUser,
		Message: message.Message,
		Date:    time.Now(),
	}

	_, status, err := bd.InsertTweet(regist)
	if err != nil {
		http.Error(w, "Error to create tweet "+err.Error(), http.StatusBadRequest)
		return
	}
	if status == false {
		http.Error(w, "Regist not created ", http.StatusBadRequest)
		return
	}

	w.WriteHeader(http.StatusCreated)

}
