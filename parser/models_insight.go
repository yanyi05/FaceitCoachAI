package main

type PositionInsight struct {
	Round  int    `json:"round"`
	Player string `json:"player"`

	NearestTeammate string `json:"nearestTeammate"`

	Distance int `json:"distance"`

	CanTrade bool `json:"canTrade"`

	Reason string `json:"reason"`
}
