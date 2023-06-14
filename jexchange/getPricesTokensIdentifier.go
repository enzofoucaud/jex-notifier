package jexchange

import (
	"encoding/json"
	"io"
	"net/http"

	"github.com/rs/zerolog/log"
)

type PricesTokensIdentifier struct {
	Rate float64 `json:"rate"`
	Unit string  `json:"unit"`
}

func GetPricesTokensIdentifier(token string) (PricesTokensIdentifier, error) {
	//Encode the data
	var (
		url = "https://api.jexchange.io/prices/" + token
	)

	// Create a new request using http
	resp, err := http.Get(url)
	if err != nil {
		log.Err(err).Msg("error making request")
		return PricesTokensIdentifier{}, err
	}

	defer func() {
		_ = resp.Body.Close()
	}()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Err(err).Msg("error reading response body")
		return PricesTokensIdentifier{}, err
	}

	var p PricesTokensIdentifier
	err = json.Unmarshal(body, &p)
	if err != nil {
		log.Err(err).Msg("error unmarshalling data")
		return PricesTokensIdentifier{}, err
	}

	return p, nil
}
