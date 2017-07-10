package server

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

func writeJson(w http.ResponseWriter, status int, data interface{}) {
	buf := new(bytes.Buffer)

	// Tenta escrever para um buffer, caso contrário retorna um erro
	if err := json.NewEncoder(buf).Encode(data); err != nil {
		log.Println("Failed rendering JSON:", err)
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("{\"error\": \"Failed processing JSON\"}"))
		return
	}

	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	w.WriteHeader(status)
	w.Write(buf.Bytes())
}

func writeJsonError(w http.ResponseWriter, status int, err error) {
	writeJson(w, status, struct {
		Error string `json:"error"`
	}{err.Error()})
}
