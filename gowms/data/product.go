package data

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// What all do we need?

// insert into database (checking if it exists first)
// update in the database
// calculate cost per item based off of pallet
// 		which i think should be done after recieved into warehouse by tms system
// know warehouse locations, how many are in each bin
// track daily sales of the item for end of day reports
// keep enough in stock to not run out
// know when deliveries are coming in
// average time from ordering to recieving

type Unit struct {
	Id          int
	Unit        string
	Description string
}

type Product struct {
	Id            int
	VendorId      int
	Product       string
	ProductCode   string
	Description   string
	UnitsCtn      int
	CtnPallet     int
	UnitsPallet   int
	CostPallet    float64
	SellingPallet float64
	CostCtn       float64
	SellingCtn    float64
	CostUnit      float64
	SellingUnit   float64
	// locations []WarehouseBin
}

func (p *Product) CostPerItem() {
	// takes the total spent on a pallet, ctn per pallet, and unitsCTN to calcualte cost per <3
	p.UnitsPallet = (p.CtnPallet * p.UnitsCtn)
	// p.CostPallet
}

func CheckExisting(db *sql.DB, product string) (bool, error) {
	// returns true if it exists
	var count int
	rows, err := db.Query("SELECT COUNT(*) FROM products WHERE product = ? ", product)
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

// needs complete rewrite of the insert, maybe need all data first
func (p Product) InsertProduct(db *sql.DB) {
	var execute bool
	var err error
	execute, err = CheckExisting(db, p.Product)
	if err != nil {
		log.Println(err)
		return
	}

	// remember, the check existing returns true if the product already exists, so it skips
	switch execute {
	case false:
		_, err := db.Exec(
			"INSERT INTO products (product, product_code, p_description, cost_pallet, units_ctn, ctn_pallet) VALUES (?,?,?,?,?,?,?,?,?)",
			p.Product, p.ProductCode, p.Description, p.CostPallet, p.UnitsCtn, p.CtnPallet,
		)
		// create table products (
		// 	id integer primary key autoincrement,
		// 	vendor_id integer,
		// 	product varchar(50) not null,
		// 	product_code varchar(60) not null,
		// 	p_description text,
		// 	units_ctn integer,
		// 	ctn_pallet integer,
		// 	units_pallet integer,
		// 	cost_pallet real,
		// 	selling_pallet real,
		// 	cost_ctn real,
		// 	selling_ctn real,
		// 	cost_unit real,
		// 	selling_unit real,
		// 	wh_bin_id integer,
		// 	foreign key (vendor_id) references vendors(id),
		// 	foreign key (wh_bin_id) references warehouse_bin(id)
		// );
		if err != nil {
			log.Fatal(err)
		}
	case true:
		fmt.Printf("Product %s already exists.\n", p.Product)
	}

}

func (p *Product) RecieveToLocation() {}
