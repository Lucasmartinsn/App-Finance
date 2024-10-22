# Application of Finance
Essa aplicação tem como por objetivo apresentar uma aplicação conjunta de Back e Front End
para interação e execulção de atividades financeiras. Essa aplicação foi deswnvolvida com intuido de 
estudo e aprimoramento de habilidades.

## Requisitos de Sistema
Golang 1.22^

PostgreSQL 14

## Instalação de Dependencias

### Intalacação do Migrate 

    git clone https://github.com/golang-migrate/migrate.git
    cd migrate
    make build
    mv ./migrate /usr/bin
Comandos de inicalização

   1 - Criando o Arquivos de banco de dados iniciais:
    
    migrate create -ext sql -dir db/migration -seq initial_tables
    
 2 - Execultando os Scripts gerados no passo acima.
 
    migrate -path db/migration -database "postgres://postgres:123@localhost:5432/fiance?sslmode=disable" -verbose up
    
### Instalando o Sqlc
Repositorio com a documentacao de instalacao

    https://docs.sqlc.dev/en/stable/overview/install.html

Para Instalar usando o Linux

    sudo snap install sqlc

Caso você altera os seguintes arquivos

    └── db
        ├── query
            ├── user.sql
            ├── account.sql
            └── category.sql
Para que as alterações seja executadas pelo Sqlc deve execultar comando abaixo.

    sqlc generate

### Intalação do sistema de monitoramento do estado da aplicação
Instalaçaõ do binarios.

    curl -sSfL https://raw.githubusercontent.com/air-verse/air/master/install.sh | sh -s -- -b $(go env GOPATH)/bin

Inicializando o Sistema:

    air init
    air
