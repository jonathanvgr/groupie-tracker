package loaders

import (
	"encoding/json"
	"io"
	"net/http"
)

// Format des donnees a envoyer a une page web
type PageData struct {
	Data any
}

// Get data from an URL and unmarshals it into an interface
func getDataFrom(url string, v interface{}) {
	// Get API data
	res, err := http.Get(url)
	if err != nil {
		panic(err.Error())
	}

	// Read body as an array of bytes
	defer res.Body.Close()
	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err.Error())
	}

	// Unmarshal (decodes as JSON) body data
	err = json.Unmarshal(body, &v)
	if err != nil {
		panic(err.Error())
	}
}