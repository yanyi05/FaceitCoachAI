package main

type Kill struct {
	Tick  int `json:"tick"`
	Round int `json:"round"`

	Killer          string `json:"killer"`
	KillerSteamID64 uint64 `json:"killerSteamId"`

	KillerTeam string `json:"killerTeam"`

	VictimTeam string `json:"victimTeam"`

	Victim          string `json:"victim"`
	VictimSteamID64 uint64 `json:"victimSteamId"`

	Assister          string `json:"assister"`
	AssisterSteamID64 uint64 `json:"assisterSteamId"`

	Weapon string `json:"weapon"`

	Headshot bool `json:"headshot"`

	Wallbang bool `json:"wallbang"`

	ThroughSmoke bool `json:"throughSmoke"`

	Blind bool `json:"blind"`

	NoScope bool `json:"noScope"`

	FlashAssist bool `json:"flashAssist"`

	Penetration int `json:"penetration"`

	Distance float32 `json:"distance"`
	Suicide  bool    `json:"suicide"`

	TeamKill bool `json:"teamKill"`
}

type Damage struct {
	Tick  int `json:"tick"`
	Round int `json:"round"`

	Attacker          string `json:"attacker"`
	AttackerSteamID64 uint64 `json:"attackerSteamId"`

	Victim          string `json:"victim"`
	VictimSteamID64 uint64 `json:"victimSteamId"`

	Weapon string `json:"weapon"`

	Damage int `json:"damage"`

	HealthDamage      int `json:"healthDamage"`
	HealthDamageTaken int `json:"healthDamageTaken"`
	ArmorDamageTaken  int `json:"armorDamageTaken"`

	ArmorDamage int `json:"armorDamage"`

	HealthRemaining int `json:"healthRemaining"`

	ArmorRemaining int `json:"armorRemaining"`

	HitGroup     string `json:"hitGroup"`
	FriendlyFire bool   `json:"friendlyFire"`
	Matched      bool   `json:"matched"`
}

type WeaponFire struct {
	Tick  int `json:"tick"`
	Round int `json:"round"`

	Player    string `json:"player"`
	SteamID64 uint64 `json:"steamId"`

	Weapon string `json:"weapon"`

	AmmoInMagazine int  `json:"ammoInMagazine,omitempty"`
	AmmoReserve    int  `json:"ammoReserve,omitempty"`
	Matched        bool `json:"matched"`
}
