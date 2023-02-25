package handlers

import (
	"log"
	"net/http"
	"os"
	"redSocial/middlewares"
	"redSocial/routers"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

func Mapping() {
	router := mux.NewRouter()

	router.HandleFunc("/registro", middlewares.CheckBD(routers.Regist)).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	log.Println("Listen and Serve on localhost:" + PORT)
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
