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
	data.PopulateVendors(db, data.RandomVendors())
	data.PopulateProducts(db, data.RandomProducts())
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
