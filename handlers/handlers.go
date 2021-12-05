package handlers

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rasjsus/rastwit/middlewares"
	"github.com/rasjsus/rastwit/routers"
	"github.com/rs/cors"
)

func Handlers() {
	router := mux.NewRouter()

	router.HandleFunc("/register", middlewares.CheckDB(routers.Register)).Methods(http.MethodPost)
	router.HandleFunc("/login", middlewares.CheckDB(routers.Login)).Methods(http.MethodPost)
	router.HandleFunc("/perfil", middlewares.CheckDB(middlewares.ValidateJWT(routers.Perfil))).Methods(http.MethodGet)
	router.HandleFunc("/perfil", middlewares.CheckDB(middlewares.ValidateJWT(routers.UpdatePerfil))).Methods(http.MethodPut)
	router.HandleFunc("/tweet", middlewares.CheckDB(middlewares.ValidateJWT(routers.TweetRouter))).Methods(http.MethodPost)
	router.HandleFunc("/tweets", middlewares.CheckDB(middlewares.ValidateJWT(routers.GetTweets))).Methods(http.MethodGet)
	router.HandleFunc("/tweet", middlewares.CheckDB(middlewares.ValidateJWT(routers.DeleteTweet))).Methods(http.MethodDelete)

	router.HandleFunc("/avatar", middlewares.CheckDB(middlewares.ValidateJWT(routers.UploadAvatar))).Methods(http.MethodPost)
	router.HandleFunc("/avatar", middlewares.CheckDB(routers.GetAvatar)).Methods(http.MethodGet)
	router.HandleFunc("/banner", middlewares.CheckDB(middlewares.ValidateJWT(routers.UploadBanner))).Methods(http.MethodPost)
	router.HandleFunc("/banner", middlewares.CheckDB(routers.GetBanner)).Methods(http.MethodGet)

	router.HandleFunc("/relations", middlewares.CheckDB(middlewares.ValidateJWT(routers.InsertRelation))).Methods(http.MethodPost)
	router.HandleFunc("/relations", middlewares.CheckDB(middlewares.ValidateJWT(routers.DeleteRelation))).Methods(http.MethodDelete)
	router.HandleFunc("/relations", middlewares.CheckDB(middlewares.ValidateJWT(routers.GetRealation))).Methods(http.MethodGet) //FAlta testear en postman

	router.HandleFunc("/users", middlewares.CheckDB(middlewares.ValidateJWT(routers.GetUsers))).Methods(http.MethodGet) //FAlta testear en postman

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
