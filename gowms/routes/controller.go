package routes

import (
	"net/http"
	"log"
	"database/sql"
	"encoding/json"
	"time"
	_ "github.com/mattn/go-sqlite3"
	"github.com/go-chi/chi/v5"
	"github.com/golang-jwt/jwt"
)

var COMPANY_NAME string = "Logans Building Products"

type Controller struct {
	db *sql.DB
	// logger here
}


func (c Controller) ConnectDatabase(database, connection string) {
	var err error
	if c.db, err = sql.Open(database, connection); err != nil {
		log.Fatal(err)
	}
	log.Printf("Connected to %s database file %s", database, connection)
}

func (c Controller) ConfigureRoutes(r chi.Router) {
	r.Method(http.MethodGet, "/login", c.LoginHandler())
	r.Method(http.MethodGet, "/", c.IndexHandler())
	r.Method(http.MethodGet, "/inventory", c.InventoryHandler())
	r.Method(http.MethodGet, "/products", c.ProductsHandler())

	r.Method(http.MethodPost, "/create-product", c.CreateProduct())
	r.Method(http.MethodPost, "/process-login", c.ProcessLogin())
}

func (c Controller) LoginHandler() http.HandlerFunc {
	// authentication switch, case true/false 
	// if !auth redirect with c.NotAuthenticated
	return func(w http.ResponseWriter, r *http.Request) {
		p := PublicPageData{Page: "login.html",CompanyName: COMPANY_NAME, CSS: CSS_URL,}
		p.RenderHTMLTemplate(w)
	}
}

func (c Controller) AuthenticateUser() bool {
	var authenticated bool

	switch authenticated {
	case true:
		return true
	case false: 
		return false 
	}
	return false
}

func (c Controller) IndexHandler() http.HandlerFunc {
	// authentication switch, case true/false 
	// if !auth redirect with c.NotAuthenticated
	return func(w http.ResponseWriter, r *http.Request) {
		p := PublicPageData{
			Page: "index.html",Title: "WMS",CompanyName: COMPANY_NAME,
    		Warehouse: "Lous",Content: "Sample Content",CSS: CSS_URL,
		}
		p.RenderHTMLTemplate(w)
	}
}

func (c Controller) InventoryHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		p := PublicPageData{
			Page: "inventory.html",Title: "WMS",CompanyName: COMPANY_NAME,
    		Warehouse: "Lous",Content: "Sample Content",CSS: CSS_URL,
		}
		p.RenderHTMLTemplate(w)
	}
}

func (c Controller) ProductsHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		p := PublicPageData{
			Page: "products.html",Title: "WMS",CompanyName: COMPANY_NAME,
    		Warehouse: "Lous",Content: "Sample Content",CSS: CSS_URL,
		}
		p.RenderHTMLTemplate(w)
	}
}

/*
Post methods gang 
*/

var jwtKey = []byte("secret_key")

var users  = map[string]string{
	"user1": "passone",
	"user2": "passtwo",
}

type Credentials struct {
	Username string `json:"username"`
	Password string `json:"password"` 
}

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims 
}

func (c Controller) ProcessLogin() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		var credentials Credentials
		err := json.NewDecoder(r.Body).Decode(&credentials)
		if err != nil {
			w.WriteHeader(http.StatusBadRequest)
		}

		expectedPassword, ok := users[credentials.Username]

		if !ok || expectedPassword != credentials.Password {
			w.WriteHeader(http.StatusUnauthorized)
		}

		expirationTime := time.Now().Add(time.Minute * 5)

		claims := &Claims{
			Username: credentials.Username,
			StandardClaims: jwt.StandardClaims{
				ExpiresAt: expirationTime,
			},
		}

		token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
		tokenstring, err := token.SignString(jwtKey)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
		}

		http.SetCookie(w, 
			&http.Cookie{
				Name: "token",
				Value: tokenstring,
				Expires: expirationTime,
			})


		http.Redirect(w, r, "/", http.StatusFound)
	}
}

func (c Controller) CreateProduct() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			log.Fatal(err)
		}

		// form actions here 
	}
}


