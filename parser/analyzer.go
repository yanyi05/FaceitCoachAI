package main

func Analyze(
	ctx *AnalyzerContext,
	result *Result,
) {

	// ==========================
	// Build Timeline Data
	// ==========================

	BuildBurst(result)
	BuildSpray(result)

	// ==========================
	// Basic Statistics
	// ==========================

	stats := buildStatsIndex(result)

	calcAccuracy(result, stats)
	calcADR(result, stats)
	calcHeadshot(result, stats)
	calcEntry(result, stats)
	calcTTD(result, stats)
	// calcTrade(result, stats)

	// ==========================
	// AI Analysis
	// ==========================

	support := AnalyzeSupport(ctx, result)
	result.AI.Support = support

	AnalyzeTrade(result, support)

	AnalyzeRotation(ctx, result)

	//AnalyzePlayerState(result)
}

func buildStatsIndex(result *Result) map[uint64]*PlayerStats {

	stats := make(map[uint64]*PlayerStats)

	for i := range result.Stats {

		stats[result.Stats[i].SteamID64] = &result.Stats[i]
	}

	return stats
}

func calcAccuracy(
	result *Result,
	stats map[uint64]*PlayerStats,
) {

	for _, fire := range result.WeaponFires {

		if s, ok := stats[fire.SteamID64]; ok {

			s.ShotsFired++
		}
	}

	for _, dmg := range result.Damages {

		if s, ok := stats[dmg.AttackerSteamID64]; ok {

			s.ShotsHit++
			s.Damage += dmg.HealthDamageTaken
		}

		if s, ok := stats[dmg.VictimSteamID64]; ok {

			s.DamageTaken += dmg.HealthDamageTaken
		}
	}

	for _, s := range stats {

		if s.ShotsFired > 0 {

			s.Accuracy =
				float64(s.ShotsHit) /
					float64(s.ShotsFired) * 100
		}
	}
}

func calcADR(
	result *Result,
	stats map[uint64]*PlayerStats,
) {

	rounds := len(result.Rounds)

	println("Rounds:", rounds)

	for _, s := range stats {

		println(
			s.Name,
			"Damage:",
			s.Damage,
			"Kills:",
			s.Kills,
		)
	}

	if rounds == 0 {
		return
	}

	for _, s := range stats {

		s.ADR =
			float64(s.Damage) /
				float64(rounds)
	}
}

func calcHeadshot(
	result *Result,
	stats map[uint64]*PlayerStats,
) {

	for _, s := range stats {

		if s.Kills > 0 {

			s.HeadshotPercent =
				float64(s.Headshots) /
					float64(s.Kills) * 100
		}
	}
}

func calcEntry(
	result *Result,
	stats map[uint64]*PlayerStats,
) {

	firstKill := make(map[int]bool)

	for _, kill := range result.Kills {

		if firstKill[kill.Round] {
			continue
		}

		firstKill[kill.Round] = true

		if s, ok := stats[kill.KillerSteamID64]; ok {
			s.EntryKills++
		}

		if s, ok := stats[kill.VictimSteamID64]; ok {
			s.EntryDeaths++
		}
	}
}

func calcTTD(
	result *Result,
	stats map[uint64]*PlayerStats,
) {

	for i := range result.WeaponFires {

		fire := &result.WeaponFires[i]

		if fire.Matched {
			continue
		}

		for j := range result.Damages {

			dmg := &result.Damages[j]

			if dmg.Matched {
				continue
			}

			if dmg.AttackerSteamID64 != fire.SteamID64 {
				continue
			}

			if dmg.Round != fire.Round {
				continue
			}

			if dmg.Tick < fire.Tick {
				continue
			}

			if dmg.Weapon != fire.Weapon {
				continue
			}

			delta := dmg.Tick - fire.Tick

			if delta > 10 {
				break
			}

			fire.Matched = true
			dmg.Matched = true

			if s, ok := stats[fire.SteamID64]; ok {

				s.TotalTTDTicks += delta
				s.TTDCount++
			}

			break
		}
	}

	for _, s := range stats {

		if s.TTDCount > 0 {

			s.AverageTTD =
				float64(s.TotalTTDTicks) /
					float64(s.TTDCount)
		}
	}
}
