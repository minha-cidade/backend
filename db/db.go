package db

import (
	"github.com/minha-cidade/transparencia-crawler/transparencia"
	"gopkg.in/mgo.v2"
	"log"
)

var config struct {
	Session  *mgo.Session
 	Database *mgo.Database
}

// Conecta ao banco de dados de forma persistente
func Conectar(data string) {
	var err error
	config.Session, err = mgo.Dial(data);
	if err != nil {
		log.Fatalln("Falha conectando ao mongodb:", err)
	}

	config.Database = config.Session.DB("despesas")
}

func EnviarDespesaAnual(d []transparencia.DespesaAnual) {
	config.Database.C("despesas").RemoveAll(nil)
	for _, v := range d {
		if err := config.Database.C("despesas").Insert(v); err != nil {
			log.Fatalln(err)
		}
	}
}
