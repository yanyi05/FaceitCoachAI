package main

// GetAllPlayerStatesAtTick 返回指定 Tick 所有玩家状态。
func GetAllPlayerStatesAtTick(
	ctx *AnalyzerContext,
	tick int,
) []PlayerState {

	var result []PlayerState

	for playerID := range ctx.PlayerByID {

		state := GetPlayerStateAtTick(ctx, playerID, tick)
		if state == nil {
			continue
		}

		result = append(result, *state)
	}

	return result
}

// GetAlivePlayersAtTick 返回指定 Tick 所有存活玩家。
func GetAlivePlayersAtTick(
	ctx *AnalyzerContext,
	tick int,
) []PlayerState {

	var result []PlayerState

	states := GetAllPlayerStatesAtTick(ctx, tick)

	for _, state := range states {
		if !state.Alive {
			continue
		}

		result = append(result, state)
	}

	return result
}
