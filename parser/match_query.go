package main

func (m *Match) GetPlayer(steamID uint64) *Player {

	for i := range m.Players {

		if m.Players[i].SteamID64 == steamID {
			return &m.Players[i]
		}
	}

	return nil
}

func (m *Match) GetRound(round int) *Round {

	for i := range m.Rounds {

		if m.Rounds[i].Number == round {
			return &m.Rounds[i]
		}
	}

	return nil
}

func (m *Match) GetPlayerByName(name string) *Player {

	for i := range m.Players {

		if m.Players[i].Name == name {
			return &m.Players[i]
		}
	}

	return nil
}

func (m *Match) GetKillsByRound(round int) []Kill {

	var result []Kill

	for _, kill := range m.Kills {

		if kill.Round == round {
			result = append(result, kill)
		}
	}

	return result
}

func (m *Match) GetRecentDamage(
	steamID uint64,
	beforeTick int,
	maxTickDiff int,
) []Damage {

	var result []Damage

	for _, damage := range m.Damages {

		if damage.AttackerSteamID64 != steamID {
			continue
		}

		if damage.Tick > beforeTick {
			continue
		}

		if beforeTick-damage.Tick > maxTickDiff {
			continue
		}

		result = append(result, damage)
	}

	return result
}

func (m *Match) GetRecentShots(
	steamID uint64,
	beforeTick int,
	maxTickDiff int,
) []Shot {

	var result []Shot

	for _, shot := range m.Shots {

		if shot.SteamID64 != steamID {
			continue
		}

		if shot.FireTick > beforeTick {
			continue
		}

		if beforeTick-shot.FireTick > maxTickDiff {
			continue
		}

		result = append(result, shot)
	}

	return result
}

func (m *Match) GetDamageHistory(
	steamID uint64,
) []Damage {

	var result []Damage

	for _, damage := range m.Damages {

		if damage.AttackerSteamID64 != steamID {
			continue
		}

		result = append(result, damage)
	}

	return result
}

func (m *Match) GetShotHistory(
	steamID uint64,
) []Shot {

	var result []Shot

	for _, shot := range m.Shots {

		if shot.SteamID64 != steamID {
			continue
		}

		result = append(result, shot)
	}

	return result
}

func (m *Match) GetLatestDeathSnapshot() *DeathSnapshot {

	if len(m.DeathSnapshots) == 0 {
		return nil
	}

	return &m.DeathSnapshots[len(m.DeathSnapshots)-1]
}
