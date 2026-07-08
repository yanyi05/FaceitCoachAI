package main

func BuildContext(result *Result) *AnalyzerContext {

	ctx := &AnalyzerContext{

		PlayerBySteam: make(map[uint64]*Player),

		PlayerIDBySteam: make(map[uint64]uint8),

		PlayerByID: make(map[uint8]*Player),

		PlayerStateByPlayer: make(map[uint8][]PlayerState),

		KillsByRound: make(map[int][]Kill),

		DamagesByRound: make(map[int][]Damage),
	}

	// 建立 Player 索引
	for i := range result.Players {

		player := &result.Players[i]

		ctx.PlayerBySteam[player.SteamID64] = player
		ctx.PlayerByID[player.PlayerID] = player
		ctx.PlayerIDBySteam[player.SteamID64] = player.PlayerID
	}

	// 建立 PlayerState 索引
	if result.PositionCache != nil {

		for _, frame := range result.PositionCache.Frames {

			for _, state := range frame.Players {

				ctx.PlayerStateByPlayer[state.PlayerID] =
					append(ctx.PlayerStateByPlayer[state.PlayerID], state)
			}
		}
	}

	return ctx
}
