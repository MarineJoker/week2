package main

import (
	"database/sql"
	"fmt"
	"github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

var ErrRecordNotFound = errors.New("record not found")

type User struct {
	ID     int
	Name  string
}

func queryByID(db *sql.DB, id int) (*User, error) {
	// SQL query
	var user User
	query := `SELECT id, name
    FROM users
    WHERE id= $1
    `
	err := db.QueryRow(query, id).Scan(&user.ID, &user.Name)
	if err!=nil {
		switch {
		case errors.Is(err, sql.ErrNoRows):
			return nil, errors.Wrapf(ErrRecordNotFound, "sql: %s %d", query, id)
		default:
			return nil, errors.Wrapf(err, "sql: %s %d", query, id)
		}
	}
	return &user, nil
}

func main() {
	cfg := mysql.Config{
		User:   "root",
		Passwd: "xxx",
		Addr:   "localhost:3306",
		DBName: "xxx",
		Net:    "tcp",
	}
	db, err := sql.Open("mysql", cfg.FormatDSN())
	if err != nil {
		fmt.Errorf("%w", err)
	}
	fmt.Println(queryByID(db, 27))
	if err != nil {
		fmt.Println(err)
		return
	}
}