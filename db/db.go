package db

import (
	_ "github.com/lib/pq"
	"github.com/jmoiron/sqlx"
	"time"
	"errors"
)

var db *sqlx.DB;

/**
 * Conecta ao banco de dados
 */
func Connect(info string) {
	db = sqlx.MustOpen("postgres", info)
}

// Armazena as informações do gastômetro
type Gastometro struct {
	Area    string `json:"area"`
	Liquido float64 `json:"liquido"`
	Pago    float64 `json:"pago"`
	Alfa    float64 `json:"alfa"`
}

type Gastometros map[string]Gastometro

/**
 * Retorna as informações do gastometro
 */
func GetGastometro() (Gastometros, error) {
	valores := []Gastometro{}
	err := db.Select(&valores,
		`SELECT
			desc_func as area,
			SUM(
				CASE
					WHEN desc_tpmo='Liquidacao de Empenho'
						THEN abs(valo_movi)
					WHEN desc_tpmo='Estorno de Liquidacao de Empenho'
						THEN -abs(valo_movi)
					ELSE 0
				END)
				AS liquido,
			SUM(
				CASE
					WHEN desc_tpmo='Pagamento de Empenho'
						THEN abs(valo_movi)
					WHEN desc_tpmo='Estorno de Pagamento de Empenho'
						THEN -abs(valo_movi)
					ELSE 0
				END)
				AS pago
			FROM despesa_lei131
			WHERE ano_empe = $1
			GROUP BY desc_func
			ORDER BY desc_func`, time.Now().Year())
	if err != nil {
		return nil, err
	}

	gastometro := make(map[string]Gastometro)
	for _, valor := range valores {
		// Calcula o alpha
		now := time.Now()
		sec := now.
			Sub(time.Date(now.Year(), 0, 0, 0, 0, 0, 0, now.Location())).
			Seconds()

		// Reais gastos por segundo
		valor.Alfa = valor.Pago / float64(sec)
		gastometro[bancoDeDadosParaInterno(valor.Area)] = valor
	}

	return gastometro, nil
}

// Repasse anual
type RepasseAnual struct {
	Valor      float64  `json:"valor"`
	Entidade   string `json:"entidade"`
	Favorecido string `json:"favorecido"`
	Codigo     string `json:"codigo"`
}

// Informações de determinada área
type InformacoesArea struct {
	Liquido  float64 `json:"liquido"`
	Pago     float64       `json:"pago"`
	TopVinte []RepasseAnual `json:"top20"`
}

func GetInformacoesArea(area string, ano int) (i InformacoesArea, err error) {
	area = internoParaBancoDeDados(area)
	if area == "" {
		err = errors.New("Área inválida")
		return
	}

	err = db.Get(&i,
		`SELECT
			SUM(
				CASE
					WHEN desc_tpmo='Liquidacao de Empenho'
						THEN abs(valo_movi)
					WHEN desc_tpmo='Estorno de Liquidacao de Empenho'
						THEN -abs(valo_movi)
					ELSE 0
				END)
				AS liquido,
			SUM(
				CASE
					WHEN desc_tpmo='Pagamento de Empenho'
						THEN abs(valo_movi)
					WHEN desc_tpmo='Estorno de Pagamento de Empenho'
						THEN -abs(valo_movi)
					ELSE 0
				END)
				AS pago
			FROM despesa_lei131
			WHERE desc_func = $1
				AND ano_empe = $2
			LIMIT 1`, area, ano)
	if err != nil {
		return
	}

	// Preenche as informações do top 20
	err = db.Select(&i.TopVinte,
		`SELECT
			SUM(valo_empe) as valor,
			nome_enti as entidade,
			nome_forn as favorecido,
			CONCAT(nume_empe, '/', ano_empe) as codigo
		FROM despesa_lei131
		WHERE desc_tpmo='Liquidacao de Empenho'
			AND desc_func = $1
			AND ano_empe = $2
		GROUP BY nome_enti, nome_forn, codigo
		ORDER BY valor DESC
		LIMIT 20`, area, ano)
	return
}
