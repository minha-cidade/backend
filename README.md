# Backend Minha Cidade em Go
## Introdução
O Backend Minha Cidade expõe uma API de fácil acesso aos dados coletados
pelo [Crawler Minha Cidade](https://github.com/minha-cidade/crawler), provendo
uma forma simples e eficiente de analisar dados públicos referentes às despesas
municipais.

## Configuração
A configuração é enviada através das seguintes variáveis de ambiente:

* `BACKEND_LISTEN_ADDRESS` (default: `":8080"`)

  Endereço o qual o servidor aceitará requests.


* `BACKEND_MONGO_CONNECTION_STRING` (default: `"mongodb://localhost"`)

  Endereço do banco de dados MongoDB seguindo o formato [Connection String](https://docs.mongodb.com/manual/reference/connection-string/).


* `BACKEND_CORS_ALLOWED_ORIGINS` (default: `"*"`)

  Endereços que o sevidor enviará no header [CORS](https://pt.wikipedia.org/wiki/Cross-origin_resource_sharing) separados
  por vírgula, sem espaço.

  Ex:

      "google.com,yahoo.com,bing.com"


## Instalação
Na pasta do repositório, execute o seguinte comando para criar uma imagem docker
com o Backend

    $ docker build -t backend .

Depois de criado a imagem, execute o seguinte comando para executar a aplicação

    $ docker run -it --rm --name backend -p 8080:8080 backend
