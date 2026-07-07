package main

type AIAnalysis struct {
	Trades []TradeAnalysis `json:"trades"`
}

type TradeAnalysis struct {
	Round int `json:"round"`

	Tick int `json:"tick"`

	Victim string `json:"victim"`

	Killer string `json:"killer"`

	Trader string `json:"trader,omitempty"`

	Success bool `json:"success"`

	Distance float64 `json:"distance"`
}
