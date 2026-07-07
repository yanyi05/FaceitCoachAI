package main

type AnalyzerContext struct {
	PlayerBySteam map[uint64]*Player

	PlayerByID map[uint8]*Player

	PlayerStateByPlayer map[uint8][]PlayerState

	KillsByRound map[int][]Kill

	DamagesByRound map[int][]Damage
}
