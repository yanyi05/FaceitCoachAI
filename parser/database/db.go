package database

import (
	"database/sql"
	"sync"

	_ "modernc.org/sqlite"
)

var (
	db   *sql.DB
	once sync.Once
)

// Init 初始化数据库，只会执行一次
func Init() error {
	var err error

	once.Do(func() {
		db, err = sql.Open("sqlite", "faceitcoach.db")
		if err != nil {
			return
		}

		err = db.Ping()
	})

	return err
}

// GetDB 返回数据库连接
func GetDB() *sql.DB {
	return db
}

// Close 关闭数据库
func Close() error {
	if db == nil {
		return nil
	}

	return db.Close()
}
