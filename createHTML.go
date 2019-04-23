package main

import (
	"bytes"
	"html/template"
	"log"
)

// CreateActHTML fill Act HTML template with data
func CreateActHTML(data Data) string {
	t, err := template.ParseFiles("templates/act.html")
	if err != nil {
		log.Fatal(err)
	}

	tpl := new(bytes.Buffer)
	if err := t.Execute(tpl, data); err != nil {
		log.Fatal(err)
	}
	str := tpl.String()
	return str
}

// CreateBillHTML fill Bill HTML template with data
func CreateBillHTML(data Data) string {
	t, err := template.ParseFiles("templates/bill.html")
	if err != nil {
		log.Fatal(err)
	}

	tpl := new(bytes.Buffer)
	if err := t.Execute(tpl, data); err != nil {
		log.Fatal(err)
	}
	str := tpl.String()
	return str
}
