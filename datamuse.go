package datamuse

import (
	"encoding/json"
	"net/http"
	"net/url"
	"time"
)

// Datamuse is the wrapper of all package's requests.
type Datamuse struct {
	apiURL *url.URL
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
	u, _ := url.Parse("https://api.datamuse.com")
	return &Datamuse{
		apiURL: u,
	}
}

// Get creates a GET request to the apiURL and parses results.
func (dm *Datamuse) Get() (Results, error) {
	c := &http.Client{Timeout: 30 * time.Second}

	resp, err := c.Get(dm.apiURL.String())
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	var res Results
	err = json.NewDecoder(resp.Body).Decode(&res)

	return res, err
}

// URL returns the apiURL as string.
func (dm *Datamuse) URL() string {
	return dm.apiURL.String()
}
