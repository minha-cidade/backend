package transparencia

import (
	"strconv"
	"encoding/csv"
_	"os/exec"
	"os"
	"log"
	"strings"
	"fmt"
	"github.com/shopspring/decimal"
)

type JoaoPessoa struct { }

const (
	downloadUrl = "http://transparencia.joaopessoa.pb.gov.br/sicoda/sql_dumps/lei131.csv.zip"
	zipPath     = "/tmp/despesas-jp.csv.zip"
	csvFilename = "despesa_lei131.csv"
	csvPath     = "/tmp/" + csvFilename
)

// Baixa o banco de dados de licitações e realiza o processamento
func (t *JoaoPessoa) Handle() []DespesaAnual {
	// Baixa o banco de dados e, caso sucesso, remove o arquivo no final
	// da função
	log.Println("-> Baixando banco de dados...")
	panicIfError(exec.Command("wget", downloadUrl, "-O", zipPath).Run())
	defer os.Remove(zipPath)

	// Descompacta o banco de dados e, caso sucesso, remove o arquivo no final
	// da função
	log.Println("-> Desempacotando arquivos...")
	panicIfError(exec.Command("unzip", zipPath, csvFilename, "-d", "/tmp").Run())
	defer os.Remove(csvPath)

	// Abre o arquivo CSV
	log.Println("-> Processando CSV...")
	file, err := os.Open(csvPath)
	panicIfError(err)
	defer file.Close()

	// Cria o leitor CSV que processará os dados
	reader := csv.NewReader(file)
	reader.FieldsPerRecord = -1 //54;
	reader.Comma = '|'

	// Calcula o total liquidado e pago para cada ano
	type chaveAreaAno struct { Area string; Ano int }
	totalAnual := make(map[chaveAreaAno]struct {
		Liquidado decimal.Decimal
		Pago      decimal.Decimal
	})

	for {
		record, err := reader.Read()
		if record == nil {
			break
		}

		panicIfError(err)

		// Pega a área e o ano do movimento financeiro para gerar a chave
		// única
		area := processarArea(record[25])
		ano, err := strconv.Atoi(record[6])
		panicIfError(err)

		// Calcula a chave área/ano
		chave := chaveAreaAno{ area, ano }

		// Realiza cálculos de valor movimentado
		tipo := record[20]
		valor, err := decimal.NewFromString(record[22])
		panicIfError(err)

		// Processa o tipo de operação
		total := totalAnual[chave]

		switch tipo {
		case "Liquidacao de Empenho":
			total.Liquidado = total.Liquidado.Add(valor.Abs())
		case "Estorno de Liquidacao de Empenho":
			total.Liquidado = total.Liquidado.Sub(valor.Abs())
		case "Pagamento de Empenho":
			total.Pago = total.Pago.Add(valor.Abs())
		case "Estorno de Pagamento de Empenho":
			total.Pago = total.Pago.Sub(valor.Abs())
		}

		totalAnual[chave] = total
	}

	totais := make([]DespesaAnual, 0, len(totalAnual))
	for chave, total := range totalAnual {
		liquidado, _ := total.Liquidado.Float64()
		pago, _ := total.Pago.Float64()

		totais = append(totais, DespesaAnual{
			Id: fmt.Sprintf("%s-%d", chave.Area, chave.Ano),
			Area: chave.Area,
			Ano: chave.Ano,
			Pago: pago,
			Liquidado: liquidado,
			Estado: "pb",
			Cidade: "joao-pessoa",
		})
	}

	return totais
}

// Retorna a sigla do estado
func (t *JoaoPessoa) Estado() string {
	return "pb"
}

// Retorna a cidade
func (t *JoaoPessoa) Cidade() string {
	return "joao-pessoa"
}

func panicIfError(err error) {
	if err != nil {
		log.Panicln(err)
	}
}

func processarArea(data string) string {
	return strings.ToLower(data)
}

func processarDinheiro(quantidade string) float64 {
	valor, err := strconv.ParseFloat(quantidade, 64)
	panicIfError(err)

	return valor;
}

/*
func formatarData(data string) time.Time {
	t, err := time.Parse("2006-01-02", data)
	panicIfError(err)

	return t
}
*/
