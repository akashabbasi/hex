package db

import (
	"database/sql"
	"log"
	"time"

	sq "github.com/Masterminds/squirrel"
	_ "github.com/go-sql-driver/mysql"
)

type Adapter struct {
	db *sql.DB
}

func NewAdapter(drivername, dsn string) *Adapter {
	// connect
	db, err := sql.Open(drivername, dsn)
	if err != nil {
		log.Fatalf("db connection failure: %v", err)
	}

	err = db.Ping()
	if err != nil {
		log.Fatal("db ping failure: %v", err)
	}

	return &Adapter{db: db}
}

func (da Adapter) CloseDbConnection() {
	if err := da.db.Close(); err != nil {
		log.Fatalf("db close failure: %v", err)
	}
}

func (da Adapter) AddToHistory(answer int32, operation string) error {
	queryStr, args, err := sq.Insert("arith_history").
		Columns("date", "answer", "operation").
		Values(
			time.Now(), answer, operation,
		).ToSql()
	if err != nil {
		return err
	}

	_, err = da.db.Exec(queryStr, args...)
	if err != nil {
		return err
	}

	return nil
}
