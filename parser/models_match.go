package main

type Match struct {
	Map string

	TickRate int

	ScoreCT int

	ScoreT int

	Winner string

	Players []Player

	Rounds []Round

	Kills []Kill

	DeathSnapshots []DeathSnapshot

	Damages []Damage

	Shots []Shot

	CurrentRound int

	CurrentTick int

	PositionCache PositionCache
}
