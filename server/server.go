package server

import (
	"github.com/auth0/go-jwt-middleware"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/handlers"
	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	"github.com/minha-cidade/backend/config"
	"github.com/minha-cidade/backend/db"
	"log"
	"net/http"
	"os"
)

func Start() {
	log.Println("Iniciando aplicação backend...")
	router := mux.NewRouter()

	// Conecta ao banco de dados
	log.Println("Conectando ao banco de dados...")
	db.Connect(config.Get().DatabaseInfo)

	// Autenticação via JWT
	var jwtMiddleware = jwtmiddleware.New(jwtmiddleware.Options{
		ValidationKeyGetter: func(_ *jwt.Token) (interface{}, error) {
			return config.Get().TokenSecretKey, nil
		},
		SigningMethod: jwt.SigningMethodHS256,
	})

	// Configura os middlewares
	apiMiddlewares := alice.New(jwtMiddleware.Handler)
	middlewares := alice.New(
		handlers.RecoveryHandler(),
		handlers.CORS(handlers.AllowedOrigins([]string{"*"})),
		stdoutLoggingHandler
	)

	// Api
	api := router.PathPrefix("/api").Subrouter()
	api.Handle("/gastometro", apiMiddlewares.Then(http.HandlerFunc(apiGastometro))).
		Methods("GET")
	api.Handle("/area/{area}/{ano}", apiMiddlewares.Then(http.HandlerFunc(apiArea))).
		Methods("GET")
	api.Handle("/get-token", http.HandlerFunc(apiGetToken)).
		Methods("GET")

	// Processa a página de 404
	api.NotFoundHandler = http.HandlerFunc(apiNotFound)

	// Escuta nesse endereço
	addr := config.Get().Address
	log.Printf("Listening... %s", addr)
	log.Fatalln(http.ListenAndServe(addr, middlewares.Then(router)))
}

func stdoutLoggingHandler(h http.Handler) http.Handler {
	return handlers.LoggingHandler(os.Stdout, h)
}
