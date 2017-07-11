package server

import (
	"errors"
	"github.com/dgrijalva/jwt-go"
	"github.com/gorilla/mux"
	"github.com/minha-cidade/backend/config"
	"github.com/minha-cidade/backend/db"
	"net/http"
	"strconv"
	"time"
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

func apiGetToken(w http.ResponseWriter, r *http.Request) {
	// Cria o token
	token := jwt.New(jwt.SigningMethodHS256)

	// Configura as informações contidas no token
	claims := token.Claims.(jwt.MapClaims)
	claims["expires"] = time.Now().Add(time.Hour * 24).Unix()

	// Assina o token com a chave secreta
	tokenString, err := token.SignedString(config.Get().TokenSecretKey)
	if err != nil {
		writeJsonError(w, http.StatusInternalServerError, err)
		return
	}

	// Envia token
	writeJson(w, http.StatusOK, struct {
		Token string `json:"token"`
	}{tokenString})
}
