package server

import (
	"bytes"
	"encoding/json"
	"log"
	"net/http"
)

func writeJson(w http.ResponseWriter, status int, data interface{}) {
	buf := new(bytes.Buffer)

	// Define o header application/json
	w.Header().Set("Content-Type", "application/json")

	// Tenta escrever para um buffer, caso contrário retorna um erro
	if err := json.NewEncoder(buf).Encode(data); err != nil {
		log.Println("Failed rendering JSON:", err)

		// Envia uma resposta sobre o erro
		w.WriteHeader(http.StatusInternalServerError)
		w.Write([]byte("{\"error\": \"Failed processing JSON\"}"))
		return
	}

	w.WriteHeader(status)
	w.Write(buf.Bytes())
}

func writeJsonError(w http.ResponseWriter, status int, err error) {
	writeJson(w, status, struct {
		Error string `json:"error"`
	}{err.Error()})
}

func chk(err error) {
	if err != nil {
		log.Fatalln(err)
	}
}
