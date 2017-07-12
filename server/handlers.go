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
	vars := mux.Vars(r)

	// Extrai as variáveis do request
	estado := vars["estado"]
	cidade := vars["cidade"]

	// Caso o ano seja inválido, retorna uma mensagem de erro dizendo que esse
	// valor não existe
	ano, err := strconv.Atoi(vars["ano"])
	if err != nil {
		writeJsonError(w, http.StatusBadRequest, AnoInvalidoError)
		return
	}

	// Pega as informações do gastometro para esse request
	gastometro, err := db.GetGastometroForYear(estado, cidade, ano)
	if err != nil {
		// TODO: englobar o bando de dados para retornar um erro mais significa-
		// tivo, como NotFound e InternalError(errDoMongo)
		writeJsonError(w, http.StatusInternalServerError, err)
		return
	}

	writeJson(w, http.StatusOK, struct {
		Gastometro []*db.Gastometro `json:"gastometro"`
	} { gastometro })
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
