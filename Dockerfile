FROM golang:latest

# Define diretório de trabalho
WORKDIR /go/src/app

# Copia os dados do arquivo pra imagem
COPY . .

# Baixa as dependências e instala o software
RUN go-wrapper download
RUN go-wrapper install

# Expõe a porta
EXPOSE 8080

# Executa
CMD ["go-wrapper", "run"]
