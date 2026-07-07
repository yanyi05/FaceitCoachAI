package main

func FindNearestPosition(playerID uint8, tick int, positions []PlayerState) *PlayerState {

	var best *PlayerState

	bestDiff := 999999999

	for i := range positions {

		p := &positions[i]

		if p.PlayerID != playerID {
			continue
		}

		diff := tick - p.Tick

		if diff < 0 {
			diff = -diff
		}

		if diff < bestDiff {

			bestDiff = diff

			best = p
		}
	}

	return best
}
