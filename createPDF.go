package main

import (
	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"log"
	"strconv"
	"strings"
	"time"
)

func CreateActPDF(data Data, input string) {
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		log.Fatal(err)
	}

	pdfg.Dpi.Set(600)
	pdfg.PageSize.Set(wkhtmltopdf.PageSizeA4)

	pdfg.AddPage(wkhtmltopdf.NewPageReader(strings.NewReader(input)))
	err = pdfg.Create()
	if err != nil {
		log.Fatal(err)
	}

	fileName := "./Act #" + strconv.Itoa(data.Act) + " " + data.SignOn.Format(time.RFC822) + ".pdf"
	err = pdfg.WriteFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
}
