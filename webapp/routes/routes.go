package routes

import (
	"net/http"
	"github.com/viniciushammett/controllers"
)

func CarregaRotas(){
	http.HandleFunc("/", controllers.Index)
}