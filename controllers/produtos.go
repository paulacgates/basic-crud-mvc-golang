package controllers

import (
	"html/template"
	"log"
	"net/http"
	"store/models"
	"strconv"
)

var temp = template.Must(template.ParseGlob("templates/*.html"))

func Index(w http.ResponseWriter, r *http.Request) {
	todosOsProdutos := models.BuscaTodosOsProdutos()
	temp.ExecuteTemplate(w, "Index", todosOsProdutos)
}

func New(w http.ResponseWriter, r *http.Request) {
	temp.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		precoConvertido, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println("precoConvertido error:", err)
		}
		quantidadeConvertida, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Println("quantidadeConvertida error:", err)
		}
		models.CriaNovoProduto(nome, descricao, precoConvertido, quantidadeConvertida)
	}
	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	idConvertida, err := strconv.Atoi(id)
	if err != nil {
		log.Println("idConvertida error:", err)
	}
	models.DeletaProduto(idConvertida)
	http.Redirect(w, r, "/", 301)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	idProduto := r.URL.Query().Get("id")
	produto := models.EditaProduto(idProduto)
	temp.ExecuteTemplate(w, "Edit", produto)
}

func Update(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		id := r.FormValue("id")
		idConvertida, err := strconv.Atoi(id)
		if err != nil {
			log.Println("idConvertida error:", err)
		}
		nome := r.FormValue("nome")
		descricao := r.FormValue("descricao")
		preco := r.FormValue("preco")
		quantidade := r.FormValue("quantidade")

		precoConvertido, err := strconv.ParseFloat(preco, 64)
		if err != nil {
			log.Println("precoConvertido error:", err)
		}
		quantidadeConvertida, err := strconv.Atoi(quantidade)
		if err != nil {
			log.Println("quantidadeConvertida error:", err)
		}
		models.AtualizaProduto(idConvertida, nome, descricao, precoConvertido, quantidadeConvertida)
	}
	http.Redirect(w, r, "/", 301)

}
