package conf

import (
	"context"
	"fmt"
	"log"
	"os"

	db "development/application/fiance/library"

	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
)

// Esse Struct Vai armazenar as principais variaveis que vao ser usadas para recuperar os dados
type Connect struct {
	Conn *db.Queries
	Cxt  context.Context
}

// Essa função init vai carregar as variaveis do arquivo .ENV
func init() {
	if err := godotenv.Load(); err != nil {
		log.Fatal("Falhar ao carregar Variaveis de ambiente")
	}
}

// A funcao Conn vai estabelecer a conexao com o banco de dados e com a library
func Conn() Connect {
	// Url de conexão com o database
	url := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s", os.Getenv("DB_USER"), os.Getenv("DB_PASSWORD"), 
		os.Getenv("DB_HOST"), os.Getenv("DB_PORT"), os.Getenv("DB_NAME"))
	// Abrindo a conexao com a library
	conn, err := pgx.Connect(context.Background(), url)
	if err != nil {
		return Connect{}
	}

	// Papulando a Struct e retornando os dados
	return Connect{
		Conn: db.New(conn),
		Cxt: context.Background(),
	}
}
