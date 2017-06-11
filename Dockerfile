FROM golang:1.6

# Instala o backend direto do GIT
RUN go get github.com/minha-cidade/backend

# Define o endereço do servidor para o endereço local, porta 8080
ENV MINHACIDADE_BACKEND_ADDRESS ":8080"

# Define as informações do banco de dados para:
# (https://godoc.org/github.com/lib/pq#hdr-Connection_String_Parameters)
ENV MINHACIDADE_BACKEND_DB_INFO "host=127.0.0.1 port=5432 user=admin password=senha123 dbname=banco sslmode=disable"

# Exporta a porta 8080
EXPOSE 8080

# Executa
CMD ["backend", "run"]
