package main

import "sort"

func SortTradeCandidates(list []TradeCandidate) {

	sort.Slice(list, func(i, j int) bool {

		// 活着优先
		if list[i].Alive != list[j].Alive {
			return list[i].Alive
		}

		// 距离近优先
		if list[i].Distance != list[j].Distance {
			return list[i].Distance < list[j].Distance
		}

		// 高度差小优先
		return abs(list[i].HeightDifference) <
			abs(list[j].HeightDifference)
	})

	for i := range list {

		list[i].Rank = i + 1
	}
}

func abs(v float64) float64 {

	if v < 0 {
		return -v
	}

	return v
}
