package server

import (
	"github.com/gorilla/mux"
	"github.com/minha-cidade/backend/db"
	"log"
	"net/http"
	"os"
)

func getConfig() (addr string, info string) {
	// Pega o endereço de bind do servidor
	addr = os.Getenv("MINHACIDADE_BACKEND_ADDR")
	if addr == "" {
		addr = ":8080"
	}

	info = os.Getenv("MINHACIDADE_BACKEND_DB_INFO")
	if info == "" {
		log.Fatalln("Invalid database configuration")
	}

	return
}

func Start() {
	log.Println("Iniciando aplicação backend...")
	router := mux.NewRouter()

	// Lê a configuração
	addr, info := getConfig()

	// Conecta ao banco de dados
	log.Println("Conectando ao banco de dados...")
	db.Connect(info)

	// Api
	api := router.PathPrefix("/api").Subrouter()
	api.HandleFunc("/gastometro", apiGastometro).Methods("GET")
	api.HandleFunc("/area/{area}/{ano}", apiArea).Methods("GET")

	// Processa a página de 404
	api.NotFoundHandler = http.HandlerFunc(apiNotFound)

	// Escuta nesse endereço
	log.Printf("Listening... %s", addr)
	log.Fatalln(http.ListenAndServe(addr, router))
}
