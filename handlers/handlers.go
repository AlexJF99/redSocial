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
	router.HandleFunc("/login", middlewares.CheckBD(routers.Login)).Methods("POST")
	router.HandleFunc("/profile", middlewares.CheckBD(middlewares.ValidJWT(routers.LookProfile))).Methods("GET")
	router.HandleFunc("/modifyProfile", middlewares.CheckBD(middlewares.ValidJWT(routers.ModifyProfile))).Methods("PUT")
	router.HandleFunc("/tweet", middlewares.CheckBD(middlewares.ValidJWT(routers.RecordTweet))).Methods("POST")
	router.HandleFunc("/readTweet", middlewares.CheckBD(middlewares.ValidJWT(routers.ReadTweet))).Methods("GET")
	router.HandleFunc("/deleteTweet", middlewares.CheckBD(middlewares.ValidJWT(routers.DeleteTweet))).Methods("DELETE")

	router.HandleFunc("/uploadAvatar", middlewares.CheckBD(middlewares.ValidJWT(routers.UploadAvatar))).Methods("POST")
	router.HandleFunc("/getAvatar", middlewares.CheckBD(middlewares.ValidJWT(routers.GetAvatar))).Methods("GET")
	router.HandleFunc("/uploadBanner", middlewares.CheckBD(middlewares.ValidJWT(routers.UploadBanner))).Methods("POST")
	router.HandleFunc("/getBanner", middlewares.CheckBD(middlewares.ValidJWT(routers.GetBanner))).Methods("GET")

	router.HandleFunc("/registRelation", middlewares.CheckBD(middlewares.ValidJWT(routers.RegistRelation))).Methods("POST")
	router.HandleFunc("/deleteRelation", middlewares.CheckBD(middlewares.ValidJWT(routers.DeleteRelation))).Methods("DELETE")
	router.HandleFunc("/getRelation", middlewares.CheckBD(middlewares.ValidJWT(routers.GetRelation))).Methods("GET")

	router.HandleFunc("/getAllUsers", middlewares.CheckBD(middlewares.ValidJWT(routers.GetAllUsers))).Methods("GET")

	router.HandleFunc("/readTweetsFollowers", middlewares.CheckBD(middlewares.ValidJWT(routers.ReadTweetFollowers))).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	log.Println("Listen and Serve on localhost:" + PORT)
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))
}
