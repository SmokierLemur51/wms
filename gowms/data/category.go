package data

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/mattn/go-sqlite3"
)

/*
	need to be able to

	update
	delete
	load all into slice
*/

type Category struct {
	Id          int    `db:"id"`
	Category    string `db:"category"`
	Description string `db:"c_description"`
}

func CheckExistingCategory(db *sql.DB, category string) (bool, error) {
	// returns true if it exists
	var count int
	rows, err := db.Query("SELECT COUNT(*) FROM categories WHERE category = ?", category)
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

func (c Category) InsertCategory(db *sql.DB) {
	var execute bool
	var err error
	execute, err = CheckExistingCategory(db, c.Category)
	if err != nil {
		log.Println(err)
		return
	}

	// remember, the check existing returns true if the category already exists, so it skips
	switch execute {
	case false:
		_, err := db.Exec(
			"INSERT INTO categories (category, c_description) VALUES (?,?)",
			c.Category, c.Description,
		)
		if err != nil {
			log.Fatal(err)
		}
	case true:
		fmt.Printf("Category <%s> already exists.\n", c.Category)
	}
}

func (c Category) UpdateCategory(db *sql.DB, ID int) {
	var exists bool
	var err error
	exists, err = CheckExistingCategory(db, c.Category)
	if err != nil {
		log.Fatal(err)
		return
	}

	// remember this executes only if it does exist of course
	switch exists {
	case false:
		fmt.Printf("Category <%s> does not exist.", c.Category)
	case true:
		result, err := db.Exec(
			"UPDATE categories SET category = ?, c_description = ? WHERE id = ?",
			c.Category, c.Description, ID,
		)
		if err != nil {
			log.Fatal(err)
		}
		rowsAffected, err := result.RowsAffected()
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("Updated %d rows.\n", rowsAffected)
	}
}

func LoadAllCategories(db *sql.DB) []Category {
	rows, err := db.Query("SELECT id, category, c_description FROM categories")
	if err != nil {
		log.Fatal(err)
	}

	var categories []Category
	for rows.Next() {
		var c Category
		if err := rows.Scan(&c.Id, &c.Category, &c.Description); err != nil {
			log.Fatal(err)
		}
		// fmt.Printf("\tCategory: %s\n\tDescription: %s\n", c.Category, c.Description)
		categories = append(categories, c)
	}
	return categories
}

func CreateSidingCategorySlice() []Category {
	return []Category{
		{Category:"Vinyl Siding", Description:"Exterior vinyl siding products."},
		{Category:"Vinyl Accessories",Description:"Accessories for vinyl siding."},
		{Category:"Mounting Blocks",Description:"Mounting blocks such as light blocks, split/electrical blocks, or water spickets."},
		{Category:"Composite Siding", Description:"Exterior composite siding products."},
		{Category:"Fiber Cement Siding",Description:"Fiber cement related products."},
		{Category:"PVC Moulding",Description:"PVC trim, moulding, and accessories."},
		{Category:"Decking",Description:"Decking products for wood based deck materials."},
		{Category:"Composite Decking",Description:"Composite decking products. PVC/plastic-film and recycled wood materials."},
		{Category:"Gutter Coil",Description:"Aluminum gutter coil. For 5\" and 6\" spools.."},
		{Category:"Gutter Accessories",Description:"Aluminum gutter accessories."},
		{Category:"Trim Coil",Description:"Aluminum roll for exterior finishing/wrapping."},
		{Category:"Windows",Description:"Windows and window accessories."},
	}
}

 
/*
The rest of this page are for a firearms distribution network
*/
func CreateArmsDealerCategorySlice() []Category {
	return []Category{
		{Category: "Handgun", Description: "Compact and versatile, our selection of handguns includes a variety of models suitable for personal protection, target shooting, and concealed carry. Available in different calibers, each handgun is designed for reliability and ease of use."},
		{Category: "Rifles", Description: "Our rifle collection features precision-engineered firearms suitable for various applications, from hunting to long-range shooting. Explore our diverse range of rifles, each crafted with attention to accuracy, durability, and performance."},
		{Category: "Shotguns", Description: "Whether for sport shooting or home defense, our shotguns deliver power and versatility. With various styles and configurations available, our shotguns are designed to meet the needs of both novice and experienced shooters."},
		{Category: "Accessories", Description: "Enhance your shooting experience with our wide array of firearm accessories. From high-capacity magazines to advanced optics and grips, our accessories are crafted to improve accuracy, customization, and overall firearm performance."},
		{Category: "Ammunition", Description: "Enhance your shooting experience with our wide array of firearm accessories. From high-capacity magazines to advanced optics and grips, our accessories are crafted to improve accuracy, customization, and overall firearm performance."},
		{Category: "Firearm Parts", Description: "Customize and maintain your firearms with our selection of quality parts. From precision barrels to enhanced triggers, our firearm parts are designed to optimize performance, accuracy, and overall functionality."},
		{Category: "Safety Equipment", Description: "Prioritize safety with our comprehensive range of safety equipment. Our ear and eye protection, shooting range accessories, and storage solutions are crafted to ensure a secure and responsible shooting experience."},
		{Category: "Apparrel and Gear", Description: "Stay comfortable and prepared with our collection of shooting sports apparel and tactical gear. From durable clothing to range bags and accessories, our gear is designed for both functionality and style."},
		{Category: "Training and Eduational Materials", Description: "Elevate your skills with our training and educational materials. Explore instructional videos, manuals, and targets to enhance your understanding of firearm safety, maintenance, and marksmanship."},
		{Category: "Security Systems", Description: "Safeguard your firearms and property with our selection of security systems. Our surveillance cameras, alarm systems, and access control solutions provide peace of mind and protect your valuable assets."},
	}
}

func PopulateDbCategories(db *sql.DB, c []Category) {
	for _, category := range c {
		category.InsertCategory(db)
	}
}
