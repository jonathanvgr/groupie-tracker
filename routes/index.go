package routes

import (
	"html/template"
	"net/http"
	"server/loaders"
)

func index(){
	http.HandleFunc("/", indexHandler)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	pageTemplate, err := template.ParseFiles("client/index.html")
	if err != nil {
		panic(err.Error())
	}

	loaders.Index(w, r, pageTemplate)
}
