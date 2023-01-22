package main

import (
	"net/http"
	"github.com/viniciushammett/models"
	"github.com/viniciushammett/routes"
)

func main() {
	routes.CarregaRotas()
	http.ListenAndServe(":8000", nil)
}
