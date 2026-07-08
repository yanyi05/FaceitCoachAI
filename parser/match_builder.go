package main

import dem "github.com/markus-wa/demoinfocs-golang/v5/pkg/demoinfocs"

func BuildMatch(
	parser dem.Parser,
	mapName string,
) Match {

	state := parser.GameState()

	match := Match{
		Map:      mapName,
		TickRate: int(parser.TickRate()),

		ScoreCT: state.TeamCounterTerrorists().Score(),
		ScoreT:  state.TeamTerrorists().Score(),
	}

	if match.ScoreCT > match.ScoreT {
		match.Winner = "CT"
	} else {
		match.Winner = "T"
	}

	for _, p := range state.Participants().All() {

		if p.SteamID64 == 0 {
			continue
		}

		match.Players = append(match.Players, Player{
			Name:      p.Name,
			SteamID64: p.SteamID64,
			PlayerID:  GetPlayerID(p.SteamID64),
			Team:      teamName(p.Team),
			IsBot:     p.IsBot,
		})
	}

	return match
}
