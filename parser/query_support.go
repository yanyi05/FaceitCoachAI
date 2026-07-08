package main

func FindSupportByVictim(

	supports []SupportAnalysis,

	victimSteamID uint64,

	round int,

) *SupportAnalysis {

	for i := range supports {

		if supports[i].VictimSteamID == victimSteamID &&
			supports[i].Round == round {

			return &supports[i]
		}
	}

	return nil
}
