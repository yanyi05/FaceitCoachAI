package main

import (
	"fmt"
	"os"
	"sort"
)

func PrintTradeDebug(result *Result) {

	for _, snap := range result.DeathSnapshots {

		fmt.Fprintf(
			os.Stdout,
			"\n==============================\n",
		)

		fmt.Fprintf(
			os.Stdout,
			"Round %d\n",
			snap.Round,
		)

		fmt.Fprintf(
			os.Stdout,
			"Victim : %s\n",
			snap.Victim,
		)

		fmt.Fprintf(
			os.Stdout,
			"Killer : %s\n",
			snap.Killer,
		)

		fmt.Fprintf(
			os.Stdout,
			"Candidates:\n",
		)

		candidates := make([]TradeCandidate, 0)

		for _, c := range snap.TradeCandidates {

			if !c.Alive {
				continue
			}

			if c.Distance > 1200 {
				continue
			}

			if c.HeightDifference > 128 || c.HeightDifference < -128 {
				continue
			}

			candidates = append(candidates, c)
		}

		sort.Slice(candidates, func(i, j int) bool {

			return candidates[i].Distance < candidates[j].Distance

		})

		for _, c := range candidates {

			fmt.Fprintf(
				os.Stdout,
				"  %-18s Dist:%6.0f HU  Height:%5.0f HU\n",
				c.Name,
				c.Distance,
				c.HeightDifference,
			)
		}
	}
}
