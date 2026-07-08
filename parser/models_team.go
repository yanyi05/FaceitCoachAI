package main

import (
	"sort"

	"github.com/markus-wa/demoinfocs-golang/v5/pkg/demoinfocs/common"
)

type TeamMember struct {
	Name      string `json:"name"`
	SteamID64 uint64 `json:"steamId"`
	PlayerID  uint8  `json:"playerId"`
}

type TeamGroup struct {
	ID        string       `json:"id"`
	Name      string       `json:"name"`
	StartSide string       `json:"startSide"`
	Players   []TeamMember `json:"players"`
}

func BuildStableTeamsFromParticipants(players []*common.Player) []TeamGroup {
	ct := make([]TeamMember, 0, 5)
	t := make([]TeamMember, 0, 5)

	for _, player := range players {
		if player == nil || player.SteamID64 == 0 {
			continue
		}

		member := TeamMember{
			Name:      player.Name,
			SteamID64: player.SteamID64,
			PlayerID:  GetPlayerID(player.SteamID64),
		}

		switch teamName(player.Team) {
		case "CT":
			ct = append(ct, member)
		case "T":
			t = append(t, member)
		}
	}

	if len(ct) != 5 || len(t) != 5 {
		return nil
	}

	sortTeamMembers(ct)
	sortTeamMembers(t)

	return []TeamGroup{
		{
			ID:        "A",
			Name:      "Team A",
			StartSide: "CT",
			Players:   ct,
		},
		{
			ID:        "B",
			Name:      "Team B",
			StartSide: "T",
			Players:   t,
		},
	}
}

func BuildStableTeamsFromPlayers(players []Player) []TeamGroup {
	ct := make([]TeamMember, 0, 5)
	t := make([]TeamMember, 0, 5)

	for _, player := range players {
		member := TeamMember{
			Name:      player.Name,
			SteamID64: player.SteamID64,
			PlayerID:  player.PlayerID,
		}

		switch player.Team {
		case "CT":
			ct = append(ct, member)
		case "T":
			t = append(t, member)
		}
	}

	if len(ct) != 5 || len(t) != 5 {
		return nil
	}

	sortTeamMembers(ct)
	sortTeamMembers(t)

	return []TeamGroup{
		{
			ID:        "A",
			Name:      "Team A",
			StartSide: "CT",
			Players:   ct,
		},
		{
			ID:        "B",
			Name:      "Team B",
			StartSide: "T",
			Players:   t,
		},
	}
}

func ApplyTeamGroupsToPlayers(players []Player, teams []TeamGroup) {
	groupBySteam := make(map[uint64]string)

	for _, team := range teams {
		for _, player := range team.Players {
			groupBySteam[player.SteamID64] = team.ID
		}
	}

	for i := range players {
		players[i].TeamGroup = groupBySteam[players[i].SteamID64]
	}
}

func sortTeamMembers(players []TeamMember) {
	sort.Slice(players, func(i, j int) bool {
		if players[i].Name == players[j].Name {
			return players[i].SteamID64 < players[j].SteamID64
		}

		return players[i].Name < players[j].Name
	})
}
