/*
	Trying to remove some redundency and too much copy & pasting

*/

package data

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)


func CheckExistence(db *sql.DB, table, column, item string) (bool, error) {
	// returns true if it exists
	var count int
	rows, err := db.Query("SELECT COUNT(*) FROM ? WHERE ? = ?", table, column, item)
	if err != nil {
		return true, err
	}
	defer rows.Close()
	for rows.Next() {
		if err := rows.Scan(&count); err != nil {
			return false, err
		}
	}
	if count > 0 {
		return true, err
	}
	// if not
	return false, nil
}

func FindDatabaseID(db *sql.DB, table, column, item string) (int) {
	rows, err := db.Query(fmt.Sprintf("SELECT id FROM %s WHERE %s = ?", table, column), item)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var id int
	for rows.Next() {
		err := rows.Scan(&id)
		if err == sql.ErrNoRows {
			fmt.Printf("\nError: %s\nItem: %s does not exist.\n")
		}
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}

	return id
}