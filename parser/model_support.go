package main

type SupportCandidate struct {
	Player string

	SteamID uint64

	PlayerID uint8
	X        int16
	Y        int16
	Z        int16

	Distance int

	HeightDifference int

	Alive bool
}
