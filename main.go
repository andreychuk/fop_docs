package main

import (
	"gopkg.in/alecthomas/kingpin.v2"
)

var (
	dataFile = kingpin.Flag("data", "JSON data file").Required().String()
	signOn   = kingpin.Flag("signOn", "SignOn date").String()
)

func main() {
	kingpin.Parse()
	data := GetData(*dataFile)
	data.setSignOn(*signOn)
	actHTML := CreateActHTML(data)
	CreateActPDF(data, actHTML)
}
