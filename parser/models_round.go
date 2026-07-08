package main

type Round struct {
	Number  int    `json:"number"`
	Winner  string `json:"winner"`
	ScoreCT int    `json:"scoreCT"`
	ScoreT  int    `json:"scoreT"`
}
