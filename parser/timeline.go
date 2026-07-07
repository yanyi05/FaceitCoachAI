package main

type KillTimeline struct {
	Tick int

	Killer uint64
	Victim uint64

	KillIndex int
}

func BuildKillTimeline(result *Result) []KillTimeline {

	timeline := make([]KillTimeline, 0, len(result.Kills))

	for i, k := range result.Kills {

		timeline = append(timeline, KillTimeline{

			Tick: k.Tick,

			Killer: k.KillerSteamID64,

			Victim: k.VictimSteamID64,

			KillIndex: i,
		})
	}

	return timeline
}
