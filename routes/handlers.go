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
	router.HandleFunc("/leoTweets", middlewares.ChequeoBD(middlewares.ValidoJWT(controllers.LeoTweet))).Methods("GET")
	router.HandleFunc("/borroTweets", middlewares.ChequeoBD(middlewares.ValidoJWT(controllers.BorroTweet))).Methods("DELETE")
	router.HandleFunc("/subirAvatar", middlewares.ChequeoBD(middlewares.ValidoJWT(controllers.SubirAvatar))).Methods("POST")
	router.HandleFunc("/subirBanner", middlewares.ChequeoBD(middlewares.ValidoJWT(controllers.SubirBanner))).Methods("POST")
	router.HandleFunc("/obtenerAvatar", middlewares.ChequeoBD(controllers.ObtenerAvatar)).Methods("GET")
	router.HandleFunc("/obtenerBanner", middlewares.ChequeoBD(controllers.ObtenerBanner)).Methods("GET")
	router.HandleFunc("/amistad", middlewares.ChequeoBD(middlewares.ValidoJWT(controllers.InsertarRelacion))).Methods("POST")
	router.HandleFunc("/bajaAmistad", middlewares.ChequeoBD(middlewares.ValidoJWT(controllers.BorroRelacion))).Methods("DELETE")
	router.HandleFunc("/consultaAmistad", middlewares.ChequeoBD(middlewares.ValidoJWT(controllers.ConsultaRelacion))).Methods("GET")
	router.HandleFunc("/verUsuarios", middlewares.ChequeoBD(middlewares.ValidoJWT(controllers.ListaUsuarios))).Methods("GET")
	router.HandleFunc("/verTweets", middlewares.ChequeoBD(middlewares.ValidoJWT(controllers.LeoTweetRelacion))).Methods("GET")

	PORT := os.Getenv("PORT")
	if PORT == "" {
		PORT = "8099"
	}

	//esto es lo que le da permiso a quien corresponda
	handler := cors.AllowAll().Handler(router)
	log.Fatal(http.ListenAndServe(":"+PORT, handler))

}
