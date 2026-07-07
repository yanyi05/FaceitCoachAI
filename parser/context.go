package main

type AnalyzerContext struct {
	PlayerBySteam map[uint64]*Player

	PlayerByID map[uint8]*Player

	PlayerIDBySteam map[uint64]uint8

	PlayerStateByPlayer map[uint8][]PlayerState

	KillsByRound map[int][]Kill

	DamagesByRound map[int][]Damage
}
