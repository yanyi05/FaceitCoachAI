package main

func GetPlayerStateAtTick(
	ctx *AnalyzerContext,
	playerID uint8,
	tick int,
) *PlayerState {

	states := ctx.PlayerStateByPlayer[playerID]

	if len(states) == 0 {
		return nil
	}

	var last *PlayerState

	for i := range states {

		if states[i].Tick > tick {
			break
		}

		last = &states[i]
	}

	return last
}
