package routes

import (
    "net/http"
    "html/template"

    "github.com/SmokierLemur51/gowms/data"
)

type PublicPageData struct {
    Page string
    Title string
    CompanyName string
    Warehouse string
    Content string
    CSS string
    Vendors []data.Vendor
    Products []data.Product
}

var CSS_URL string = "/static/css/main.css"

func (p PublicPageData) RenderHTMLTemplate(w http.ResponseWriter, data PublicPageData) {
    tmpl, err := template.ParseFiles("templates/" + p.Page)
    if err != nil {
        return
    }
    err = tmpl.Execute(w, data)
    if err != nil {
        return
    }
    // this prevents the superflous hanlder err 
    w.Header().Set("Content-Type", "text/html; charset=utf-8")
}