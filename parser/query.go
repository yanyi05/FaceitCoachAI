package main

func GetPlayerStates(
	ctx *AnalyzerContext,
	playerID uint8,
) []PlayerState {

	return ctx.PlayerStateByPlayer[playerID]
}
