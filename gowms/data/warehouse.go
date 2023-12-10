package data

import (
	_ "github.com/mattn/go-sqlite3"
)

type WarehouseBin struct {
	Id   int
	RFID string
	// ShelfLocation [][]Shelf ??
	Contains Product
	Quanity  int
}

func (w WarehouseBin) InsertBin() {}

func (w *WarehouseBin) RecieveProductIntoBin(p *Product) {} // does this need * ?

func (w *WarehouseBin) AddProductQuantity(p Product, q int) {}

func (w *WarehouseBin) RemoveProduct(p Product, q int) {}
