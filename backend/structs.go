package main

type TokenInfo struct {
	Address           string         `json:"address"`
	TotalSupply       string         `json:"totalSupply"`
	Name              string         `json:"name"`
	Symbol            string         `json:"symbol"`
	Decimals          string         `json:"decimals"`
	Price             TokenPriceInfo `json:"price"`
	PublicTags        []string       `json:"publicTags,omitempty"`
	Owner             string         `json:"owner"`
	CountOps          int64          `json:"countOps"`
	TotalIn           int64          `json:"totalIn"`
	TotalOut          int64          `json:"totalOut"`
	TransfersCount    int64          `json:"transfersCount"`
	EthTransfersCount int64          `json:"ethTransfersCount,omitempty"`
	HoldersCount      int64          `json:"holdersCount"`
	IssuancesCount    int64          `json:"issuancesCount"`
	Image             string         `json:"image,omitempty"`
	Description       string         `json:"description,omitempty"`
	Website           string         `json:"website,omitempty"`
	LastUpdated       int64          `json:"lastUpdated"`
	Address_Type      string
    TopHolder         []string
}

type Holder struct {
    Address string `json:"address"`
    Balance float64 `json:"balance"`
    Share float64 `json:"share"`
}

type Holders struct {
    Holders []Holder `json:"holders"`
}

// type Input struct {
//     tokenAddr string `json:"tokenAddress"`
// }


type TokenPriceInfo struct {
	Rate            float64 `json:"rate"`
	Currency        string  `json:"currency"`
	Diff            float64 `json:"diff"`
	Diff7d          float64 `json:"diff7d"`
	Diff30d         float64 `json:"diff30d"`
	MarketCapUsd    float64 `json:"marketCapUsd"`
	AvailableSupply float64   `json:"availableSupply"`
	Volume24h       float64 `json:"volume24h"`
	Ts              int64   `json:"ts"`
}

type Operation struct {
	Timestamp int64   `json:"timestamp"`
	From      string  `json:"from"`
	To        string  `json:"to"`
	Hash      string  `json:"hash"`
	Value     float64 `json:"value"`
	Input     string  `json:"input"`
	Success   bool    `json:"success"`
	Token     string
}

type Operations []Operation
