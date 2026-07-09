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

// GetPlayerPositionAtTick 返回玩家坐标。
func GetPlayerPositionAtTick(
	ctx *AnalyzerContext,
	playerID uint8,
	tick int,
) (int16, int16, int16) {

	state := GetPlayerStateAtTick(ctx, playerID, tick)
	if state == nil {
		return 0, 0, 0
	}

	return state.X, state.Y, state.Z
}

// GetPlayerHealthAtTick 返回 HP。
func GetPlayerHealthAtTick(
	ctx *AnalyzerContext,
	playerID uint8,
	tick int,
) uint8 {

	state := GetPlayerStateAtTick(ctx, playerID, tick)
	if state == nil {
		return 0
	}

	return state.HP
}

// GetPlayerArmorAtTick 返回 Armor。
func GetPlayerArmorAtTick(
	ctx *AnalyzerContext,
	playerID uint8,
	tick int,
) uint8 {

	state := GetPlayerStateAtTick(ctx, playerID, tick)
	if state == nil {
		return 0
	}

	return state.Armor
}

// GetPlayerMoneyAtTick 返回金钱。
func GetPlayerMoneyAtTick(
	ctx *AnalyzerContext,
	playerID uint8,
	tick int,
) int {

	state := GetPlayerStateAtTick(ctx, playerID, tick)
	if state == nil {
		return 0
	}

	return state.Money
}

// IsAliveAtTick 判断玩家是否存活。
func IsAliveAtTick(
	ctx *AnalyzerContext,
	playerID uint8,
	tick int,
) bool {

	state := GetPlayerStateAtTick(ctx, playerID, tick)
	if state == nil {
		return false
	}

	return state.Alive
}

// GetPlayerTeamAtTick 返回玩家队伍。
func GetPlayerTeamAtTick(
	ctx *AnalyzerContext,
	playerID uint8,
	tick int,
) string {

	state := GetPlayerStateAtTick(ctx, playerID, tick)
	if state == nil {
		return ""
	}

	return state.Team
}

// GetPlayerViewAngleAtTick 返回玩家视角。
func GetPlayerViewAngleAtTick(
	ctx *AnalyzerContext,
	playerID uint8,
	tick int,
) (float32, float32) {

	state := GetPlayerStateAtTick(ctx, playerID, tick)
	if state == nil {
		return 0, 0
	}

	return state.ViewYaw, state.ViewPitch
}

// GetPlayerLastPlaceAtTick 返回地图区域名称。
func GetPlayerLastPlaceAtTick(
	ctx *AnalyzerContext,
	playerID uint8,
	tick int,
) string {

	state := GetPlayerStateAtTick(ctx, playerID, tick)
	if state == nil {
		return ""
	}

	return state.LastPlace
}
