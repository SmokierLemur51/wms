	/*
	Trying to remove some redundency and too much copy & pasting

	Contains:
		CheckExistence()
		FindDatabaseID()
		ToTitleCase() << pairs nicely with the above 
*/

package data

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

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
			fmt.Println(err)
			return false, err
		}
	}
	if count > 0 {
		fmt.Printf("<%s> exists in table <%s>.\n", item, table)
		return true, err
	}
	// if not
	fmt.Printf("<%s> does not exist in table <%s>.\n", item, table)
	return false, nil
}

func FindDatabaseID(db *sql.DB, table, column, item string) int {
	rows, err := db.Query(fmt.Sprintf("SELECT id FROM %s WHERE %s = ?", table, column), item)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var id int
	for rows.Next() {
		err := rows.Scan(&id)
		if err == sql.ErrNoRows {
			fmt.Printf("\nError: %s\nItem: %s does not exist.\n", err, item)
		}
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	fmt.Printf("Found item <%s> with an ID of <%d>", item, id)
	return id
}

func ToTitleCase(s string) string {
	// might never use this <3
	return strings.Title(strings.ToLower(s))
}
