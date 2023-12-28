package routes

import (
	"errors"
	"html/template"
	"net/http"
	"server/loaders"
	"strings"
)

func artist() {
	http.HandleFunc("/artist/", artistHandler)
}

// Handles dynamic URL
// if URL is only "/artist" => Redirect to root page
// if URL is "/artist/<artist ID>" => Fetch artist data according to ID in the URL
func artistHandler(w http.ResponseWriter, r *http.Request) {
    if r.URL.Path == "/artist/" {
		http.Redirect(w, r, "/", http.StatusPermanentRedirect)
		return
    }
	
	id, err := getArtistFromURL(r.URL.Path)
	if err != nil {
		http.Error(w, "Not Found", http.StatusNotFound)
		return
	} 

	pageTemplate, err := template.ParseFiles("client/artist.html")
	if err != nil {
		panic(err.Error())
	}

	loaders.Artist(w, r, id, pageTemplate)
}

func getArtistFromURL(url string) (string, error) {
	if !(strings.Contains(url, "/artist/")) {
		return "", errors.New("bad request")
	}

	id, found := strings.CutPrefix(url, "/artist/")
	if !found {
		return "", errors.New("not found")
	}

	return id, nil
}