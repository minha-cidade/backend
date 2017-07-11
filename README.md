# Back-end Minha Cidade em Golang
O back-end tem a função de expor a API que será usada pelo front-end web e
mobile do projeto Minha Cidade, desenvolvido na HackFest 2017. Usa como base
o banco de dados obtido do [Portal de Transparência de João Pessoa](http://transparencia.joaopessoa.pb.gov.br/),
administrado pela [Secretaria da Transparência Pública](http://www.joaopessoa.pb.gov.br/secretarias/setransp/).
A base de dados pode ser obtida na [página de download](http://transparencia.joaopessoa.pb.gov.br/download) (item *dados de despesas, receitas e entidades*)
do mesmo.

## Instalação

### Docker

    $ docker build -t backend .

#### Execução

    $ docker run -it --rm --name backend -p 8080:8080 backend

### Manual

    $ go install github.com/minha-cidade/backend

#### Execução

    $ backend
