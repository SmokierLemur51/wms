package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/SmokierLemur51/gowms/data"
	"github.com/SmokierLemur51/gowms/routes"

	_ "github.com/mattn/go-sqlite3"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

var db *sql.DB

func init() {
	var err error
	db, err = sql.Open("sqlite3", "testing.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	pops := data.Product{
		Product: "Pops", ProductCode: "874", Description: "Kevin Malone - The Office",
		UnitsCtn: 12, CtnPallet: 48, CostPallet: 4500.00,
	}
	hydro := data.Product{
		Product: "HydroFlask", ProductCode: "jhafdkljklfjdlk", Description: "Drink water, aesthetically.",
		UnitsCtn: 6, CtnPallet: 45, CostPallet: 6200.00,
	}

	pops.InsertProduct(db)
	hydro.InsertProduct(db)

}

func main() {
	var PORT string = ":5000"
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	routes.ConfigureRoutes(r)
	log.Println("Starting server on port ", PORT)
	http.ListenAndServe(PORT, r)
}
