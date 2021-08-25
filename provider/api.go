package provider

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/kustov-an/define/common"
)

const baseUrl = "https://api.dictionaryapi.dev/api/v2/entries/en_US/"

type ApiProvider struct {
	Client *http.Client
}

var ErrorNetwork = errors.New("network error")
var ErrorMissingDefinition = errors.New("missing definition")

type Response struct {
	Meanings []struct {
		Definitions []struct {
			Definition string `json:"definition"`
		} `json:"definitions"`
	} `json:"meanings"`
}

func (r *Response) toArray() []string {
	result := []string{}
	for _, meaning := range r.Meanings {
		for _, definition := range meaning.Definitions {
			result = append(result, definition.Definition)
		}
	}
	return result
}

func (a *ApiProvider) Define(word string) (common.Definition, error) {
	url := baseUrl + word
	resp, err := a.Client.Get(url)
	if err != nil {
		return common.Definition{}, ErrorNetwork
	}
	defer resp.Body.Close()
	bytes, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return common.Definition{}, ErrorNetwork
	}

	response := []Response{}
	err = json.Unmarshal(bytes, &response)
	if err != nil {
		return common.Definition{}, ErrorMissingDefinition
	}
	return common.Definition{
		Meanings: response[0].toArray(),
	}, nil
}
