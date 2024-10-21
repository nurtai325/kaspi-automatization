package handlers

import (
	"html/template"
)

var templates *template.Template

func ParseViews() error {
	templ, err := template.ParseFiles(
		"templ/customers.html",
		"templ/add-customer.html",
		"templ/qrcode.html",
	)
	templates = templ
	return err
}
