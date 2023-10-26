package infrastructure

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type DB struct {
	Read  *sqlx.DB
	Write *sqlx.DB
}

func NewDB() (*DB, error) {
	dbConnect := fmt.Sprintf("%s:%s@tcp(%s)/%s?parseTime=true&loc=Local&charset-utf8&timeout=%s&tls=%s",
		"app-user",
		"test",
		"db:3306",
		"household_account_book",
		"30s",
		"false",
	)

	readDB, err := sqlx.Open("mysql", dbConnect)
	if err != nil {
		return nil, fmt.Errorf("sqlx.Open: %w", err)
	}
	writeDB, err := sqlx.Open("mysql", dbConnect)
	if err != nil {
		return nil, fmt.Errorf("sqlx.Open: %w", err)
	}

	return &DB{
		Read:  readDB,
		Write: writeDB,
	}, nil
}

func (db *DB) Close() {
	_ = db.Read.Close()
	_ = db.Write.Close()
}
