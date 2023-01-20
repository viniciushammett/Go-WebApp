package main

import (
	"html/template"
	"net/http"
)

type Produto struct {
	Nome       string
	Descricao  string
	Preco      float64
	Quantidade int
}

var temp = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	http.HandleFunc("/", index)
	http.ListenAndServe(":8000", nil)
}

func index(w http.ResponseWriter, r *http.Request) {
	produtos := []Produto{
		{Nome: "Camiseta", Descricao: "Azul bem bonita", Preco: 39, Quantidade: 5},
		{"Tenis", "Confortavel como eu sempre quis", 89, 3},
		{"Fone", "Muito bom de bonito", 59, 2},
		{"Produto novo", "Muito bom mesmo", 1.99, 1},
	}

	temp.ExecuteTemplate(w, "index", produtos)

}
