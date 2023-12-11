package data

import (
	"fmt"
	"log"
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
)

type Vendor struct {
	Id 		int
	Vendor  string
	Street1 string
	Street2 string
	City	string
	State	string
	Zip		string
	StockProducts []Product
	SpecialOrderProducts []Product
}

// populating random values for testing
func RandomVendors() []Vendor {
	return []Vendor{
		{Vendor: "Mars, Incorporated", Street1: "6885 Elm Street McLean", City: "McLean", State: "VA", Zip: "22101"},
		{Vendor: "Ford Motor Company", Street1: "1 American Rd", City: "Dearborn", State: "MI", Zip: "48126"},
		{Vendor: "Dell", Street1: "1 Dell Way", City:"Round Rock", State: "TX", Zip:"78664"},
	}
}

func PopulateVendors(db *sql.DB, v []Vendor) {
	for _, vendor := range v {
		vendor.InsertVendor(db)
	}
}

// vendor methods
func (v Vendor) InsertVendor(db *sql.DB, ) {
	var execute bool
	var err error
	execute, err = CheckExisting(db, v.Vendor)
	if err != nil {
		log.Println(err)
		return
	}
	// remember, the check existing returns true if the vendor already exists, so it skips
	switch execute {
	case false:
		_, err := db.Exec(
			"INSERT INTO vendors (vendor, address_street, address_street_2, address_city, address_state, address_zip) VALUES (?,?,?,?,?,?)",
			v.Vendor, v.Street1, v.Street2, v.City, v.State, v.Zip,
		)
		if err != nil {
			log.Fatal(err)
		}
	case true:
		fmt.Printf("Vendor %s already exists.\n", v.Vendor)
	}
}

func (v Vendor) LoadAllStockedProducts() {}

func (v Vendor) LoadSpecialOrderProducts() {}

// queries
func CheckVendorExistence(db *sql.DB, vendor string) (bool, error) {
	// returns true if it exists
	var count int
	rows, err := db.Query("SELECT COUNT(*) FROM vendors WHERE vendor = ? ", vendor)
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

func FindVendorId(db *sql.DB, v string) Vendor {}