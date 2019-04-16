package main

import (
	"gopkg.in/alecthomas/kingpin.v2"
	"log"
	"os"
	"strconv"
	"time"
)

var (
	dataFile = kingpin.Flag("data", "JSON data file").Required().String()
	signOn   = kingpin.Flag("signOn", "SignOn date").String()
)

func init() {
	if _, err := os.Stat("./reports"); os.IsNotExist(err) {
		os.Mkdir("./reports", os.ModeDir)
	}
}

func main() {
	kingpin.Parse()

	data := GetData(*dataFile)
	data.setSignOn(*signOn)

	actHTML := CreateActHTML(data)
	billHTML := CreateBillHTML(data)

	path := "./reports/" + data.SignOn.Format("2006-01-02")
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, os.ModePerm)
	}
	actName := path + "/Act #" + strconv.Itoa(data.Act) + " " + data.SignOn.Format(time.RFC822) + ".pdf"
	billName := path + "/Bill #" + strconv.Itoa(data.Act) + " " + data.SignOn.Format(time.RFC822) + ".pdf"

	CreatePDF(actHTML, actName)
	CreatePDF(billHTML, billName)
	log.Println("Done!!!")
}
