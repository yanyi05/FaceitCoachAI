package main

import "fmt"

func AnalyzePosition(
	ctx *AnalyzerContext,
	result *Result,
) {

	fmt.Println()
	fmt.Println("========== Position Analyzer ==========")

	for _, kill := range result.Kills {

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

		nearestDistance := int(^uint(0) >> 1)

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

			dist := DistanceSquared(
				*state,
				*teammateState,
			)

			if dist < nearestDistance {

				nearestDistance = dist

				nearestPlayer = &teammate
			}
		}

		if nearestPlayer != nil {

			fmt.Printf(
				"Round %d | %s | Nearest: %s | Distance²: %d\n",
				kill.Round,
				kill.Victim,
				nearestPlayer.Name,
				nearestDistance,
			)
		}
	}
}
