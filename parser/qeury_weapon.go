package main

import "strings"

// GetPlayerWeaponAtTick 返回玩家当前武器。
// 如果找不到玩家状态则返回空字串。
func GetPlayerWeaponAtTick(
	ctx *AnalyzerContext,
	playerID uint8,
	tick int,
) string {

	state := GetPlayerStateAtTick(ctx, playerID, tick)

	if state == nil {
		return ""
	}

	return state.Weapon
}

// IsScopedAtTick 判断玩家当前是否开镜。
func IsScopedAtTick(
	ctx *AnalyzerContext,
	playerID uint8,
	tick int,
) bool {

	state := GetPlayerStateAtTick(ctx, playerID, tick)

	if state == nil {
		return false
	}

	return state.IsScoped
}

// IsReloadingAtTick 判断玩家当前是否正在换弹。
func IsReloadingAtTick(
	ctx *AnalyzerContext,
	playerID uint8,
	tick int,
) bool {

	state := GetPlayerStateAtTick(ctx, playerID, tick)

	if state == nil {
		return false
	}

	return state.IsReloading
}

// HasWeaponAtTick 判断玩家当前是否持有指定武器。
func HasWeaponAtTick(
	ctx *AnalyzerContext,
	playerID uint8,
	tick int,
	weapon string,
) bool {

	return GetPlayerWeaponAtTick(
		ctx,
		playerID,
		tick,
	) == weapon
}

// GetAmmoInMagazine 返回当前弹匣子弹数。
func GetAmmoInMagazine(
	ctx *AnalyzerContext,
	playerID uint8,
	tick int,
) int {

	state := GetPlayerStateAtTick(ctx, playerID, tick)
	if state == nil {
		return 0
	}

	return state.AmmoInMagazine
}

// GetAmmoReserve 返回备用子弹数。
func GetAmmoReserve(
	ctx *AnalyzerContext,
	playerID uint8,
	tick int,
) int {

	state := GetPlayerStateAtTick(ctx, playerID, tick)
	if state == nil {
		return 0
	}

	return state.AmmoReserve
}

// GetZoomLevel 返回开镜等级。
func GetZoomLevel(
	ctx *AnalyzerContext,
	playerID uint8,
	tick int,
) int {

	state := GetPlayerStateAtTick(ctx, playerID, tick)
	if state == nil {
		return 0
	}

	return state.ZoomLevel
}

// HasAmmo 判断当前弹匣是否还有子弹。
func HasAmmo(
	ctx *AnalyzerContext,
	playerID uint8,
	tick int,
) bool {

	return GetAmmoInMagazine(
		ctx,
		playerID,
		tick,
	) > 0
}

func IsPistol(
	ctx *AnalyzerContext,
	playerID uint8,
	tick int,
) bool {

	w := strings.ToLower(
		GetPlayerWeaponAtTick(ctx, playerID, tick),
	)

	switch w {
	case "glock",
		"usp-s",
		"usp",
		"p2000",
		"p250",
		"five-seven",
		"cz75-auto",
		"dual berettas",
		"tec-9",
		"desert eagle",
		"revolver":
		return true
	}

	return false
}

func IsRifle(
	ctx *AnalyzerContext,
	playerID uint8,
	tick int,
) bool {

	w := strings.ToLower(
		GetPlayerWeaponAtTick(ctx, playerID, tick),
	)

	switch w {
	case "ak-47",
		"m4a1",
		"m4a1-s",
		"famas",
		"galil ar",
		"aug",
		"sg 553":
		return true
	}

	return false
}

func IsSMG(
	ctx *AnalyzerContext,
	playerID uint8,
	tick int,
) bool {

	w := strings.ToLower(
		GetPlayerWeaponAtTick(ctx, playerID, tick),
	)

	switch w {
	case "mac-10",
		"mp9",
		"mp7",
		"mp5-sd",
		"ump-45",
		"bizon",
		"p90":
		return true
	}

	return false
}

func IsSniper(
	ctx *AnalyzerContext,
	playerID uint8,
	tick int,
) bool {

	w := strings.ToLower(
		GetPlayerWeaponAtTick(ctx, playerID, tick),
	)

	switch w {
	case "awp",
		"ssg 08",
		"g3sg1",
		"scar-20":
		return true
	}

	return false
}

func IsAWP(
	ctx *AnalyzerContext,
	playerID uint8,
	tick int,
) bool {

	return strings.EqualFold(
		GetPlayerWeaponAtTick(ctx, playerID, tick),
		"awp",
	)
}

func IsAK47(
	ctx *AnalyzerContext,
	playerID uint8,
	tick int,
) bool {

	return strings.EqualFold(
		GetPlayerWeaponAtTick(ctx, playerID, tick),
		"ak-47",
	)
}

func IsDeagle(
	ctx *AnalyzerContext,
	playerID uint8,
	tick int,
) bool {

	w := GetPlayerWeaponAtTick(ctx, playerID, tick)

	return strings.EqualFold(w, "desert eagle") ||
		strings.EqualFold(w, "deagle")
}
