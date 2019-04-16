// CLI client for the ICNDB JSON API (http://www.icndb.com/api).
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"html"
	"log"
	"net/http"
)

// Ref: http://www.icndb.com/api/
const URLTemplate = "https://api.icndb.com/jokes/random/%d"

// APIResp represents an API response from the ICNDB.
type APIResp struct {
	Value []struct {
		Joke string
	}
}

func main() {

	// Accept a CLI flag specifying the number of jokes to print.
	nJokes := flag.Int("n", 1, "number of random jokes to retreive")
	flag.Parse()

	// Make an API request.
	url := fmt.Sprintf(URLTemplate, *nJokes)
	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}
	defer resp.Body.Close()

	// Decode the JSON API response into an APIResp struct.
	var ar APIResp
	err = json.NewDecoder(resp.Body).Decode(&ar)
	if err != nil {
		log.Fatal(err)
	}

	// Print the jokes from our API response.
	// We also unescape any HTML character entity references.
	// Ref: https://en.wikipedia.org/wiki/List_of_XML_and_HTML_character_entity_references#Character_entity_references_in_HTML
	for _, item := range ar.Value {
		fmt.Println(html.UnescapeString(item.Joke))
	}
}
