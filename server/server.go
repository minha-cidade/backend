package server

import (
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
	chk(db.Connect())

	// Configura os middlewares
	middlewares := alice.New(
		handlers.RecoveryHandler(),
		handlers.CORS(handlers.AllowedOrigins(config.Get().AllowedOrigins)))

	// Caso uma reverse proxy seja configurada
	if config.Get().ReverseProxy {
		middlewares = middlewares.Append(handlers.ProxyHeaders)
	}

	// Adiciona logging
	middlewares = middlewares.Append(stdoutLoggingHandler)

	// Api v1
	api := router.PathPrefix("/api/v1").Subrouter()
	api.Handle("/cidades",
		http.HandlerFunc(apiCidades)).
		Methods("GET")
	api.Handle("/cidades/{cidade}",
		http.HandlerFunc(apiCidade)).
		Methods("GET")

	// Processa a página de 404
	router.NotFoundHandler = http.HandlerFunc(apiNotFound)

	// Escuta nesse endereço
	addr := config.Get().Address
	log.Printf("Listening... %s", addr)
	log.Fatalln(http.ListenAndServe(addr, middlewares.Then(router)))
}

func stdoutLoggingHandler(h http.Handler) http.Handler {
	return handlers.LoggingHandler(os.Stdout, h)
}
