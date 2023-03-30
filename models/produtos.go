package models

import (
	"fmt"
	"store/db"
)

type Produto struct {
	Id         int
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

func BuscaTodosOsProdutos() []Produto {
	db := db.ConectaComBancoDeDados()

	selectDeTodosOsProdutos, err := db.Query("select * from produtos")
	if err != nil {
		panic(err.Error())
	}

	p := Produto{}
	produtos := []Produto{}

	for selectDeTodosOsProdutos.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = selectDeTodosOsProdutos.Scan(&id, &nome, &descricao, &preco, &quantidade)
		if err != nil {
			panic(err.Error())
		}

		p.Id = id
		p.Nome = nome
		p.Descricao = descricao
		p.Preco = preco
		p.Quantidade = quantidade
		produtos = append(produtos, p)
	}
	defer db.Close()
	return produtos
}

func CriaNovoProduto(nome string, descricao string, preco float64, quantidade int) {
	database := db.ConectaComBancoDeDados()

	insereDadosNoBanco, err := database.Prepare("insert into produtos (nome, descricao, preco, quantidade) values (:nome,:descricao,:preco,:quantidade)")
	db.HandleError("prepare insert statement", err)
	sqlResult, err := insereDadosNoBanco.Exec(nome, descricao, preco, quantidade)
	db.HandleError("exec insert statement", err)
	rowCount, _ := sqlResult.RowsAffected()
	fmt.Println("Inserted number of rows = ", rowCount)
	defer database.Close()
}

func DeletaProduto(id int) {
	database := db.ConectaComBancoDeDados()

	deletaDadosNoBanco, err := database.Prepare("delete from produtos where id = :id")
	db.HandleError("prepare delete statement", err)
	sqlResult, err := deletaDadosNoBanco.Exec(id)
	db.HandleError("exec delete statement", err)
	rowCount, _ := sqlResult.RowsAffected()
	fmt.Println("deleted number of rows = ", rowCount)
	defer database.Close()
}

func EditaProduto(id string) Produto {
	database := db.ConectaComBancoDeDados()

	produtoBanco, err := database.Query("select * from produtos where id = :id", id)
	db.HandleError("query select", err)

	produtoAtualizar := Produto{}

	for produtoBanco.Next() {
		var id, quantidade int
		var nome, descricao string
		var preco float64

		err = produtoBanco.Scan(&id, &nome, &descricao, &preco, &quantidade)
		db.HandleError("Scan", err)

		produtoAtualizar.Id = id
		produtoAtualizar.Nome = nome
		produtoAtualizar.Descricao = descricao
		produtoAtualizar.Preco = preco
		produtoAtualizar.Quantidade = quantidade
	}

	defer database.Close()
	return produtoAtualizar
}

func AtualizaProduto(id int, nome string, descricao string, preco float64, quantidade int) {
	database := db.ConectaComBancoDeDados()

	atualizaDadosNoBanco, err := database.Prepare("update produtos set nome = :nome, descricao = :descricao, preco= :preco, quantidade= :quantidade where id = :id")
	db.HandleError("prepare update statement", err)
	sqlResult, err := atualizaDadosNoBanco.Exec(nome, descricao, preco, quantidade, id)
	db.HandleError("exec update statement", err)
	rowCount, _ := sqlResult.RowsAffected()
	fmt.Println("updated number of rows = ", rowCount)
	defer database.Close()
}
