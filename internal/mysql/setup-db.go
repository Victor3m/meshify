package mysql

import (
	"database/sql"
	"log/slog"
	"strconv"
)

func SetupDB(db *sql.DB) error {
	logSQLResult(db.Exec("CREATE DATABASE IF NOT EXISTS meshify"))
	logSQLResult(db.Exec("USE meshify"))
	logSQLResult(db.Exec("CREATE TABLE IF NOT EXISTS users (id INT NOT NULL AUTO_INCREMENT, username VARCHAR(20) NOT NULL, password VARCHAR(20) NOT NULL, PRIMARY KEY (id))"))
	logSQLResult(db.Exec("CREATE TABLE IF NOT EXISTS user_tokens (id INT NOT NULL AUTO_INCREMENT, token VARCHAR(200) NOT NULL, user_id INT NOT NULL, PRIMARY KEY (id), FOREIGN KEY (user_id) REFERENCES users(id))"))

	return nil
}

func logSQLResult(result sql.Result, err error) {
	if err != nil {
		slog.Error(err.Error())
	} else {
		insertId, _ := result.LastInsertId()
		rowsAffected, _ := result.RowsAffected()
		slog.Info("insertId: %s rowsAffected: %s", strconv.Itoa(int(insertId)), strconv.Itoa(int(rowsAffected)))
	}
}
