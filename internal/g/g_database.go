package g

import (
	"database/sql"
	"sync"

	_ "github.com/jackc/pgx/v5"
)

var (
	rdb    *sql.DB
	wdb    *sql.DB
	dbOnce sync.Once
)

func DB(read ...bool) *sql.DB {
	dbOnce.Do(func() {
		var err error
		wdb, err = sql.Open("pgx", Cfg().Database.Write)
		if err != nil {
			panic(err)
		}
		if Cfg().Database.Read != "" {
			rdb, err = sql.Open("pgx", Cfg().Database.Read)
			if err != nil {
				panic(err)
			}
		}
	})
	if len(read) > 0 && read[0] && rdb != nil {
		return rdb
	}
	return wdb
}

func RDB() *sql.DB {
	return DB(true)
}

func WDB() *sql.DB {
	return DB()
}

func DbClose() error {
	if err := wdb.Close(); err != nil {
		return err
	}
	if rdb != nil {
		return rdb.Close()
	}
	return nil
}
