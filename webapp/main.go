package main

import (
	"html/template"
	"net/http"
	_ "github.com/lib/pq"
)

func conectaComBancoDeDados() *sql.DB {
	conexao := "user=postgres dbname=alura_lojas password=12345678 host=localhost sslmode=disable"
	db, err := sql.Open("postgres", conexao)
	if err != nil {
		panic(err.Error())
	}
	return db
}

type Produto struct {
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	db := conectaComBancoDeDados()
	defer db.Close()
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	produtos := []Produto{
		{Nome: "Camiseta", Descricao: "Azul bem bonita", Preco: 39, Quantidade: 5},
		{"Tenis", "Confortavel", 89, 3},
		{"Fone", "Muito bom ", 59, 2},
		{"Produto novo", "Muito bom mesmo", 1.99, 1},
	}

	temp.ExecuteTemplate(w, "index", produtos)

}
