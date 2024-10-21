package main

import (
	"database/sql"
	"flag"
	"log/slog"

	_ "github.com/go-sql-driver/mysql"
)

var userStmt = `CREATE TABEL users(id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
 name VARCHAR(255) NOT NULL,
 email VARCHAR(255) NOT NULL,
 hashed_password CHAR(60) NOT NULL,
 created DATETIME NOT NULL);
 ALTER TABLE users ADD CONSTRAINT users_uc_email UNIQUE (email);
`

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

	DB, err := createConnectionPool(dsn)
	if err != nil {
		slog.Error(err.Error())
	}

	createTabel(DB, userStmt)

}
