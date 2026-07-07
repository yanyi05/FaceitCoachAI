package main

type PlayerState struct {
	Tick int

	PlayerID uint8

	X int16
	Y int16
	Z int16

	HP    uint8
	Armor uint8

	Alive bool

	ViewYaw   int16
	ViewPitch int16

	Velocity uint16
}

type TickFrame struct {
	Tick int `json:"tick"`

	Players []PlayerState `json:"players"`

	Events []string `json:"events,omitempty"`
}
