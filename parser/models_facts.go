package main

type Facts struct {
	Match       Match            `json:"match"`
	Teams       []TeamGroup      `json:"teams"`
	Players     []Player         `json:"players"`
	Rounds      []Round          `json:"rounds"`
	Stats       []PlayerStats    `json:"stats"`
	Combat      CombatFacts      `json:"combat"`
	Positioning PositioningFacts `json:"positioning"`
	Utility     UtilityFacts     `json:"utility"`
	Economy     EconomyFacts     `json:"economy"`
	Objective   ObjectiveFacts   `json:"objective"`
}

type CombatFacts struct {
	Kills       []Kill          `json:"kills"`
	Damages     []Damage        `json:"damages"`
	WeaponFires []WeaponFire    `json:"weaponFires"`
	Shots       []Shot          `json:"shots"`
	Trades      []TradeAnalysis `json:"trades"`
}

type PositioningFacts struct {
	PlayerStates []PlayerState      `json:"playerStates"`
	Support      []SupportAnalysis  `json:"support"`
	Rotations    []RotationAnalysis `json:"rotations"`
}

type UtilityFacts struct {
	Grenades []any `json:"grenades"`
	Flashes  []any `json:"flashes"`
	Smokes   []any `json:"smokes"`
	Molotovs []any `json:"molotovs"`
}

type EconomyFacts struct {
	Rounds []any `json:"rounds"`
}

type ObjectiveFacts struct {
	BombEvents []any `json:"bombEvents"`
}

func BuildFacts(result *Result) Facts {
	return Facts{
		Match:   result.Match,
		Teams:   result.Teams,
		Players: result.Players,
		Rounds:  result.Rounds,
		Stats:   result.Stats,
		Combat: CombatFacts{
			Kills:       result.Kills,
			Damages:     result.Damages,
			WeaponFires: result.WeaponFires,
			Shots:       result.Shots,
			Trades:      result.AI.Trades,
		},
		Positioning: PositioningFacts{
			PlayerStates: nil,
			Support:      result.AI.Support,
			Rotations:    result.AI.Rotations,
		},
		Utility: UtilityFacts{
			Grenades: []any{},
			Flashes:  []any{},
			Smokes:   []any{},
			Molotovs: []any{},
		},
		Economy: EconomyFacts{
			Rounds: []any{},
		},
		Objective: ObjectiveFacts{
			BombEvents: []any{},
		},
	}
}
