package main

func BuildOpeningStats(result *Result) {

	// 重置数据
	for i := range result.Stats {
		result.Stats[i].OpeningKills = 0
		result.Stats[i].OpeningDeaths = 0
	}

	// 每一回合第一杀
	firstKill := make(map[int]bool)

	for _, k := range result.Kills {

		if firstKill[k.Round] {
			continue
		}

		firstKill[k.Round] = true

		for i := range result.Stats {

			if result.Stats[i].SteamID64 == k.KillerSteamID64 {
				result.Stats[i].OpeningKills++
			}

			if result.Stats[i].SteamID64 == k.VictimSteamID64 {
				result.Stats[i].OpeningDeaths++
			}
		}
	}
}
