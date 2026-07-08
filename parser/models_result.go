package main

type Metadata struct {
	// Parser
	ParserVersion string `json:"parserVersion"`

	// Demo
	Game              string `json:"game"`
	DemoFormat        string `json:"demoFormat"`
	DemoinfocsVersion string `json:"demoinfocsVersion"`

	// Units
	TimeUnit     string `json:"timeUnit"`
	DistanceUnit string `json:"distanceUnit"`
	PositionUnit string `json:"positionUnit"`
	VelocityUnit string `json:"velocityUnit"`

	// Demo Information
	TickRate int `json:"tickRate"`
}

type Result struct {
	Success bool `json:"success"`

	Metadata Metadata `json:"metadata"`

	Match Match `json:"match"`

	Players        []Player        `json:"players"`
	Teams          []TeamGroup     `json:"teams"`
	Kills          []Kill          `json:"kills"`
	DeathSnapshots []DeathSnapshot `json:"deathSnapshots"`
	Damages        []Damage        `json:"damages"`
	Rounds         []Round         `json:"rounds"`
	Stats          []PlayerStats   `json:"stats"`

	Message string `json:"message,omitempty"`

	WeaponFires   []WeaponFire   `json:"weaponFires"`
	PositionCache *PositionCache `json:"PositionCache"`
	Shots         []Shot         `json:"shots"`

	AI AIAnalysis `json:"analysis"`

	Facts Facts `json:"facts"`
}
