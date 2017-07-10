package main

import (
	"github.com/minha-cidade/transparencia-crawler/transparencia"
	"github.com/minha-cidade/transparencia-crawler/manager"
	"github.com/minha-cidade/transparencia-crawler/db"
)

func main() {
	db.Conectar("localhost")

	m := manager.New()
	m.Add(&transparencia.JoaoPessoa{})
	m.Run()
}
