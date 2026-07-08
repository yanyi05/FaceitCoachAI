package main

import (
	"fmt"
	"os"
	"sort"
)

func AnalyzeSupport(
	ctx *AnalyzerContext,
	result *Result,
) []SupportAnalysis {
	var analyses []SupportAnalysis
	fmt.Fprintln(os.Stderr)
	fmt.Fprintln(os.Stderr, "========== Position Analyzer ==========")

	for _, kill := range result.Kills {

		if kill.Round == 0 {
			continue
		}

		player, ok := ctx.PlayerBySteam[kill.VictimSteamID64]
		if !ok {
			continue
		}

		state := GetPlayerStateAtTick(
			ctx,
			player.PlayerID,
			kill.Tick,
		)

		if state == nil {
			continue
		}

		analysis := SupportAnalysis{

			Round: kill.Round,

			Victim: player.Name,

			VictimSteamID: player.SteamID64,
		}

		nearestDistance := int(^uint(0) >> 1)

		nearestHeight := 0

		var nearestPlayer *Player

		for _, teammate := range result.Players {

			if teammate.Team != player.Team {
				continue
			}

			if teammate.PlayerID == player.PlayerID {
				continue
			}

			teammateState := GetPlayerStateAtTick(
				ctx,
				teammate.PlayerID,
				kill.Tick,
			)

			if teammateState == nil {
				continue
			}

			if !teammateState.Alive {
				continue
			}

			dist := Distance(
				*state,
				*teammateState,
			)

			height := HeightDifference(
				*state,
				*teammateState,
			)

			candidate := SupportCandidate{

				Player: teammate.Name,

				X: teammateState.X,

				Y: teammateState.Y,

				Z: teammateState.Z,

				SteamID: teammate.SteamID64,

				PlayerID: teammate.PlayerID,

				Distance: dist,

				HeightDifference: height,

				Alive: teammateState.Alive,
			}

			analysis.Candidates = append(
				analysis.Candidates,
				candidate,
			)

			if dist < nearestDistance {

				nearestDistance = dist

				nearestHeight = height

				nearestPlayer = &teammate
			}

		}

		if nearestPlayer != nil {

			sort.Slice(

				analysis.Candidates,

				func(i, j int) bool {

					return analysis.Candidates[i].Distance <
						analysis.Candidates[j].Distance

				},
			)

			analysis.SupportPlayer = analysis.Candidates[0].Player

			analysis.SupportSteamID = nearestPlayer.SteamID64

			analysis.Distance = nearestDistance

			analysis.HeightDifference = nearestHeight

			analyses = append(

				analyses,

				analysis,
			)
		}
	}

	for _, a := range analyses {

		fmt.Fprintln(os.Stderr, "Candidates:")

		for i, c := range a.Candidates {

			fmt.Fprintf(
				os.Stderr,

				"%d. %s | %d units | Height:%d units | XYZ:(%d,%d,%d)\n",

				i+1,

				c.Player,

				c.Distance,

				c.HeightDifference,

				c.X,

				c.Y,

				c.Z,
			)
		}

		fmt.Fprintf(
			os.Stderr,
			"Round %d | %s <- %s | Distance:%d units | Height:%d units\n",
			a.Round,
			a.Victim,
			a.SupportPlayer,
			a.Distance,
			a.HeightDifference,
		)
	}

	return analyses
}
