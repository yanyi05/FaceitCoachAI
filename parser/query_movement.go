package main

import "math"

// GetVelocityAtTick 返回玩家速度。
func GetVelocityAtTick(
	ctx *AnalyzerContext,
	playerID uint8,
	tick int,
) float32 {

	state := GetPlayerStateAtTick(ctx, playerID, tick)

	if state == nil {
		return 0
	}

	return state.Velocity
}

// IsWalkingAtTick 判断玩家是否走路。
func IsWalkingAtTick(
	ctx *AnalyzerContext,
	playerID uint8,
	tick int,
) bool {

	state := GetPlayerStateAtTick(ctx, playerID, tick)

	if state == nil {
		return false
	}

	return state.IsWalking
}

// IsStandingAtTick 判断玩家是否站立。
func IsStandingAtTick(
	ctx *AnalyzerContext,
	playerID uint8,
	tick int,
) bool {

	state := GetPlayerStateAtTick(ctx, playerID, tick)

	if state == nil {
		return false
	}

	return state.IsStanding
}

// IsAirborneAtTick 判断玩家是否在空中。
func IsAirborneAtTick(
	ctx *AnalyzerContext,
	playerID uint8,
	tick int,
) bool {

	state := GetPlayerStateAtTick(ctx, playerID, tick)

	if state == nil {
		return false
	}

	return state.IsAirborne
}

// IsMovingAtTick 判断玩家是否正在移动。
func IsMovingAtTick(
	ctx *AnalyzerContext,
	playerID uint8,
	tick int,
) bool {

	return GetVelocityAtTick(
		ctx,
		playerID,
		tick,
	) > 5
}

// IsDuckingAtTick 判断玩家是否蹲下。
func IsDuckingAtTick(
	ctx *AnalyzerContext,
	playerID uint8,
	tick int,
) bool {

	state := GetPlayerStateAtTick(ctx, playerID, tick)

	if state == nil {
		return false
	}

	return state.IsDucking
}

// IsBlindedAtTick 判断玩家是否被闪光影响。
func IsBlindedAtTick(
	ctx *AnalyzerContext,
	playerID uint8,
	tick int,
) bool {

	state := GetPlayerStateAtTick(ctx, playerID, tick)

	if state == nil {
		return false
	}

	return state.IsBlinded
}

// GetFlashDurationAtTick 返回剩余致盲时间。
func GetFlashDurationAtTick(
	ctx *AnalyzerContext,
	playerID uint8,
	tick int,
) float32 {

	state := GetPlayerStateAtTick(ctx, playerID, tick)

	if state == nil {
		return 0
	}

	return state.FlashDuration
}

// GetVelocityVectorAtTick 返回速度向量。
func GetVelocityVectorAtTick(
	ctx *AnalyzerContext,
	playerID uint8,
	tick int,
) (float32, float32, float32) {

	state := GetPlayerStateAtTick(ctx, playerID, tick)

	if state == nil {
		return 0, 0, 0
	}

	return state.VelocityX, state.VelocityY, state.VelocityZ
}

// GetSpeed2DAtTick 返回水平速度。
func GetSpeed2DAtTick(
	ctx *AnalyzerContext,
	playerID uint8,
	tick int,
) float32 {

	vx, vy, _ := GetVelocityVectorAtTick(
		ctx,
		playerID,
		tick,
	)

	return float32(math.Sqrt(float64(vx*vx + vy*vy)))
}
