package main

import (
	// "database/sql"
	"log"
	// "fmt"
	"net/http"

	// "github.com/SmokierLemur51/gowms/data"
	"github.com/SmokierLemur51/gowms/routes"

	// _ "github.com/mattn/go-sqlite3"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	// "github.com/go-chi/jwtauth/v5"
)

// var db *sql.DB
// var tokenAuth *jwtauth.JWTAuth


// func init() {
// 	var err error
// 	db, err = sql.Open("sqlite3", "testing.db")
// 	if err != nil {
// 		log.Fatal(err)
// 	}
// 	defer db.Close()
// 	// jwtauth 
// 	tokenAuth = jwtauth.New("HS256", []byte("secret"), nil)
	
// 	_, tokenString, _ := tokenAuth.Encode(map[string]interface{}{"user_id": 123})
// 	fmt.Printf("DEBUG: a sample jwt is %v\n\n", tokenString)

	// populate categories
	// data.PopulateDbCategories(db, data.CreateSidingCategorySlice())

	// populate vendors
	// data.PopulateVendors(db, data.CreateSidingVendorSlice())

	// populate products
	
// }

func main() {

	var PORT string = ":5000"
	r := chi.NewRouter()
	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("./static"))))

	// r.Group(func(r chi.Router) {
	// 	// seek, verify and validate JWT tokens
	// 	r.Use(jwtauth.Verifier(tokenAuth))

	// 	// Handle valid / invalid tokens. In this example, we use
	// 	// the provided authenticator middleware, but you can write your
	// 	// own very easily, look at the Authenticator method in jwtauth.go
	// 	// and tweak it, its not scary.
	// 	r.Use(jwtauth.Authenticator(tokenAuth))
	// 	r.Method(http.MethodGet, "/testing", routes.Handler(routes.AuthTestHandler))
	// })

	// routes.ConfigureRoutes(r)

	c := routes.Controller{}
	c.ConnectDatabase("sqlite3", "testing.db")
	c.ConfigureRoutes(r)

	log.Println("Starting server on port ", PORT)
	http.ListenAndServe(PORT, r)
}
