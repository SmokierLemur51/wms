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


type Product struct {
	Id            int
	VendorId      int
	Status		  int
	Category      int
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
	// Locations 	   []WarehouseBin   
	Location []Shelf 	  
}

func CostToSelling(cost, margin float64) float64 {
	return cost / (1.00-(margin/100.00))
}

func RandomProducts() []Product {
	return []Product{
		{},
	}	
}


func PopulateProducts(db *sql.DB, p []Product) {
	for _, product := range p {
		product.InsertProduct(db)
	}
}

func CheckExistingProduct(db *sql.DB, product string) (bool, error) {
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
	execute, err = CheckExistingProduct(db, p.Product)
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


// functions
func LoadAllStockProducts(db *sql.DB) []Product {
	var products []Product
	rows, err := db.Query("SELECT * FROM products")
	if err != nil {
		log.Fatal(err)
	}
	for rows.Next() {
		var p Product
		if err := rows.Scan(&p.Id, &p.VendorId, &p.Status, &p.Product, &p.ProductCode, &p.Description, &p.UnitsCtn, 
			&p.CtnPallet, &p.UnitsPallet, &p.CostPallet, &p.SellingPallet, &p.CostCtn, &p.SellingCtn, 
			&p.CostUnit, &p.SellingUnit, &p.Location); err != nil {
				log.Fatal(err)
			}

		fmt.Println(p.Product)
		products = append(products, p)
	}
	return products
}
