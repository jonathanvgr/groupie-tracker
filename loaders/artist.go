package loaders

import (
	"fmt"
	"html/template"
	"net/http"
	"server/constants"
)

type artistData struct {
	Artist constants.Artist
	Relations constants.Relations
}

func Artist(w http.ResponseWriter, r *http.Request, id string, t *template.Template) {
	// Get API data
	var fullData artistData

	getDataFrom(fmt.Sprintf("%s/%s", constants.API_Artist_URL, id), &fullData.Artist)
	getDataFrom(fmt.Sprintf("%s/%s", constants.API_Relation_URL, id), &fullData.Relations)

	// Send data to the web page
	pageData := constants.PageData{Data: fullData}
	t.Execute(w, pageData)
}
