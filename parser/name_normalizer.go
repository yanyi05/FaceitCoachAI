package main

func BuildNameMap(players []Player) map[uint64]string {
	names := make(map[uint64]string)

	for _, p := range players {
		names[p.SteamID64] = p.Name
	}

	return names
}

func NormalizeResultNames(result *Result) {
	names := BuildNameMap(result.Players)

	for i := range result.Kills {

		if name, ok := names[result.Kills[i].KillerSteamID64]; ok {
			result.Kills[i].Killer = name
		}

		if name, ok := names[result.Kills[i].VictimSteamID64]; ok {
			result.Kills[i].Victim = name
		}

		if result.Kills[i].AssisterSteamID64 != 0 {
			if name, ok := names[result.Kills[i].AssisterSteamID64]; ok {
				result.Kills[i].Assister = name
			}
		}
	}
}
