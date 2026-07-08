package main

type TickFrame struct {
	Tick int `json:"tick"`

	Players []PlayerState `json:"players"`

	Events []string `json:"events,omitempty"`
}

type PositionCache struct {
	Frames map[int]*TickFrame
}
