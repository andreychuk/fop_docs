package main

import (
	"bytes"
	"html/template"
	"log"
)

func CreateActHTML(data Data) string {
	//t := template.New("templates/act.html")
	t, err := template.ParseFiles("templates/act.html")
	if err != nil {
		log.Fatal("wwwww ", err)
	}

	tpl := new(bytes.Buffer)
	if err := t.Execute(tpl, data); err != nil {
		log.Fatal("qqqqq ", err)
	}
	str := tpl.String()
	return str
}
