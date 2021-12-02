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

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8080"
	}
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
