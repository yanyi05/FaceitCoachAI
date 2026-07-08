package main

import (
	"fmt"
	"math"
	"os"

	dem "github.com/markus-wa/demoinfocs-golang/v5/pkg/demoinfocs"
	"github.com/markus-wa/demoinfocs-golang/v5/pkg/demoinfocs/events"
	"github.com/markus-wa/demoinfocs-golang/v5/pkg/demoinfocs/msg"
)

type DebugMode int

const (
	DebugNone DebugMode = iota
	DebugTrade
	DebugAPI
	DebugSupport
	DebugCrossfire
)

func EventParser(path string, debug DebugMode) error {
	f, err := os.Open(path)

	defer f.Close()

	parser := dem.NewParser(f)
	defer parser.Close()

	stats := make(map[uint64]*PlayerStats)

	roundNumber := 0
	var match Match
	var mapName string
	var teams []TeamGroup
	var weaponFires []WeaponFire
	var positionCache *PositionCache
	var deathSnapshots []DeathSnapshot
	var matchStarted bool

	parser.RegisterNetMessageHandler(func(m *msg.CSVCMsg_ServerInfo) {
		mapName = m.GetMapName()
	})

	parser.RegisterEventHandler(func(events.MatchStart) {
		matchStarted = true
	})

	parser.RegisterEventHandler(func(e events.PlayerHurt) {

		if !matchStarted {
			return
		}

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
		match.Damages = append(match.Damages, Damage{

			Tick:  parser.GameState().IngameTick(),
			Round: parser.GameState().TotalRoundsPlayed() + 1,

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

		for i := len(match.Shots) - 1; i >= 0; i-- {

			if match.Shots[i].SteamID64 != e.Attacker.SteamID64 {
				continue
			}

			if match.Shots[i].Hit {
				continue
			}

			match.Shots[i].Hit = true

			match.Shots[i].HitTick = parser.GameState().IngameTick()

			match.Shots[i].TimeToDamage =
				match.Shots[i].HitTick -
					match.Shots[i].FireTick

			match.Shots[i].Victim = e.Player.Name

			match.Shots[i].VictimSteamID64 = e.Player.SteamID64

			match.Shots[i].Damage = e.HealthDamageTaken

			match.Shots[i].HitGroup = hitGroupName(e.HitGroup)

			break
		}
	})

	parser.RegisterEventHandler(func(e events.WeaponFire) {

		if !matchStarted {
			return
		}

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
			Round: parser.GameState().TotalRoundsPlayed() + 1,

			Player:    e.Shooter.Name,
			SteamID64: e.Shooter.SteamID64,

			Weapon: weapon,
		})
		match.Shots = append(match.Shots, Shot{

			ShotIndex: len(match.Shots) + 1,

			FireTick: parser.GameState().IngameTick(),

			Round: parser.GameState().TotalRoundsPlayed() + 1,

			Player: e.Shooter.Name,

			SteamID64: e.Shooter.SteamID64,

			Weapon: weapon,
		})
	})
	parser.RegisterEventHandler(func(e events.Kill) {

		if !matchStarted {
			return
		}

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

		fmt.Println("================================")

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

		match.Kills = append(match.Kills, Kill{
			Tick:  parser.GameState().IngameTick(),
			Round: parser.GameState().TotalRoundsPlayed() + 1,

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

		victimWeapon := ""
		if e.Victim.ActiveWeapon() != nil {
			victimWeapon = e.Victim.ActiveWeapon().String()
		}

		killerWeapon := ""
		if e.Killer.ActiveWeapon() != nil {
			killerWeapon = e.Killer.ActiveWeapon().String()
		}

		snapshot := DeathSnapshot{

			Tick:  parser.GameState().IngameTick(),
			Round: roundNumber,

			// =========================
			// Victim
			// =========================

			Victim:        e.Victim.Name,
			VictimSteamID: e.Victim.SteamID64,

			VictimX: e.Victim.Position().X,
			VictimY: e.Victim.Position().Y,
			VictimZ: e.Victim.Position().Z,

			VictimViewX: e.Victim.ViewDirectionX(),
			VictimViewY: e.Victim.ViewDirectionY(),

			VictimWeapon: victimWeapon,

			VictimHP:    e.Victim.Health(),
			VictimArmor: e.Victim.Armor(),

			VictimMoney: e.Victim.Money(),

			VictimScoped:  e.Victim.IsScoped(),
			VictimDucking: e.Victim.IsDucking(),

			VictimFlashDuration: float64(e.Victim.FlashDuration),

			VictimPlace: e.Victim.LastPlaceName(),

			// =========================
			// Killer
			// =========================

			Killer:        e.Killer.Name,
			KillerSteamID: e.Killer.SteamID64,

			KillerX: e.Killer.Position().X,
			KillerY: e.Killer.Position().Y,
			KillerZ: e.Killer.Position().Z,

			KillerViewX: e.Killer.ViewDirectionX(),
			KillerViewY: e.Killer.ViewDirectionY(),

			KillerWeapon: killerWeapon,

			KillerHP:    e.Killer.Health(),
			KillerArmor: e.Killer.Armor(),

			KillerMoney: e.Killer.Money(),

			KillerScoped:  e.Killer.IsScoped(),
			KillerDucking: e.Killer.IsDucking(),

			KillerFlashDuration: float64(e.Killer.FlashDuration),

			KillerPlace: e.Killer.LastPlaceName(),
		}

		for _, p := range parser.GameState().Participants().Playing() {

			if p == nil {
				continue
			}

			if p.SteamID64 == 0 {
				continue
			}

			// 死者自己不用记录
			if p.SteamID64 == e.Victim.SteamID64 {
				continue
			}

			// 只记录死者队友
			if p.Team != e.Victim.Team {
				continue
			}

			pos := p.Position()

			dx := pos.X - e.Victim.Position().X
			dy := pos.Y - e.Victim.Position().Y
			dz := pos.Z - e.Victim.Position().Z

			distance := math.Sqrt(
				dx*dx +
					dy*dy +
					dz*dz,
			)

			weapon := ""

			if p.ActiveWeapon() != nil {
				weapon = p.ActiveWeapon().String()
			}

			snapshot.TradeCandidates = append(
				snapshot.TradeCandidates,
				TradeCandidate{

					Rank: 0,

					Name:      p.Name,
					SteamID64: p.SteamID64,

					X: pos.X,
					Y: pos.Y,
					Z: pos.Z,

					Distance: distance,

					HeightDifference: pos.Z - e.Victim.Position().Z,

					Alive: p.IsAlive(),

					HP:    p.Health(),
					Armor: p.Armor(),

					Weapon: weapon,

					Money: p.Money(),

					FlashDuration: float64(p.FlashDuration),

					Scoped: p.IsScoped(),

					Ducking: p.IsDucking(),

					Velocity: 0,

					ViewX: p.ViewDirectionX(),
					ViewY: p.ViewDirectionY(),

					// ========= Killer Snapshot =========

					KillerViewX: e.Killer.ViewDirectionX(),
					KillerViewY: e.Killer.ViewDirectionY(),

					KillerHP:    e.Killer.Health(),
					KillerArmor: e.Killer.Armor(),

					KillerScoped: e.Killer.IsScoped(),

					KillerFlashDuration: float64(e.Killer.FlashDuration),

					KillerPlace:  e.Killer.LastPlaceName(),
					KillerWeapon: killerWeapon,

					KillerMoney: e.Killer.Money(),

					KillerX: e.Killer.Position().X,
					KillerY: e.Killer.Position().Y,
					KillerZ: e.Killer.Position().Z,
				},
			)
		}

		SortTradeCandidates(snapshot.TradeCandidates)

		deathSnapshots = append(deathSnapshots, snapshot)

		for i := len(match.Shots) - 1; i >= 0; i-- {

			if match.Shots[i].SteamID64 != e.Killer.SteamID64 {
				continue
			}

			if match.Shots[i].Kill {
				continue
			}

			if !match.Shots[i].Hit {
				continue
			}

			if match.Shots[i].VictimSteamID64 != e.Victim.SteamID64 {
				continue
			}

			match.Shots[i].Kill = true

			match.Shots[i].Headshot = e.IsHeadshot

			match.Shots[i].Wallbang = e.IsWallBang()

			match.Shots[i].ThroughSmoke = e.ThroughSmoke

			match.Shots[i].Blind = e.AttackerBlind

			match.Shots[i].NoScope = e.NoScope

			match.Shots[i].FlashAssist = e.AssistedFlash

			match.Shots[i].Distance = e.Distance

			match.Shots[i].Penetration = e.PenetratedObjects

			match.Shots[i].TimeToKill =
				parser.GameState().IngameTick() -
					match.Shots[i].FireTick

			break
		}

		for i := len(match.Shots) - 1; i >= 0; i-- {

			if match.Shots[i].SteamID64 != e.Killer.SteamID64 {
				continue
			}

			if match.Shots[i].VictimSteamID64 != e.Victim.SteamID64 {
				continue
			}

			match.Shots[i].Kill = true

			match.Shots[i].Headshot = e.IsHeadshot

			match.Shots[i].Wallbang = e.IsWallBang()

			match.Shots[i].ThroughSmoke = e.ThroughSmoke

			match.Shots[i].Blind = e.AttackerBlind

			match.Shots[i].NoScope = e.NoScope

			match.Shots[i].FlashAssist = e.AssistedFlash

			match.Shots[i].Distance = e.Distance

			match.Shots[i].Penetration = e.PenetratedObjects

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

		if e.Assister != nil {

			if s, ok := stats[e.Assister.SteamID64]; ok {
				s.Assists++
			}
		}
	})

	parser.RegisterEventHandler(func(e events.RoundEnd) {

		if !matchStarted {
			return
		}

		roundNumber++

		state := parser.GameState()

		match.Rounds = append(match.Rounds, Round{
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

	match = BuildMatch(
		parser,
		mapName,
	)

	teams = BuildStableTeamsFromPlayers(match.Players)

	ApplyTeamGroupsToPlayers(match.Players, teams)

	for i := range match.Shots {

		if !match.Shots[i].Hit {

			match.Shots[i].Miss = true
		}
	}
	if err != nil {
		output(Result{
			Success: false,
			Message: err.Error(),
		})
		return err
	}

	fmt.Fprintln(os.Stderr, "Kill Count:", len(match.Kills))

	var statsList []PlayerStats

	for _, s := range stats {
		statsList = append(statsList, *s)
	}

	result := Result{
		Success: true,

		Metadata: Metadata{
			ParserVersion:     "0.1.0",
			Game:              "Counter-Strike 2",
			DemoFormat:        "CS2 Demo",
			DemoinfocsVersion: "5.2.0",

			TimeUnit:     "tick",
			DistanceUnit: "hammer_units",
			PositionUnit: "hammer_units",
			VelocityUnit: "hammer_units_per_second",

			TickRate: int(parser.TickRate()),
		},

		Match: match,

		Players: match.Players,

		Teams: teams,

		Kills: match.Kills,

		DeathSnapshots: deathSnapshots,

		Damages: match.Damages,

		WeaponFires: weaponFires,

		Shots: match.Shots,

		PositionCache: nil,

		Rounds: match.Rounds,

		Stats: statsList,
	}

	positionCache, err = BuildPositionCache(path)

	if err != nil {
		return err
	}

	result.PositionCache = positionCache

	ctx := BuildContext(&result)

	ValidateResult(&result)

	NormalizeResultNames(&result)

	Analyze(ctx, &result)

	BuildOpeningStats(&result)

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

	switch debug {

	case DebugTrade:

		PrintTradeDebug(&result)

	case DebugAPI:

		PrintAPIDebug(&result)

	default:

		output(result)

	}

	return nil
}
