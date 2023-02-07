package coincap

import "fmt"

type assetsResponce struct {
	Assets    []Asset `json:"data"`
	Timestamp int64   `json:"timestamp"`
}

type assetResponce struct {
	Asset     Asset `json:"data"`
	Timestamp int64 `json:"timestamp"`
}

type Asset struct {
	ID                string `json:"id"`
	Rank              string `json:"rank"`
	Symbol            string `json:"symbol"`
	Name              string `json:"name"`
	Supple            string `json:"supple"`
	MaxSupply         string `json:"maxSupply"`
	MarketCapUSD      string `json:"marketCapUsd"`
	VolumeUsd24Hr     string `json:"volumeUsd24Hr"`
	PriceUsd          string `json:"priceUsd"`
	ChangePercent24Hr string `json:"changePercent24Hr"`
	Vwap24Hr          string `json:"vwap24Hr"`
	Explorer          string `json:"explorer"`
}

func (d Asset) Info() string {
	return fmt.Sprintf("[ID] %s | [RANK] %s | [SYMBOL] %s | [NAME] %s | [PRICE] %s",
		d.ID, d.Rank, d.Symbol, d.Name, d.PriceUsd)
}
