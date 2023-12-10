package routes

import (
    "net/http"
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