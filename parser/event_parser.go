package main

import (
	"fmt"
	"os"

	dem "github.com/markus-wa/demoinfocs-golang/v5/pkg/demoinfocs"
	"github.com/markus-wa/demoinfocs-golang/v5/pkg/demoinfocs/events"
)

func EventParser(path string) error {
	f, err := os.Open(path)

	defer f.Close()

	parser := dem.NewParser(f)
	defer parser.Close()

	var players []Player
	var kills []Kill
	stats := make(map[uint64]*PlayerStats)
	var rounds []Round
	roundNumber := 0
	var match Match
	var teams []TeamGroup
	var damages []Damage
	var weaponFires []WeaponFire
	var shots []Shot
	var positions []PlayerState

	parser.RegisterEventHandler(func(events.RoundFreezetimeEnd) {
		if len(teams) > 0 {
			return
		}

		teams = BuildStableTeamsFromParticipants(
			parser.GameState().Participants().Playing(),
		)
	})

	parser.RegisterEventHandler(func(e events.PlayerHurt) {

		if e.Attacker == nil || e.Player == nil {
			return
		}

		if e.Attacker.SteamID64 == 0 || e.Player.SteamID64 == 0 {
			return
		}

		weapon := ""

		if e.Weapon != nil {
			weapon = e.Weapon.String()
		}

		println(
			e.Attacker.Name,
			"->",
			e.Player.Name,
			"Damage:",
			e.HealthDamage,
		)
		damages = append(damages, Damage{

			Tick:  parser.GameState().IngameTick(),
			Round: parser.GameState().TotalRoundsPlayed(),

			Attacker:          e.Attacker.Name,
			AttackerSteamID64: e.Attacker.SteamID64,

			Victim:          e.Player.Name,
			VictimSteamID64: e.Player.SteamID64,

			Weapon: weapon,

			Damage: e.HealthDamage,

			HealthDamage: e.HealthDamage,
			ArmorDamage:  e.ArmorDamage,

			HealthDamageTaken: e.HealthDamageTaken,
			ArmorDamageTaken:  e.ArmorDamageTaken,

			HealthRemaining: e.Player.Health(),
			ArmorRemaining:  e.Player.Armor(),

			HitGroup: hitGroupName(e.HitGroup),
		})

		for i := len(shots) - 1; i >= 0; i-- {

			if shots[i].SteamID64 != e.Attacker.SteamID64 {
				continue
			}

			if shots[i].Hit {
				continue
			}

			shots[i].Hit = true

			shots[i].HitTick = parser.GameState().IngameTick()

			shots[i].TimeToDamage =
				shots[i].HitTick -
					shots[i].FireTick

			shots[i].Victim = e.Player.Name

			shots[i].VictimSteamID64 = e.Player.SteamID64

			shots[i].Damage = e.HealthDamageTaken

			shots[i].HitGroup = hitGroupName(e.HitGroup)

			break
		}
	})

	parser.RegisterEventHandler(func(e events.WeaponFire) {

		if e.Shooter == nil {
			return
		}

		if e.Shooter.SteamID64 == 0 {
			return
		}

		weapon := ""

		if e.Weapon != nil {
			weapon = e.Weapon.String()
		}

		weaponFires = append(weaponFires, WeaponFire{

			Tick:  parser.GameState().IngameTick(),
			Round: roundNumber,

			Player:    e.Shooter.Name,
			SteamID64: e.Shooter.SteamID64,

			Weapon: weapon,
		})
		shots = append(shots, Shot{

			ShotIndex: len(shots) + 1,

			FireTick: parser.GameState().IngameTick(),

			Round: roundNumber,

			Player: e.Shooter.Name,

			SteamID64: e.Shooter.SteamID64,

			Weapon: weapon,
		})
	})
	parser.RegisterEventHandler(func(e events.Kill) {

		if e.Killer == nil || e.Victim == nil {
			return
		}

		// 过滤 DemoRecorder / 世界实体
		if e.Killer.SteamID64 == 0 || e.Victim.SteamID64 == 0 {
			return
		}

		if _, ok := stats[e.Killer.SteamID64]; !ok {
			stats[e.Killer.SteamID64] = &PlayerStats{
				SteamID64: e.Killer.SteamID64,
			}
		}

		stats[e.Killer.SteamID64].Name = e.Killer.Name

		if _, ok := stats[e.Victim.SteamID64]; !ok {
			stats[e.Victim.SteamID64] = &PlayerStats{
				SteamID64: e.Victim.SteamID64,
			}
		}

		stats[e.Victim.SteamID64].Name = e.Victim.Name

		// 先过滤自杀（以后可以单独统计）
		if e.Killer.SteamID64 == e.Victim.SteamID64 {
			return
		}

		println("EVENT:", e.Killer.Name, "->", e.Victim.Name)

		assisterName := ""
		assisterSteamID := uint64(0)

		if e.Assister != nil {
			assisterName = e.Assister.Name
			assisterSteamID = e.Assister.SteamID64
		}

		fmt.Fprintln(
			os.Stderr,
			"Kill Debug",
			"Tick:", parser.GameState().IngameTick(),
			"RoundNumber:", roundNumber,
			"TotalRoundsPlayed:", parser.GameState().TotalRoundsPlayed(),
			"Weapon:", e.Weapon.String(),
			"Killer:", e.Killer.Name,
			"Victim:", e.Victim.Name,
		)

		kills = append(kills, Kill{
			Tick:  parser.GameState().IngameTick(),
			Round: roundNumber,

			Killer:          e.Killer.Name,
			KillerSteamID64: e.Killer.SteamID64,
			KillerTeam:      teamName(e.Killer.Team),

			Victim:          e.Victim.Name,
			VictimSteamID64: e.Victim.SteamID64,
			VictimTeam:      teamName(e.Victim.Team),

			Assister:          assisterName,
			AssisterSteamID64: assisterSteamID,

			Weapon: e.Weapon.String(),

			Headshot: e.IsHeadshot,

			Wallbang: e.IsWallBang(),

			ThroughSmoke: e.ThroughSmoke,

			Blind: e.AttackerBlind,

			NoScope: e.NoScope,

			FlashAssist: e.AssistedFlash,

			Penetration: e.PenetratedObjects,

			Distance: e.Distance,
		})

		for i := len(shots) - 1; i >= 0; i-- {

			if shots[i].SteamID64 != e.Killer.SteamID64 {
				continue
			}

			if shots[i].Kill {
				continue
			}

			if !shots[i].Hit {
				continue
			}

			if shots[i].VictimSteamID64 != e.Victim.SteamID64 {
				continue
			}

			shots[i].Kill = true

			shots[i].Headshot = e.IsHeadshot

			shots[i].Wallbang = e.IsWallBang()

			shots[i].ThroughSmoke = e.ThroughSmoke

			shots[i].Blind = e.AttackerBlind

			shots[i].NoScope = e.NoScope

			shots[i].FlashAssist = e.AssistedFlash

			shots[i].Distance = e.Distance

			shots[i].Penetration = e.PenetratedObjects

			shots[i].TimeToKill =
				parser.GameState().IngameTick() -
					shots[i].FireTick

			break
		}

		for i := len(shots) - 1; i >= 0; i-- {

			if shots[i].SteamID64 != e.Killer.SteamID64 {
				continue
			}

			if shots[i].VictimSteamID64 != e.Victim.SteamID64 {
				continue
			}

			shots[i].Kill = true

			shots[i].Headshot = e.IsHeadshot

			shots[i].Wallbang = e.IsWallBang()

			shots[i].ThroughSmoke = e.ThroughSmoke

			shots[i].Blind = e.AttackerBlind

			shots[i].NoScope = e.NoScope

			shots[i].FlashAssist = e.AssistedFlash

			shots[i].Distance = e.Distance

			shots[i].Penetration = e.PenetratedObjects

			break
		}

		if s, ok := stats[e.Killer.SteamID64]; ok {
			s.Kills++

			if e.IsHeadshot {
				s.Headshots++
			}
		}

		if s, ok := stats[e.Victim.SteamID64]; ok {
			s.Deaths++
		}
	})
	parser.RegisterEventHandler(func(e events.RoundEnd) {

		roundNumber++

		state := parser.GameState()

		rounds = append(rounds, Round{
			Number:  roundNumber,
			Winner:  teamName(e.Winner),
			ScoreCT: state.TeamCounterTerrorists().Score(),
			ScoreT:  state.TeamTerrorists().Score(),
		})
	})

	err = parser.ParseToEnd()

	if err != nil {
		output(Result{
			Success: false,
			Message: err.Error(),
		})
		return err
	}

	state := parser.GameState()

	match = Match{
		Map:      "",
		TickRate: int(parser.TickRate()),

		ScoreCT: state.TeamCounterTerrorists().Score(),
		ScoreT:  state.TeamTerrorists().Score(),
	}
	if match.ScoreCT > match.ScoreT {
		match.Winner = "CT"
	} else {
		match.Winner = "T"
	}

	state = parser.GameState()

	for _, p := range state.Participants().All() {

		if p.SteamID64 == 0 {
			continue
		}

		players = append(players, Player{
			Name:      p.Name,
			SteamID64: p.SteamID64,
			PlayerID:  GetPlayerID(p.SteamID64),
			Team:      teamName(p.Team),
			IsBot:     p.IsBot,
		})
	}

	if len(teams) == 0 {
		teams = BuildStableTeamsFromPlayers(players)
	}

	ApplyTeamGroupsToPlayers(players, teams)

	for i := range shots {

		if !shots[i].Hit {

			shots[i].Miss = true
		}
	}
	if err != nil {
		output(Result{
			Success: false,
			Message: err.Error(),
		})
		return err
	}

	fmt.Fprintln(os.Stderr, "Kill Count:", len(kills))

	var statsList []PlayerStats

	for _, s := range stats {
		statsList = append(statsList, *s)
	}

	result := Result{
		Success: true,

		Match: match,

		Players: players,

		Teams: teams,

		Kills: kills,

		Damages: damages,

		WeaponFires: weaponFires,

		Shots: shots,

		PlayerStates: positions,

		Rounds: rounds,

		Stats: statsList,
	}

	positions, err = CollectPlayerStates(path)

	if err != nil {
		return err
	}

	result.PlayerStates = positions

	ctx := BuildContext(&result)

	ValidateResult(&result)

	Analyze(ctx, &result)

	GenerateInsights(&result)

	for _, s := range result.Stats {

		println(
			s.Name,
			"Trade:",
			s.TradeKills,
			"/",
			s.TradeDeaths,
		)
	}

	ValidateResult(&result)

	result.Facts = BuildFacts(&result)

	output(result)

	return nil
}
