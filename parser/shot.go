package main

type Shot struct {
	ShotIndex int `json:"shotIndex"`
	// Fire

	FireTick int `json:"fireTick"`

	Round int `json:"round"`

	Player string `json:"player"`

	SteamID64 uint64 `json:"steamId"`

	Weapon string `json:"weapon"`

	//--------------------------------------------------

	// Hit

	Hit bool `json:"hit"`

	Miss bool `json:"miss"`

	HitTick int `json:"hitTick"`

	TimeToDamage int `json:"timeToDamage"`

	Victim string `json:"victim"`

	VictimSteamID64 uint64 `json:"victimSteamId"`

	Damage int `json:"damage"`

	HitGroup string `json:"hitGroup"`

	//--------------------------------------------------

	// Kill

	Kill bool `json:"kill"`

	Headshot bool `json:"headshot"`

	Wallbang bool `json:"wallbang"`

	ThroughSmoke bool `json:"throughSmoke"`

	Blind bool `json:"blind"`

	NoScope bool `json:"noScope"`

	FlashAssist bool `json:"flashAssist"`

	Distance float32 `json:"distance"`

	Penetration int `json:"penetration"`

	BurstIndex int `json:"burstIndex"`

	BurstShot int `json:"burstShot"`

	TimeToKill int `json:"timeToKill"`
}
