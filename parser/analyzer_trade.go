package main

import "fmt"

const (
	TradeWindowTicks = 320
	TradeMaxDistance = 500
)

func AnalyzeTrade(result *Result) {

	for i := 0; i < len(result.Kills); i++ {

		first := result.Kills[i]

		for j := i + 1; j < len(result.Kills); j++ {

			second := result.Kills[j]

			if second.Tick-first.Tick > TradeWindowTicks {
				break
			}

			if second.VictimSteamID64 != first.KillerSteamID64 {
				continue
			}

			if second.KillerTeam != first.VictimTeam {
				continue
			}

			result.AI.Trades = append(result.AI.Trades, TradeAnalysis{
				Round:    first.Round,
				Tick:     first.Tick,
				Victim:   first.Victim,
				Killer:   first.Killer,
				Trader:   second.Killer,
				Success:  true,
				Distance: 0,
			})

			fmt.Printf(
				"[TRADE] %s traded %s (enemy: %s)\n",
				second.Killer,
				first.Victim,
				first.Killer,
			)

			break
		}
	}
}
