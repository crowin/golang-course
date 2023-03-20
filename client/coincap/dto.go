package coincap

import (
	"encoding/json"
	"io"
	"net/http"
)

type response[T any] struct {
	Timestamp int64 `json:"timestamp"`
	Data      T     `json:"data"`
}

type Asset struct {
	Id                string `json:"id"`
	Rank              string `json:"rank"`
	Symbol            string `json:"symbol"`
	Name              string `json:"name"`
	Supply            string `json:"supply"`
	MaxSupply         string `json:"maxSupply"`
	MarketCapUsd      string `json:"marketCapUsd"`
	VolumeUsd24Hr     string `json:"volumeUsd24Hr"`
	PriceUsd          string `json:"priceUsd"`
	ChangePercent24Hr string `json:"changePercent24Hr"`
	Vwap24Hr          string `json:"vwap24Hr"`
}

func (c *response[T]) toJson(r http.Response) (*response[T], error) {
	defer r.Body.Close() //look at client handler

	b, err := io.ReadAll(r.Body)

	if err != nil {
		return nil, err
	}

	var obj *response[T]
	if err = json.Unmarshal(b, &obj); err != nil {
		return nil, err
	}
	return obj, nil
}
