package main

import (
	"log"
	"net/http"
	"server/routes"
)

func main(){
	// Load CSS files into "/style/" route
	http.Handle("/style/", http.StripPrefix("/style/", http.FileServer(http.Dir("./client/style"))))
	// Same for JS
	http.Handle("/js/", http.StripPrefix("/js/", http.FileServer(http.Dir("./client/js"))))
	
	routes.Init()
	log.Fatal(http.ListenAndServe(":3000", nil))
}