package db

import (
	"database/sql"
	"fmt"
	"os"

	_ "github.com/sijms/go-ora/v2"
)

func ConectaComBancoDeDados() *sql.DB {
	//connectionString := dbParams["username"] + ":" + dbParams["password"] + "@" + dbParams["server"] + ":" + dbParams["port"] + "/" + dbParams["service"]
	conexao := "you database"
	db, err := sql.Open("oracle", conexao)
	if err != nil {
		panic(fmt.Errorf("error in sql.Open: %w", err))
	}
	return db
}

func CriarTabelaProdutos(db *sql.DB, createTableStatement string) {
	_, err := db.Exec(createTableStatement)
	if err != nil {
		HandleError("create table", err)
	}

}

func HandleError(msg string, err error) {
	if err != nil {
		fmt.Println(msg, err)
		os.Exit(1)
	}
}
