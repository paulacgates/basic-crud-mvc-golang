package main

import (
	"store/routes"

	_ "github.com/sijms/go-ora/v2"
)

// const createTableStatement = "create table paulacg.teste(id Number(19) GENERATED BY DEFAULT ON NULL AS IDENTITY,nome varchar(100) not null,descricao varchar(100) not null,preco number not null,quantidade int not null,primary key(id))"

func main() {
	routes.CarregaRotas()
}
