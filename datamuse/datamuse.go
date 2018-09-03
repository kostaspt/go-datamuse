package datamuse

import (
	"encoding/json"
	"net/http"
	"net/url"
	"time"
)

// Datamuse is the wrapper of all package's requests.
type Datamuse struct {
	APIURL *url.URL
}

// Results represent a list of Datamuse's results.
type Results []struct {
	Word           string   `json:"word"`
	Score          int      `json:"score"`
	SyllablesCount uint     `json:"numSyllables,omitempty"`
	Tags           []string `json:"tags,omitempty"`
}

// New create a new Datamuse instance.
func New() *Datamuse {
	dm := new(Datamuse)

	url, _ := url.Parse("https://api.datamuse.com")

	dm.APIURL = url

	return dm
}

// Get creates a GET request to the APIURL and parses results.
func (dm *Datamuse) Get() (Results, error) {
	client := &http.Client{Timeout: 13 * time.Second}

	resp, err := client.Get(dm.APIURL.String())

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	results := &Results{}

	err = json.NewDecoder(resp.Body).Decode(results)

	return *results, err
}

// Hyperlink returns the APIURL as string.
func (dm *Datamuse) Hyperlink() string {
	return dm.APIURL.String()
}
