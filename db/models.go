package db

// Transferência realizada entre um credor e um favorecido
type Transferencia struct {
	Favorecido string  `bson:"favorecido" json:"favorecido"`
	Credor     string  `bson:"credor" json:"credor"`
	Valor      float64 `bson:"valor" json:"valor"`
}

// Informações do gastometro
type Gastometro struct {
	Ano int `bson:"ano" json:"ano"`

	IDCidade string `bson:"idCidade" json:"idCidade"`
	IDEstado string `bson:"idEstado" json:"idEstado"`
	IDArea   string `bson:"idArea" json:"idArea"`

	Cidade string `bson:"cidade" json:"cidade"`
	Estado string `bson:"estado" json:"estado"`
	Area   string `bson:"area" json:"area"`

	Pago      float64 `bson:"pago" json:"pago"`
	Liquidado float64 `bson:"liquidado" json:"liquidado"`
	Empenhado float64 `bson:"empenhado" json:"empenhado"`

	// Top vinte transações realizadas com dinheiro público nesse espaço
	// de tempo
	TopVinte []Transferencia `bson:"topVinte" json:"topVinte"`
}
