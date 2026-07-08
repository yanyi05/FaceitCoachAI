package main

type RotationAnalysis struct {
	Round int

	Player string

	SteamID uint64

	StartTick int

	EndTick int

	StartX int16
	StartY int16
	StartZ int16

	EndX int16
	EndY int16
	EndZ int16

	Distance int

	DurationTicks int
}
