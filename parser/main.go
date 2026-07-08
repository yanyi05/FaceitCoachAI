package main

import (
	"encoding/json"
	"os"

	"github.com/markus-wa/demoinfocs-golang/v5/pkg/demoinfocs/common"
	"github.com/markus-wa/demoinfocs-golang/v5/pkg/demoinfocs/events"

	"faceitcoachai/parser/database"
)

func main() {

	if err := database.Init(); err != nil {
		panic(err)
	}

	if err := database.Migrate(); err != nil {
		panic(err)
	}

	defer database.Close()

	if len(os.Args) < 2 {

		println("Usage:")
		println("parser.exe demo.dem")
		println("parser.exe event demo.dem")
		println("parser.exe frame demo.dem")

		return
	}

	mode := "event"
	demo := os.Args[1]

	if len(os.Args) >= 3 {
		mode = os.Args[1]
		demo = os.Args[2]
	}

	switch mode {

	case "event":

		err := EventParser(demo, DebugNone)

		if err != nil {
			panic(err)
		}

	case "debug-trade":

		err := EventParser(demo, DebugTrade)

		if err != nil {
			panic(err)
		}

	case "debug-api":

		err := EventParser(demo, DebugAPI)

		if err != nil {
			panic(err)
		}

	case "frame":

		_, err := BuildPositionCache(demo)

		if err != nil {
			panic(err)
		}

	default:

		println("Unknown mode")
	}
}

func output(v any) {
	data, err := json.MarshalIndent(v, "", "    ")
	if err != nil {
		panic(err)
	}

	// 1. 输出到终端（React 会用）
	os.Stdout.Write(data)

	// 2. 自动建立 output 目录
	os.MkdirAll("output", 0755)

	// 3. 写入 JSON 文件
	err = os.WriteFile("output/event.json", data, 0644)
	if err != nil {
		panic(err)
	}
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
