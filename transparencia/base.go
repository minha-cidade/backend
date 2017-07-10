package transparencia

import (
	"time"
)

// Informação de uma transferência bancária realizada pelo município
type Transferencia struct {
	Favorecido    string    `json:"favorecido" bson:"favorecido"`
	Valor         float64   `json:"valor" bson:"valor"`
	Data          time.Time `json:"data" bson:"data"`
}

// Despesa anual de uma determinada área para um determinado ano
type DespesaAnual struct {
	Id        string  `json:"-" bson:"_id"`
	Ano       int     `json:"ano" bson:"ano"`
	Area      string  `json:"area" bson:"area"`
	Estado    string  `json:"estado" bson:"estado"`
	Cidade    string  `json:"cidade" bson:"cidade"`
	// Quantidade total em reais liquidada durante o ano
	Liquidado float64 `json:"liquidado" bson:"liquidado"`
	// Quantidade total de reais paga durante o ano
	Pago      float64 `json:"pago" bson:"pago"`

	// Top vinte transações realizadas para essa determinada área
	TopVinte  []Transferencia `json:"top20" bson:"topVinte"`
}

// Interface que toda transparência indexada deve processar
type Transparencia interface {
	Handle() []DespesaAnual
	Estado() string
	Cidade() string
}
