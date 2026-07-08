package main

type PositionFrame struct {
	Tick int

	SteamID64 uint64

	X float64
	Y float64
	Z float64

	ViewX float32
	ViewY float32

	HP int

	Armor int

	Money int

	Weapon string

	FlashDuration float64

	Scoped bool

	Ducking bool

	Walking bool

	Airborne bool

	Reloading bool

	Alive bool
}

func (c *PositionCache) AddFrame(frame PositionFrame) {
	// 已废弃，由 BuildPositionCache() 负责建立 PositionCache。
}
