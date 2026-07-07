package main

func BuildContext(result *Result) *AnalyzerContext {

	ctx := &AnalyzerContext{

		PlayerBySteam: make(map[uint64]*Player),

		PlayerByID: make(map[uint8]*Player),

		PlayerStateByPlayer: make(map[uint8][]PlayerState),

		KillsByRound: make(map[int][]Kill),

		DamagesByRound: make(map[int][]Damage),
	}

	// 建立 Player 索引
	for i := range result.Players {

		player := &result.Players[i]

		ctx.PlayerBySteam[player.SteamID64] = player
	}

	// 建立 PlayerState 索引
	for _, state := range result.PlayerStates {

		ctx.PlayerStateByPlayer[state.PlayerID] =
			append(ctx.PlayerStateByPlayer[state.PlayerID], state)
	}

	return ctx
}
