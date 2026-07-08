package main

type AIAnalysis struct {
	Support []SupportAnalysis `json:"support"`

	Trades []TradeAnalysis `json:"trades"`

	Rotations []RotationAnalysis `json:"rotations"`
}

type TradeAnalysis struct {
	Round int `json:"round"`

	Tick int `json:"tick"`

	Victim        string `json:"victim"`
	VictimSteamID uint64 `json:"victimSteamId"`

	Killer        string `json:"killer"`
	KillerSteamID uint64 `json:"killerSteamId"`

	Trader        string `json:"trader,omitempty"`
	TraderSteamID uint64 `json:"traderSteamId"`

	ExpectedTrader string `json:"expectedTrader"`

	ExpectedTraderSteamID uint64 `json:"expectedTraderSteamId"`
	Success               bool   `json:"success"`

	TradeTimeTicks int `json:"tradeTimeTicks"`

	SupportDistance int `json:"supportDistance"`

	HeightDifference int `json:"heightDifference"`

	TradeQuality string `json:"tradeQuality"`

	Distance float64 `json:"distance"`
}

type SupportAnalysis struct {
	Round int `json:"round"`

	Victim string `json:"victim"`

	VictimSteamID uint64 `json:"victimSteamId"`

	SupportPlayer string `json:"supportPlayer"`

	SupportSteamID uint64 `json:"supportSteamId"`

	Distance int `json:"distance"`

	HeightDifference int `json:"heightDifference"`

	Candidates []SupportCandidate `json:"candidates"`
}
