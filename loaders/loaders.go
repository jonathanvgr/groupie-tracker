package loaders

import (
	"encoding/json"
	"io"
	"net/http"
)

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