package routes

import (
	"log"
	"net/http"
	"os"

	"github.com/gorilla/mux"
	"github.com/rs/cors"

	"github.com/federicoalonso/socialnetfa/app/controllers"
	"github.com/federicoalonso/socialnetfa/app/middlewares"
)

// Routes controla las rutas
func Routes() {
	router := mux.NewRouter()

	router.HandleFunc("/registro", middlewares.ChequeoBD(controllers.Registro)).Methods("POST")
	router.HandleFunc("/login", middlewares.ChequeoBD(controllers.Login)).Methods("POST")
	router.HandleFunc("/perfil", middlewares.ChequeoBD(middlewares.ValidoJWT(controllers.VerPerfil))).Methods("GET")
	router.HandleFunc("/modificarPerfil", middlewares.ChequeoBD(middlewares.ValidoJWT(controllers.ModificoRegistro))).Methods("PUT")
	router.HandleFunc("/tweet", middlewares.ChequeoBD(middlewares.ValidoJWT(controllers.GraboTweet))).Methods("POST")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8099"
	}

	//esto es lo que le da permiso a quien corresponda
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
