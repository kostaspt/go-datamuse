package datamuse

import (
	"encoding/json"
	"net/http"
	"net/url"
	"time"
)

type Datamuse struct {
	ApiUrl *url.URL
}

type Results []struct {
	Word           string   `json:"word"`
	Score          int      `json:"score"`
	SyllablesCount uint     `json:"numSyllables,omitempty"`
	Tags           []string `json:"tags,omitempty"`
}

func New() *Datamuse {
	dm := new(Datamuse)

	url, _ := url.Parse("https://api.datamuse.com")

	dm.ApiUrl = url

	return dm
}

func (dm *Datamuse) Get() (Results, error) {
	client := &http.Client{Timeout: 13 * time.Second}

	resp, err := client.Get(dm.ApiUrl.String())

	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()

	results := &Results{}

	err = json.NewDecoder(resp.Body).Decode(results)

	return *results, err
}

func (dm *Datamuse) Hyperlink() string {
	return dm.ApiUrl.String()
}
