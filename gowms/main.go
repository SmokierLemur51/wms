package main

import (
	"log"
    "net/http"
    "database/sql"

    "github.com/SmokierLemur51/gowms/routes"
    "github.com/SmokierLemur51/gowms/data"

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

    yarn := data.Product{
        Product: "Yarn", ProductCode: "aj3if89ajfdklaf", Description: "Some high quality yarn",
        Cost: 3.99, Selling: (3.99/.80), UnitsCtn: 20, CtnPallets: 40,
    }
    hydro := data.Product{
        Product: "Hydroflask", ProductCode: "jfakj3184", Description: "Drink water like a tiktok girl!",
        Cost: 35.00, Selling: (35.00/.80), UnitsCTN: 10, CtnPallets: 20,
    }
    yarn.InsertProduct(db)
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
