package main

func findPlayer(result *Result, steam uint64) *Player {

	for i := range result.Players {

		if result.Players[i].SteamID64 == steam {

			return &result.Players[i]
		}
	}

	return nil
}
