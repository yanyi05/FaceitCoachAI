package main

import (
	"encoding/json"
	"os"

	"github.com/markus-wa/demoinfocs-golang/v5/pkg/demoinfocs/common"
	"github.com/markus-wa/demoinfocs-golang/v5/pkg/demoinfocs/events"
)

func main() {

	if len(os.Args) < 3 {

		println("Usage:")
		println("parser.exe event demo.dem")
		println("parser.exe frame demo.dem")

		return
	}

	mode := os.Args[1]
	demo := os.Args[2]

	switch mode {

	case "event":

		err := EventParser(demo)

		if err != nil {
			panic(err)
		}

	case "frame":

		_, err := CollectPlayerStates(demo)

		if err != nil {
			panic(err)
		}

	default:

		println("Unknown mode")
	}
}

func output(v any) {

	data, _ := json.MarshalIndent(v, "", "    ")

	os.Stdout.Write(data)
}
func teamName(team common.Team) string {
	switch team {
	case common.TeamCounterTerrorists:
		return "CT"
	case common.TeamTerrorists:
		return "T"
	case common.TeamSpectators:
		return "Spectator"
	default:
		return "Unknown"
	}
}
func hitGroupName(hit events.HitGroup) string {
	switch hit {
	case events.HitGroupHead:
		return "Head"
	case events.HitGroupChest:
		return "Chest"
	case events.HitGroupStomach:
		return "Stomach"
	case events.HitGroupLeftArm:
		return "LeftArm"
	case events.HitGroupRightArm:
		return "RightArm"
	case events.HitGroupLeftLeg:
		return "LeftLeg"
	case events.HitGroupRightLeg:
		return "RightLeg"
	case events.HitGroupGear:
		return "Gear"
	default:
		return "Generic"
	}
}
