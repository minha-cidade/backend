package db

import (
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
	"github.com/minha-cidade/backend/config"
)

var conn struct {
	Session  *mgo.Session
	Database *mgo.Database
}

// Conecta ao banco de dados de forma persistente
func Connect() (err error) {
	conn.Session, err = mgo.Dial(config.Get().DatabaseInfo)
	if err != nil {
		return
	}

	conn.Database = conn.Session.DB("despesas")
	return
}

func GetGastometroForYear(estado string, cidade string, ano int) (gastometro []*Gastometro, err error) {
	err = conn.Database.C("gastometro").Find(bson.M{
		"idCidade": cidade,
		"idEstado": estado,
		"ano":      ano,
	}).All(&gastometro)
	return
}
