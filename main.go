package main

import (
	"gopkg.in/alecthomas/kingpin.v2"
	"log"
	"os"
	"time"
)

var (
	signOn = kingpin.Flag("signOn", "SignOn date").String()
)

func init() {
	if _, err := os.Stat("./reports"); os.IsNotExist(err) {
		os.Mkdir("./reports", os.ModeDir)
	}
}

func main() {
	kingpin.Parse()

	data := GetData("data.json")
	data.setSignOn(*signOn)

	actHTML := CreateActHTML(data)
	billHTML := CreateBillHTML(data)

	path := "./reports/" + data.SignOn.Format("2006-01-02")
	if _, err := os.Stat(path); os.IsNotExist(err) {
		os.Mkdir(path, os.ModePerm)
	}
	actName := path + "/Act #" + data.Act + " " + data.SignOn.Format(time.RFC822) + ".pdf"
	billName := path + "/Bill #" + data.Act + " " + data.SignOn.Format(time.RFC822) + ".pdf"

	CreatePDF(actHTML, actName)
	CreatePDF(billHTML, billName)
	log.Println("Done!!!")
}
