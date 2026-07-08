package main

type Player struct {
	Name      string `json:"name"`
	SteamID64 uint64 `json:"steamId"`
	Team      string `json:"team"`
	TeamGroup string `json:"teamGroup"`
	PlayerID  uint8  `json:"playerId"`
	IsBot     bool   `json:"isBot"`
}

type PlayerStats struct {
	Name      string `json:"name"`
	SteamID64 uint64 `json:"steamId"`

	Kills     int `json:"kills"`
	Deaths    int `json:"deaths"`
	Assists   int `json:"assists"`
	Headshots int `json:"headshots"`

	ADR             float64 `json:"adr"`
	KAST            float64 `json:"kast"`
	Rating          float64 `json:"rating"`
	Accuracy        float64 `json:"accuracy"`
	HeadshotPercent float64 `json:"headshotPercent"`

	Damage      int `json:"damage"`
	DamageTaken int `json:"damageTaken"`

	ShotsFired int `json:"shotsFired"`
	ShotsHit   int `json:"shotsHit"`

	EntryKills    int     `json:"entryKills"`
	EntryDeaths   int     `json:"entryDeaths"`
	EntryAttempts int     `json:"entryAttempts"`
	EntrySuccess  float64 `json:"entrySuccess"`

	AverageTTD float64 `json:"averageTTD"`

	TotalTTDTicks int `json:"totalTTDTicks"`

	TTDCount   int `json:"ttdCount"`
	TradeKills int `json:"tradeKills"`

	TradeDeaths int `json:"tradeDeaths"`

	TradeSuccess float64 `json:"tradeSuccess"`

	OpeningKills  int `json:"openingKills"`
	OpeningDeaths int `json:"openingDeaths"`
}
