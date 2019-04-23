package main

import (
	"github.com/SebastiaanKlippert/go-wkhtmltopdf"
	"log"
	"strings"
)

// CreatePDF
func CreatePDF(input, fileName string) {
	pdfg, err := wkhtmltopdf.NewPDFGenerator()
	if err != nil {
		log.Fatal(err)
	}

	pdfg.Dpi.Set(75)
	pdfg.PageSize.Set(wkhtmltopdf.PageSizeA4)

	pdfg.AddPage(wkhtmltopdf.NewPageReader(strings.NewReader(input)))
	err = pdfg.Create()
	if err != nil {
		log.Fatal(err)
	}

	err = pdfg.WriteFile(fileName)
	if err != nil {
		log.Fatal(err)
	}
}
