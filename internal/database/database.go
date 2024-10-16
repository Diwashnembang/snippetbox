package main

import (
	"database/sql"
	"flag"
	"log/slog"

	_ "github.com/go-sql-driver/mysql"
)

var userStmt = `CREATE TABEL users(id INT PRIMARY KEY,email TEXT NOT NULL, password TEXT NOT NULL)`

func createTabel(db *sql.DB, stmt string) {
	db.Exec(stmt)
}

func createConnectionPool(dsn *string) (*sql.DB, error) {
	DB, err := sql.Open("mysql", *dsn)
	if err != nil {
		slog.Error(err.Error())
		return nil, err
	}
	if err := DB.Ping(); err != nil {
		slog.Error(err.Error())
		return nil, err
	}
	return DB, nil
}

func main() {
	dsn := flag.String("dsn", "web:pass@/snippetbox?parseTime=true", "snippetbox web user connection")
	flag.Parse()

	_, err := createConnectionPool(dsn)
	if err != nil {
		slog.Error(err.Error())
	}

	// createTabel(DB,userStmt)

}
