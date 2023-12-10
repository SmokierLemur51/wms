package data

import (
	"database/sql"
	"log"
	"fmt"
	_ "github.com/mattn/go-sqlite3"
)


// need to add function to calculate cost per item by cost paid for pallet/ctn of item
type Product struct {
	Id int
	Product string
	ProductCode string
	Description string
	UnitsCtn int
	CtnPallet int
	UnitsPallet int
	Selling float64
	CostPallet float64
	CostCtn float64
}

func (p *Product) CostPerItem() {
	// takes the total spent on a pallet, ctn per pallet, and unitsCTN to calcualte cost per <3
	p.UnitsPallet = (p.CtnPallet * p.UnitsCtn)
	p.CostPallet
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

	if count > 0 {return true, err}

	// if not
	return true, nil
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
		"INSERT INTO products (product, product_code, description, cost, selling, cost_pallet ,units_ctn, ctn_pallet, units_pallet) VALUES (?,?,?,?,?,?,?,?,?)",
		p.Product, p.ProductCode, p.Description, p.Cost, p.Selling, p.CostPallet, p.UnitsCtn, p.CtnPallet, p.UnitsPallet,
		)
		if err != nil{
			log.Fatal(err)
		}
	 case true:
	 	fmt.Printf("Product %s already exists.\n", p.Product)
	 } 

}


