# Backend Minha Cidade em Go
## Introdução
O Backend Minha Cidade expõe uma API de fácil acesso aos dados coletados
pelo (Crawler Minha Cidade)[https://github.com/minha-cidade/crawler], provendo
uma forma simples e eficiente de analisar dados públicos referentes às despesas
municipais.

## Configuração
Uma vez com o repositório baixado em seu computador, copie o arquivo
`config.template.yml` para `config.yml` e edite seu conteúdo conforme é
explicado nele.

## Instalação
Na pasta do repositório, execute o seguinte comando para criar uma imagem docker
com o Backend

    $ docker build -t backend .

Depois de criado a imagem, execute o seguinte comando para executar a aplicação

    $ docker run -it --rm --name backend -p 8080:8080 backend
