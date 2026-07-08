package main

import (
	"fmt"
	"os"
)

const (
	TradeWindowTicks = 320
	TradeMaxDistance = 500
)

func AnalyzeTrade(

	result *Result,

	supports []SupportAnalysis,

) {

	for i := 0; i < len(result.Kills); i++ {

		first := result.Kills[i]
		if first.Round == 0 {
			continue
		}

		support := FindSupportByVictim(
			supports,
			first.VictimSteamID64,
			first.Round,
		)

		for j := i + 1; j < len(result.Kills); j++ {

			expectedTrader := ""

			expectedTraderSteamID := uint64(0)

			distance := 0
			height := 0

			if support != nil {

				expectedTrader = support.SupportPlayer

				expectedTraderSteamID = support.SupportSteamID

				distance = support.Distance

				height = support.HeightDifference
			}

			second := result.Kills[j]
			if second.KillerSteamID64 == second.VictimSteamID64 {
				continue
			}

			if second.Tick-first.Tick > TradeWindowTicks {
				break
			}

			if second.VictimSteamID64 != first.KillerSteamID64 {
				continue
			}

			if second.KillerTeam != first.VictimTeam {
				continue
			}

			if second.KillerSteamID64 == first.VictimSteamID64 {
				continue
			}

			if second.KillerSteamID64 == first.KillerSteamID64 {
				continue
			}

			quality := "Poor"

			reaction := second.Tick - first.Tick

			if reaction <= 32 {

				quality = "Excellent"

			} else if reaction <= 64 {

				quality = "Good"

			} else if reaction <= 128 {

				quality = "Late"

			}

			result.AI.Trades = append(result.AI.Trades, TradeAnalysis{

				Round: first.Round,

				Tick: first.Tick,

				Victim:        first.Victim,
				VictimSteamID: first.VictimSteamID64,

				Killer:        first.Killer,
				KillerSteamID: first.KillerSteamID64,

				Trader:        second.Killer,
				TraderSteamID: second.KillerSteamID64,

				ExpectedTrader: expectedTrader,

				ExpectedTraderSteamID: expectedTraderSteamID,

				Success: true,

				TradeTimeTicks:  second.Tick - first.Tick,
				TradeQuality:    quality,
				SupportDistance: distance,

				HeightDifference: height,
			})

			fmt.Fprintf(
				os.Stderr,

				"[TRADE] Round %d | Expected:%s | Actual:%s | Time:%d ticks\n",

				first.Round,

				expectedTrader,

				second.Killer,

				second.Tick-first.Tick,
			)

			break
		}
	}
}
