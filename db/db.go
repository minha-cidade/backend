package db

import (
	"github.com/minha-cidade/backend/config"
	"gopkg.in/mgo.v2"
	"gopkg.in/mgo.v2/bson"
)

var conn struct {
	Session    *mgo.Session
	Database   *mgo.Database
	Collection *mgo.Collection
}

// Conecta ao banco de dados de forma persistente
func Connect() (err error) {
	conn.Session, err = mgo.Dial(config.Get().DatabaseInfo)
	if err != nil {
		return
	}

	conn.Database = conn.Session.DB("despesas")
	conn.Collection = conn.Database.C("gastometro")
	return
}

func SearchGastometro(cidade string, area string, ano int) (
	gastometro []*Gastometro, err error) {

	query := bson.M{"idCidade": cidade}
	if area != "" {
		query["idArea"] = area
	}

	if ano != 0 {
		query["ano"] = ano
	}

	err = conn.Collection.Find(query).All(&gastometro)
	return
}

func GetCidades() (cidades []*Cidade, err error) {
	err = conn.Collection.Pipe([]bson.M{
		{
			"$group": bson.M{
				"_id": bson.M{
					"id":   "$idCidade",
					"nome": "$cidade",
				},
				"anos": bson.M{"$addToSet": "$ano"},
				"areas": bson.M{
					"$addToSet": bson.M{
						"id":   "$idArea",
						"nome": "$area",
					},
				},
			},
		},
		{
			"$project": bson.M{
				"_id":   false,
				"anos":  true,
				"areas": true,
				"id":    "$_id.id",
				"nome":  "$_id.nome",
			},
		},
	}).All(&cidades)
	return
}
