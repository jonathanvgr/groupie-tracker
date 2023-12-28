package loaders

import (
	"html/template"
	"net/http"
	"server/constants"
)

// This function gets the data from the API, converts it into a struct and send the converted data to the HTML template
// It's important to convert into a struct so we can use templates directives
func Index(w http.ResponseWriter, r *http.Request, t *template.Template) {
	var artists []constants.Artist
	getDataFrom(constants.API_Artist_URL, &artists)
	
	// Send data to the web page
	pageData := constants.PageData{Data: artists}
	t.Execute(w, pageData)
}