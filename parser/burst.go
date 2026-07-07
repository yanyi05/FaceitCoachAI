package main

const BurstGap = 8

func BuildBurst(result *Result) {

	if len(result.Shots) == 0 {
		return
	}

	// 每个玩家独立计算 Burst
	lastShot := make(map[uint64]*Shot)

	currentBurst := make(map[uint64]int)

	for i := range result.Shots {

		shot := &result.Shots[i]

		last, ok := lastShot[shot.SteamID64]

		// 第一枪
		if !ok {

			currentBurst[shot.SteamID64] = 1

			shot.BurstIndex = 1

			shot.BurstShot = 1

			lastShot[shot.SteamID64] = shot

			continue
		}

		// 是否属于同一个 Burst
		if shot.FireTick-last.FireTick <= BurstGap {

			shot.BurstIndex = currentBurst[shot.SteamID64]

			shot.BurstShot = last.BurstShot + 1

		} else {

			currentBurst[shot.SteamID64]++

			shot.BurstIndex = currentBurst[shot.SteamID64]

			shot.BurstShot = 1
		}

		lastShot[shot.SteamID64] = shot
	}
}
