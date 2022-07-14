# User Service

Serviço será responsável pelo processo de cadastro, autenticação e autorização de usuários na plataforma, tendo como resposta um token JWT para ser tratado no front e no Gateway para possíveis controles de acesso.

### 📋 Pré-requisitos

```
- Go lang
- Docker
```

Após a instalação, é necessário baixar as dependencias seguindo o código abaixo:
```
$ go mod tidy
```

### 🔧 Start

Após clonar o repositório, iniciaremos nosso banco de dados com o docker:

```
docker-compose up
```

É esperado uma saída final como:

```
"database system is ready to accept connections"
```

O que indica que nosso banco de dados está configurado e disponível.

## ⚙️ Executando a API

Para executar a api é esperado um comando simples do Go, em um terminal na raiz do projeto:

```
go run main.go
```

Neste caso é esperado uma saída como:

```
"Listening and serving HTTP on :5000"
```

## 📦 Desenvolvimento

Para consumir esta API, temos tais endpoints:

- GET  - /user-api/usuarios  = Listar todos os usuários cadastrados, retornando status code 200;
- POST - /user-api/novo      = Cria um novo usuário, retornando status code 204;
- POST - /user-api/login     = Faz a autenticação do usuário, retornando um token JWT e um admin = true, caso o seja;

## 🛠️ Construído com

* [GoLang](https://go.dev/) - Linguagem de Programação GO
* [Gin](https://gin-gonic.com/) - Gin Web Framework
* [Docker](https://www.docker.com/) - Utilização de container
* [Postgres](https://www.docker.com/) - Postgres SQL na versão 14
