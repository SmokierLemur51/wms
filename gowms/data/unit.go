package data

import (
	"database/sql"
	// "fmt"
	// "log"

	_ "github.com/mattn/go-sqlite3"
)


type Unit struct {
	Id          int
	Unit        string
	Description string
}

// I think a status struct may be an unnecesarry use of resources, a database query may suffice.
type Status struct {}

func FindStatusId(db *sql.DB, s string) (int) {
	return 1
}