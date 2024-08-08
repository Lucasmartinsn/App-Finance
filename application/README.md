# App-Finance

## Intalacação do Migrate 
    git clone https://github.com/golang-migrate/migrate.git
    cd migrate
    make build
    mv ./migrate /usr/bin
## Comandos de inicalização
    1 - Criando o Arquivo DB iniciais migrate create -ext sql -dir db/migration -seq initial_tables
    
    2 - Execultando os Scripts gerados no passo acima. migrate -path db/migration -database "postgres://postgres:123@localhost:5432/fiance?sslmode=disable" -verbose up
    
    3 - Apos execultar os dois primeiros passos tudo ja esta pronto. Porem caso você precise alterar o arquivo de migration do arquivo .up basta execultar o comando sqlc generate