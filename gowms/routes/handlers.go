package routes

import (
    "net/http"
    // "log"
    // "strconv"
    // "github.com/SmokierLemur51/gowms/data"
)


func IndexHandler(w http.ResponseWriter, r *http.Request) error {
    page := PublicPageData{
        Page: "index.html",
        Title: "WMS",
        CompanyName: "Precision Parts Distributors",
        Warehouse: "Lous",
        Content: "Sample Content",
        CSS: CSS_URL,
    }

    page.RenderHTMLTemplate(w, page)
    return nil
}

func InventoryHandler(w http.ResponseWriter, r *http.Request) error {
    page := PublicPageData{
        Page: "inventory.html",
        Title: "WMS - Inventory",
        CompanyName: "Precision Parts Distributors",
        Warehouse: "Lous",
        Content: "Sample Content",
        CSS: CSS_URL,
    }

    page.RenderHTMLTemplate(w, page)
    return nil
}

// non ajax version, simply http.Post
func CreateProductHandler(w http.ResponseWriter, r *http.Request) error {
    err := r.ParseForm()
    if err != nil {
        http.Error(w, err.Error(), http.StatusInternalServerError)
        return err
    }
    // converting units per ctn to int
    uCtn := r.FormValue("unitsCtn")
    if uc, err := strconv.Atoi(uCtn); err != nil {
        log.Printf("Error converting %s to integer.", uCtn)
        return err
    }
    // converting ctns per pallet to int
    ctnPal := r.FormValue("ctnPallet")
    if ctp, err := strconv.Atoi(ctnPal); err != nil {
        log.Printf("Error converting %s to integer.", ctnPal)
        return err        
    }
    // converting cost per pallet to float64
    costPal := r.FormValue("costPallet")
    if csp, err := strconv.ParseFloat(costPal, 64); err != nil {
        log.Printf("Error converting %s to float64.", costPal)
        return err
    }

    p := data.Product {
        VendorId: data.FindVendorId(r.FormValue("vendor")),
        Status: data.FindStatusId(r.FormValue("status")),
        Product: r.FormValue("product"),
        ProductCode: r.FormValue("productCode"),
        Description: r.FormValue("description"),
        UnitsCtn: uc,
        CtnPallet: ctp,
        CostPallet: csp,
    }

    if p.Product == "" || p.ProductCode == "" || p.CostPallet == 0.0 || p.CtnPallet == 0 || p.UnitsCtn == 0 {
        http.Error(w, "Missing form fields", http.StatusBadRequest)
        return nil
    }


    http.Redirect(w, r, "/inventory", http.StatusSeeOther)
}

// testing theory here 
func TestHandler(w http.ResponseWriter, r *http.Request) error {} 