package server

import (
	"errors"
	"github.com/gorilla/mux"
	"github.com/minha-cidade/backend/db"
	"net/http"
	"strconv"
)

var (
	NotFoundError    = errors.New("Not found")
	AnoInvalidoError = errors.New("Ano inválido")
)

func apiNotFound(w http.ResponseWriter, r *http.Request) {
	writeJsonError(w, http.StatusNotFound, NotFoundError)
}

func apiGastometro(w http.ResponseWriter, r *http.Request) {
	gastometro, err := db.GetGastometro()
	if err != nil {
		writeJsonError(w, http.StatusInternalServerError, err)
		return
	}

	writeJson(w, http.StatusOK, gastometro)
}

func apiArea(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	area := vars["area"]
	ano, err := strconv.Atoi(vars["ano"])
	if err != nil {
		writeJsonError(w, http.StatusBadRequest, AnoInvalidoError)
		return
	}

	info, err := db.GetInformacoesArea(area, ano)
	if err != nil {
		writeJsonError(w, http.StatusInternalServerError, err)
		return
	}

	writeJson(w, http.StatusOK, info)
}
