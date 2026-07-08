package database

import "fmt"

// Migrate 自动创建数据库表
func Migrate() error {
	db := GetDB()

	query := `
	CREATE TABLE IF NOT EXISTS matches (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		demo_name TEXT NOT NULL,
		map TEXT,
		game_mode TEXT,
		tick_rate INTEGER,
		duration INTEGER,
		score_ct INTEGER,
		score_t INTEGER,
		winner TEXT,
		created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);
	`

	_, err := db.Exec(query)
	if err != nil {
		return fmt.Errorf("create matches table failed: %w", err)
	}

	return nil
}
