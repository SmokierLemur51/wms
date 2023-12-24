package routes

import (
	"net/http"
	"log"
	"fmt"
	"database/sql"
	// "encoding/json"
	"time"
	_ "github.com/mattn/go-sqlite3"
	"github.com/go-chi/chi/v5"
	"github.com/golang-jwt/jwt/v5"
)

// jwt docs
var (
	key []byte 
	t *jwt.Token
	s string
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

	// testing
	r.Method(http.MethodGet, "/testing", c.TestingHandler())
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

type CustomClaims struct {
	Username string
	Password string
	jwt.RegisteredClaims
}

func createToken(username, password string) (string, error) {
	claims := CustomClaims{
		Username: username,
		Password: password,
		RegisteredClaims: jwt.RegisteredClaims{
			ExpiresAt: jwt.NewNumericDate(time.Now().Add(24 * time.Hour)),
			IssuedAt: jwt.NewNumericDate(time.Now()),
			NotBefore: jwt.NewNumericDate(time.Now()),
			Subject: username,
		},
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	signedToken, err := token.SignedString(key)
	if err != nil {
		return "", err
	}
	return signedToken, nil
}

func (c Controller) ProcessLogin() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := r.ParseForm(); err != nil {
			log.Fatal(err)
		}

		token, err := createToken(r.FormValue("username"), r.FormValue("password"))
		if err != nil {
			log.Fatal(err)
		}

		// w.Write([]byte(token))
		http.SetCookie(w, &http.Cookie{
			Name: "jwt_token",
			Value: token,
			Expires: time.Now().Add(time.Hour * 24),
			HttpOnly: true,
		})

		http.Redirect(w, r, "/testing", http.StatusFound)
	}
}

func (c Controller) TestingHandler() http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		cookie, err := r.Cookie("jwt_token")
		if err != nil {
			http.Error(w, "No auth", http.StatusUnauthorized)
		}

		tokenString := cookie.Value

		// parse and validate token
		claims := &CustomClaims{}
		token, err := jwt.ParseWithClaims(tokenString, claims, func(token *jwt.Token) (interface{}, error) {
			return key, nil
		})

		if err != nil {
			http.Error(w, "Failed to parse token.", http.StatusUnauthorized)
		}

		// if !token.Valid {
		// 	http.Error(w, "Invalid token", http.StatusUnauthorized)
		// }

		if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
			fmt.Printf("%v %v", claims.Username, claims.RegisteredClaims.Issuer)
		} else {
			fmt.Println(err)
		}

		username := claims.Username
		password := claims.Password

		w.Write([]byte(fmt.Sprintf("Protected area accessed by user: %s\t password: %s", username, password)))
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
