package database

import "fmt"

type Match struct {
	DemoName string
	Map      string
	GameMode string
	TickRate int
	Duration int
	ScoreCT  int
	ScoreT   int
	Winner   string
}

// SaveMatch 保存一场比赛
func SaveMatch(match Match) (int64, error) {

	db := GetDB()

	result, err := db.Exec(`
	INSERT INTO matches
	(
		demo_name,
		map,
		game_mode,
		tick_rate,
		duration,
		score_ct,
		score_t,
		winner
	)
	VALUES
	(
		?,
		?,
		?,
		?,
		?,
		?,
		?,
		?
	)
	`,
		match.DemoName,
		match.Map,
		match.GameMode,
		match.TickRate,
		match.Duration,
		match.ScoreCT,
		match.ScoreT,
		match.Winner,
	)

	if err != nil {
		return 0, fmt.Errorf("save match failed: %w", err)
	}

	return result.LastInsertId()
}
