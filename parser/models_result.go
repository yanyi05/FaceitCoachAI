package main

type Result struct {
	Success bool  `json:"success"`
	Match   Match `json:"match"`

	Players []Player      `json:"players"`
	Kills   []Kill        `json:"kills"`
	Damages []Damage      `json:"damages"` // ← 新增这里
	Rounds  []Round       `json:"rounds"`
	Stats   []PlayerStats `json:"stats"`

	Message string `json:"message,omitempty"`

	WeaponFires  []WeaponFire  `json:"weaponFires"`
	PlayerStates []PlayerState `json:"playerStates"`
	Shots        []Shot        `json:"shots"`

	AI AIAnalysis `json:"analysis"`
}
