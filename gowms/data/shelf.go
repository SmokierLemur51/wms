package data

import (
	_ "github.com/mattn/go-sqlite3"
)

type Shelf struct {
	Id   int
	RFID string
	// ShelfLocation [][]Shelf ??
	Contains Product
	Quanity  int
}


func (s Shelf) InsertShelf() {}

func (s *Shelf) RecieveProductIntoShelf(p *Product) {} // does this need * ?

func (s *Shelf) AddProductQuantity(p Product, q int) {}

func (s *Shelf) RemoveProduct(p Product, q int) {}
