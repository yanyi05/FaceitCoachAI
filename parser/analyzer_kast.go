package main

func calcKAST(result *Result, stats map[uint64]*PlayerStats) {

	type RoundState struct {
		Kill    bool
		Assist  bool
		Survive bool
	}

	roundStats := make(map[uint64]map[int]*RoundState)

	// 初始化
	for _, p := range result.Players {

		roundStats[p.SteamID64] = make(map[int]*RoundState)

		for _, r := range result.Rounds {
			roundStats[p.SteamID64][r.Number] = &RoundState{
				Survive: true,
			}
		}
	}

	// Kill
	for _, k := range result.Kills {

		roundStats[k.KillerSteamID64][k.Round].Kill = true

		roundStats[k.VictimSteamID64][k.Round].Survive = false

		if k.AssisterSteamID64 != 0 {
			roundStats[k.AssisterSteamID64][k.Round].Assist = true
		}
	}

	// 统计
	for steamID, rounds := range roundStats {

		success := 0

		total := len(rounds)

		for _, rs := range rounds {

			if rs.Kill || rs.Assist || rs.Survive {
				success++
			}
		}

		if s, ok := stats[steamID]; ok && total > 0 {
			s.KAST = float64(success) / float64(total) * 100
		}
	}
}
