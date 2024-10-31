package main

import (
	"database/sql"
	"flag"
	"log"
	"log/slog"

	_ "github.com/go-sql-driver/mysql"
)

var userStmt = `CREATE TABLE users(id INTEGER NOT NULL PRIMARY KEY AUTO_INCREMENT,
 name VARCHAR(255) NOT NULL,
 email VARCHAR(255) NOT NULL,
 hashed_password CHAR(60) NOT NULL,
 created DATETIME NOT NULL);
 `
var uniqueEmailStmt = "ALTER TABLE users ADD CONSTRAINT users_uc_email UNIQUE (email);"

func executeStmts(db *sql.DB, stmts ...string) error {
	for _, stmt := range stmts {
		_, err := db.Exec(stmt)
		if err != nil {
			return err
		}
	}
	return nil
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
	dsn := flag.String("dsn", "root:root@/snippetbox?parseTime=true", "snippetbox web user connection")
	flag.Parse()

	DB, err := createConnectionPool(dsn)
	if err != nil {
		slog.Error(err.Error())
	}

	err = executeStmts(DB, userStmt, uniqueEmailStmt)
	if err != nil {
		log.Fatal(err)
	}

}
