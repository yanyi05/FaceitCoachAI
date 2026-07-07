package main

type Match struct {
	Map      string `json:"map"`
	TickRate int    `json:"tickRate"`
	ScoreCT  int    `json:"scoreCT"`
	ScoreT   int    `json:"scoreT"`
	Winner   string `json:"winner"`
}
