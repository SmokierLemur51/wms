package routes

import (
    "net/http"
    "log"
    "database/sql"
    "github.com/SmokierLemur51/gowms/data"
    _ "github.com/mattn/go-sqlite3"
)

var COMPANY_NAME string = "Siding Distribution"

func IndexHandler(w http.ResponseWriter, r *http.Request) error {
    // this should have an introduction of the project
    page := PublicPageData{
        Page: "index.html",
        Title: "WMS",
        CompanyName: COMPANY_NAME,
        Warehouse: "Lous",
        Content: "Sample Content",
        CSS: CSS_URL,
    }

    page.RenderHTMLTemplate(w)
    return nil
}

func InventoryHandler(w http.ResponseWriter, r *http.Request) error {
    page := PublicPageData{
        Page: "inventory.html",
        Title: "WMS - Inventory",
        CompanyName: COMPANY_NAME,
        Warehouse: "Lous",
        Content: "Sample Content",
        CSS: CSS_URL,
    }

    page.RenderHTMLTemplate(w)
    return nil
}


func ProductsHandler(w http.ResponseWriter, r *http.Request) error {
    db, err := sql.Open("sqlite3", data.DB_FILE)
    if err != nil {
        return err
    }
    page := PublicPageData{
        Page: "products.html",
        Title: "WMS - Products",
        CompanyName: COMPANY_NAME,
        Warehouse: "Lous",
        Content: "Sample Content",
        CSS: CSS_URL,
        Products: data.LoadAllStockProducts(db), 
    }

    page.RenderHTMLTemplate(w)
    return nil
}


// the form is not complete
// non ajax version, simply http.Post
func CreateProductHandler(w http.ResponseWriter, r *http.Request) error {
    err := r.ParseForm()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return err
    }
    // converting units per ctn to int
    uCtn, err := ConvertStrInt(r.FormValue("unitsCtn"))
    if err != nil {
        return err
    }

    // converting ctns per pallet to int
    ctnPal, err := ConvertStrInt(r.FormValue("ctnPallet"))
    if err != nil {
        return err
    }

    // converting cost per pallet to float64
    costPal, err := ConvertStrFloat64(r.FormValue("costPallet"))
    if err != nil {
        return err
    }
    db, err := sql.Open("sqlite3", "testing.db")
    if err != nil {
        log.Fatal(err)
        return err
    }
    p := data.Product {
        VendorId: data.FindDatabaseID(db, "vendors", "vendor", r.FormValue("vendor")),
        Status: data.FindDatabaseID(db, "stock_status", "disposition" ,r.FormValue("status")),
        Product: r.FormValue("product"),
        ProductCode: r.FormValue("productCode"),
        Description: r.FormValue("description"),
        UnitsCtn: uCtn,
        CtnPallet: ctnPal,
        CostPallet: costPal,
    }

    if p.Product == "" || p.ProductCode == "" || p.CostPallet == 0.0 || p.CtnPallet == 0 || p.UnitsCtn == 0 {
        http.Error(w, "Missing form fields", http.StatusBadRequest)
        return nil
    } else {
        db, err := sql.Open("sqlite3", "testing.db")
        if err != nil {
            return err
        }
        p.InsertProduct(db)
    }


    http.Redirect(w, r, "/inventory", http.StatusSeeOther)
    return nil
}
