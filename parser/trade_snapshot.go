package main

type DeathSnapshot struct {
	Tick  int
	Round int

	// Victim
	Victim        string
	VictimSteamID uint64

	VictimX float64
	VictimY float64
	VictimZ float64

	VictimViewX float32
	VictimViewY float32

	VictimWeapon string

	VictimHP    int
	VictimArmor int

	VictimMoney int

	VictimScoped  bool
	VictimDucking bool

	VictimFlashDuration float64

	VictimPlace string

	// Killer
	Killer        string
	KillerSteamID uint64

	KillerX float64
	KillerY float64
	KillerZ float64

	KillerViewX float32
	KillerViewY float32

	KillerWeapon string

	KillerHP    int
	KillerArmor int

	KillerMoney int

	KillerScoped  bool
	KillerDucking bool

	KillerFlashDuration float64

	KillerPlace string

	TradeCandidates []TradeCandidate
}

type TradeCandidate struct {
	Rank      int    `json:"rank"`
	Name      string `json:"name"`
	SteamID64 uint64 `json:"steamId"`

	X float64 `json:"x"`
	Y float64 `json:"y"`
	Z float64 `json:"z"`

	Distance         float64 `json:"distance"`
	HeightDifference float64 `json:"heightDifference"`

	Alive bool `json:"alive"`

	HP    int `json:"hp"`
	Armor int `json:"armor"`

	Weapon string `json:"weapon"`

	Money int `json:"money"`

	FlashDuration float64 `json:"flashDuration"`

	Scoped bool `json:"scoped"`

	Ducking bool `json:"ducking"`

	Velocity float64 `json:"velocity"`

	ViewX float32 `json:"viewX"`
	ViewY float32 `json:"viewY"`

	KillerViewX float32 `json:"killerViewX"`
	KillerViewY float32 `json:"killerViewY"`

	KillerWeapon string `json:"killerWeapon"`

	KillerMoney int `json:"killerMoney"`

	KillerX float64 `json:"killerX"`
	KillerY float64 `json:"killerY"`
	KillerZ float64 `json:"killerZ"`

	KillerHP    int `json:"killerHP"`
	KillerArmor int `json:"killerArmor"`

	KillerScoped bool `json:"killerScoped"`

	KillerFlashDuration float64 `json:"killerFlashDuration"`

	KillerPlace string `json:"killerPlace"`
}
