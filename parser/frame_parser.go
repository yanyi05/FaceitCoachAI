package main

import (
	"encoding/json"
	"fmt"
	"math"
	"os"

	dem "github.com/markus-wa/demoinfocs-golang/v5/pkg/demoinfocs"
	"github.com/markus-wa/demoinfocs-golang/v5/pkg/demoinfocs/events"
)

func BuildPositionCache(path string) (*PositionCache, error) {

	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	parser := dem.NewParser(f)
	defer parser.Close()

	cache := PositionCache{
		Frames: make(map[int]*TickFrame),
	}

	lastPosition := make(map[uint64]PlayerState)

	lastTick := -1

	parser.RegisterEventHandler(func(e events.FrameDone) {

		state := parser.GameState()

		tick := state.IngameTick()

		if tick == lastTick {
			return
		}

		lastTick = tick

		if tick%PositionCacheInterval != 0 {
			return
		}

		for _, p := range state.Participants().Playing() {

			if p == nil {
				continue
			}

			if p.SteamID64 == 0 {
				continue
			}

			pos := p.Position()

			id := GetPlayerID(p.SteamID64)

			current := PlayerState{
				LastPlace: p.LastPlaceName(),
				Tick:      tick,
				Round:     state.TotalRoundsPlayed() + 1,

				PlayerID:  id,
				SteamID64: p.SteamID64,
				Team:      teamName(p.Team),

				X: int16(pos.X),
				Y: int16(pos.Y),
				Z: int16(pos.Z),

				ViewYaw:   float32(p.ViewDirectionX()),
				ViewPitch: float32(p.ViewDirectionY()),

				Alive: p.IsAlive(),

				IsWalking:  p.IsWalking(),
				IsStanding: p.IsStanding(),
				IsDucking:  p.IsDucking(),
				IsAirborne: p.IsAirborne(),

				IsScoped:      p.IsScoped(),
				IsBlinded:     p.IsBlinded(),
				FlashDuration: float32(p.FlashDurationTimeRemaining().Seconds()),

				HasHelmet:    p.HasHelmet(),
				HasDefuseKit: p.HasDefuseKit(),

				IsPlanting: p.IsPlanting,
				IsDefusing: p.IsDefusing,

				HP:    uint8(p.Health()),
				Armor: uint8(p.Armor()),
				Money: p.Money(),
			}

			weapon := p.ActiveWeapon()

			if weapon != nil {
				current.Weapon = weapon.String()
				current.AmmoInMagazine = weapon.AmmoInMagazine()
				current.AmmoReserve = weapon.AmmoReserve()
				current.ZoomLevel = int(weapon.ZoomLevel())
			}

			last, ok := lastPosition[p.SteamID64]

			if ok {

				dx := int(current.X) - int(last.X)
				dy := int(current.Y) - int(last.Y)
				dz := int(current.Z) - int(last.Z)

				movementDistance := float32(math.Sqrt(
					float64(dx*dx + dy*dy + dz*dz),
				))

				sampleTime := float32(PositionCacheInterval) / float32(parser.TickRate())

				current.Velocity = movementDistance / sampleTime

				current.VelocityX = float32(dx) / sampleTime
				current.VelocityY = float32(dy) / sampleTime
				current.VelocityZ = float32(dz) / sampleTime

				moved := current.Velocity > 16

				hpChanged := current.HP != last.HP
				armorChanged := current.Armor != last.Armor

				if !moved && !hpChanged && !armorChanged {
					continue
				}
			}

			frame, ok := cache.Frames[current.Tick]

			if !ok {
				frame = &TickFrame{
					Tick: current.Tick,
				}
				cache.Frames[current.Tick] = frame
			}

			frame.Players = append(frame.Players, current)

			lastPosition[p.SteamID64] = current
		}

	})

	lastPrintTick := -1000

	for {

		ok, err := parser.ParseNextFrame()

		if err != nil {
			return nil, err
		}

		if !ok {
			break
		}

		tick := parser.GameState().IngameTick()

		if tick != lastPrintTick && tick%1000 == 0 {

			fmt.Fprintln(os.Stderr, "Tick:", tick)

			lastPrintTick = tick
		}
	}

	fmt.Fprintln(os.Stderr, "Finished.")

	playerCount := 0

	for _, frame := range cache.Frames {
		playerCount += len(frame.Players)
	}

	fmt.Fprintln(os.Stderr, "Frames:", len(cache.Frames))
	fmt.Fprintln(os.Stderr, "Player States:", playerCount)

	data, err := json.MarshalIndent(cache, "", "  ")

	if err != nil {
		return nil, err
	}

	err = os.WriteFile("player_states.json", data, 0644)
	if err != nil {
		return nil, err
	}

	fmt.Fprintln(os.Stderr, "Saved to player_states.json")

	return &cache, nil

}
