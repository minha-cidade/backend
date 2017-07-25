package server

import (
	"errors"
	"github.com/gorilla/mux"
	"github.com/minha-cidade/backend/db"
	"net/http"
	"strconv"
)

var (
	ErrNotFound    = errors.New("Not found")
	ErrAnoInvalido = errors.New("Ano inválido")
)

func apiNotFound(w http.ResponseWriter, r *http.Request) {
	writeJsonError(w, http.StatusNotFound, ErrNotFound)
}

func apiCidades(w http.ResponseWriter, r *http.Request) {
	cidades, err := db.GetCidades()
	if err != nil {
		writeJsonError(w, http.StatusInternalServerError, err)
		return
	}

	writeJson(w, http.StatusOK, struct {
		Cidades []*db.Cidade `json:"cidades"`
	} { cidades })
}

func apiCidade(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)

	// Extrai as variáveis do request
	cidade := vars["cidade"]

	// Extrai as variáveis de parâmetro
	query := r.URL.Query()

	// Extrai a área
	area := query.Get("area")

	// Extrai o ano
	ano := 0
	if anoStr := query.Get("ano"); anoStr != "" {
		var err error
		ano, err = strconv.Atoi(anoStr)
		if err != nil {
			writeJsonError(w, http.StatusBadRequest, ErrAnoInvalido)
			return
		}
	}

	// Pega as informações do gastometro para esse request
	gastometro, err := db.SearchGastometro(cidade, area, ano)
	if err != nil {
		// TODO: englobar o bando de dados para retornar um erro mais significa-
		// tivo, como NotFound e InternalError(errDoMongo)
		writeJsonError(w, http.StatusInternalServerError, err)
		return
	}

	// Caso não tenha encontrado nenhum resultado, envia um resultado vazio
	if gastometro == nil {
		gastometro = []*db.Gastometro{ }
	}

	writeJson(w, http.StatusOK, struct {
		Gastometro []*db.Gastometro `json:"gastometro"`
	} { gastometro })
}
