package data

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

// What all do we need?

// done		insert into database (checking if it exists first)
// done		calculate cost per item based off of pallet

// update in the database
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
	Status		  int
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
	// Locations     []WarehouseBin
}

func RandomProducts() []Product {
	return []Product{
		{Product: "Pops", ProductCode: "874", Description: "Kevin Malone - The Office",
			UnitsCtn: 12, CtnPallet: 48, CostPallet: 4500.00},
		{Product: "HydroFlask", ProductCode: "jhafdkljklfjdlk", Description: "Drink water, aesthetically.",
			UnitsCtn: 6, CtnPallet: 45, CostPallet: 6200.00},
	}	
}


func PopulateProducts(db *sql.DB, p []Product) {
	for _, product := range p {
		product.InsertProduct(db)
	}
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
			"INSERT INTO products (product, product_code, p_description, units_ctn, ctn_pallet, units_pallet, cost_pallet, cost_ctn, cost_unit) VALUES (?,?,?,?,?,?,?,?,?)",
			p.Product, p.ProductCode, p.Description, p.UnitsCtn, p.CtnPallet, (p.CtnPallet*p.UnitsCtn), 
			p.CostPallet, (p.CostPallet/(float64(p.CtnPallet))), ((p.CostPallet/float64(p.CtnPallet))/float64(p.UnitsCtn)),
		)
		if err != nil {
			log.Fatal(err)
		}
	case true:
		fmt.Printf("Product %s already exists.\n", p.Product)
	}

}

func (p *Product) UpdateCostPallet(cost float64) {}

func (p *Product) RecieveToLocation() {}

func (p *Product) UpdateLocation() {}

func (p *Product) UpdatePricing(margin float64) {} 

