package main

func AnalyzeEntry(result *Result) {

	// 清空旧数据
	for i := range result.Stats {

		result.Stats[i].EntryKills = 0
		result.Stats[i].EntryDeaths = 0

		result.Stats[i].EntryAttempts = 0
		result.Stats[i].EntrySuccess = 0
	}

	// 每回合第一杀
	firstKillRound := make(map[int]bool)

	for _, kill := range result.Kills {

		if firstKillRound[kill.Round] {
			continue
		}

		firstKillRound[kill.Round] = true

		for i := range result.Stats {

			if result.Stats[i].SteamID64 == kill.KillerSteamID64 {

				result.Stats[i].EntryKills++
				result.Stats[i].EntryAttempts++
			}

			if result.Stats[i].SteamID64 == kill.VictimSteamID64 {

				result.Stats[i].EntryDeaths++
				result.Stats[i].EntryAttempts++
			}
		}
	}

	for i := range result.Stats {

		if result.Stats[i].EntryAttempts == 0 {
			continue
		}

		result.Stats[i].EntrySuccess =
			float64(result.Stats[i].EntryKills) /
				float64(result.Stats[i].EntryAttempts) * 100
	}
}
