package jexchange

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/rs/zerolog/log"
)

type Token struct {
	Chart      bool   `json:"chart"`
	Decimals   int    `json:"decimals"`
	Identifier string `json:"identifier"`
	IsMetaEsdt bool   `json:"isMetaEsdt"`
	Name       string `json:"name"`
}

func GetTokens() ([]Token, error) {
	//Encode the data
	var (
		url = "https://api.jexchange.io/tokens"
	)

	// Create a new request using http
	resp, err := http.Get(url)
	if err != nil {
		log.Err(err).Msg("error making request")
		return []Token{}, err
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Err(err).Msg("error reading response body")
		return []Token{}, err
	}

	var t []Token
	err = json.Unmarshal(body, &t)
	if err != nil {
		log.Err(err).Msg("error unmarshalling data")
		return []Token{}, err
	}

	return t, nil
}
