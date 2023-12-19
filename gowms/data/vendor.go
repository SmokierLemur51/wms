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

// vendor methods
func (v Vendor) InsertVendor(db *sql.DB, ) {
	var execute bool
	var err error
	execute, err = CheckExistingVendor(db, v.Vendor)
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

func (v *Vendor) UpdateVendor(db *sql.DB) {}


func (v Vendor) LoadAllStockProducts(db *sql.DB) {
	// whole func needs some review
	query := `
		SELECT id, vendor_id, status_id, product, product_code, p_description, units_ctn, ctn_pallet, units_pallet,
		cost_pallet, selling_pallet, cost_ctn, selling_ctn, cost_unit, selling_unit FROM proudcts
		WHERE status_id = 1 AND vendor_id = ?
	`
	rows, err := db.Query(query, v.Id)
	if err != nil {
		log.Fatal(err)
	} 
	for rows.Next() {
		var p Product
		if err := rows.Scan(
			&p.Id, &p.VendorId, &p.Status, &p.Product, &p.ProductCode, &p.Description, &p.UnitsCtn,
			&p.CtnPallet, &p.UnitsPallet, &p.CostPallet, &p.SellingPallet, &p.CostCtn, &p.SellingCtn,
			&p.CostUnit, &p.SellingUnit,
		)
		err != nil {
			log.Fatal(err)
		}
		v.StockProducts = append(v.StockProducts, p) 
	}
}

func (v Vendor) LoadSpecialOrderProducts(db *sql.DB) {
	// whole func needs some review
	query := `
		SELECT id, vendor_id, status_id, product, product_code, p_description, units_ctn, ctn_pallet, units_pallet,
		cost_pallet, selling_pallet, cost_ctn, selling_ctn, cost_unit, selling_unit FROM proudcts
		WHERE status_id = 2 AND vendor_id = ?
	`
	rows, err := db.Query(query, v.Id)
	if err != nil {
		log.Fatal(err)
	} 
	for rows.Next() {
		var p Product
		if err := rows.Scan(
			&p.Id, &p.VendorId, &p.Status, &p.Product, &p.ProductCode, &p.Description, &p.UnitsCtn,
			&p.CtnPallet, &p.UnitsPallet, &p.CostPallet, &p.SellingPallet, &p.CostCtn, &p.SellingCtn,
			&p.CostUnit, &p.SellingUnit,
		)
		err != nil {
			log.Fatal(err)
		}
		v.StockProducts = append(v.StockProducts, p) 
	}
}

// queries
func CheckExistingVendor(db *sql.DB, vendor string) (bool, error) {
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

func LoadVendorByName(db *sql.DB, vendor string) Vendor {
	rows, err := db.Query(
		"SELECT id, vendor, address_street, address_street_2, address_city, address_state, address_zip FROM vendors WHERE vendor = ?",
		vendor,
	)
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()
	var v Vendor
	for rows.Next() {
		err := rows.Scan(&v.Id, &v.Vendor, &v.Street1, &v.Street2, &v.City, &v.State, &v.Zip)
		if err != nil {
			log.Fatal(err)
		}
		log.Printf("\nID: %d\tVendor: %s", v.Id, v.Vendor)
	}

	if err := rows.Err(); err != nil {
		log.Fatal(err)
	}
	return v
}

// populate original values
func CreateSidingVendorSlice() []Vendor {
	return []Vendor{
		{Vendor:"Mastic",Street1:"5020 Weston Parkway",Street2:"Suite 400",City:"Cary",State:"NC",Zip:"27513"},
		{Vendor:"Royal",Street1:"91 Royal Group Crescent",Street2:"",City:"Woodbridge",State:"ON", Zip:"L4H 1X9"},
		{Vendor:"US Lumber",Street1:"6002 Sunnyside Rd",Street2:"",City:"Indianapolis",State:"IN", Zip:"46"},
		{Vendor:"Lumbermans",Street1:"4433 Stafford Ave SW",Street2:"",City:"Grand Rapids",State:"MI", Zip:"49548"},
		{Vendor:"Palmer-Donavin",Street1:"3620 Langley Dr",Street2:"",City:"Hebron",State:"KY", Zip:"41048"},
		{Vendor:"Wincore",Street1:"250 Staunton Turnpike",Street2:"",City:"Parkersburg",State:"WV", Zip:"26104"},		
		{Vendor:"Atrium",Street1:"Somewhere I cant find",Street2:"",City:"Nowhere",State:"<3", Zip:"11111"},		
	}
}

func PopulateVendors(db *sql.DB, v []Vendor) {
	for _, vendor := range v {
		vendor.InsertVendor(db)
	}
}