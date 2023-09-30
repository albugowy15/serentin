package db

import (
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

type Database struct {
	Conn *sqlx.DB
}

func Connect(host string, port int, username string, password string, database string) *Database {
	connectionStr := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s", username, password, host, port, database)
	conn := sqlx.MustConnect("mysql", connectionStr)
	fmt.Println("Connected to database")
	return &Database{
		Conn: conn,
	}
}

func (db *Database) Close() {
	db.Conn.Close()
}
