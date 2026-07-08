package main

import (
	"encoding/json"
	"fmt"
	"os"

	dem "github.com/markus-wa/demoinfocs-golang/v5/pkg/demoinfocs"
	"github.com/markus-wa/demoinfocs-golang/v5/pkg/demoinfocs/events"
)

func CollectPlayerStates(path string) ([]PlayerState, error) {

	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	parser := dem.NewParser(f)
	defer parser.Close()

	var positions []PlayerState

	lastPosition := make(map[uint64]PlayerState)

	lastTick := -1

	parser.RegisterEventHandler(func(e events.FrameDone) {

		state := parser.GameState()

		tick := state.IngameTick()

		if tick == lastTick {
			return
		}

		lastTick = tick

		if tick%4 != 0 {
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
				Tick: tick,

				PlayerID: id,

				X: int16(pos.X),

				Y: int16(pos.Y),

				Z: int16(pos.Z),

				HP: uint8(p.Health()),

				Armor: uint8(p.Armor()),

				Alive: p.IsAlive(),
			}

			last, ok := lastPosition[p.SteamID64]

			if ok {

				dx := current.X - last.X
				dy := current.Y - last.Y

				moved := dx*dx+dy*dy > 16*16

				hpChanged := current.HP != last.HP

				armorChanged := current.Armor != last.Armor

				if !moved && !hpChanged && !armorChanged {

					continue
				}
			}

			positions = append(positions, current)

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

	count := len(positions)

	fmt.Fprintln(os.Stderr, "Positions Collected:", count)

	data, err := json.MarshalIndent(positions, "", "  ")
	if err != nil {
		return nil, err
	}

	err = os.WriteFile("player_states.json", data, 0644)
	if err != nil {
		return nil, err
	}

	fmt.Fprintln(os.Stderr, "Saved to player_states.json")

	return positions, nil

}
