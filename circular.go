package main

type Circular struct {
	ID           int    `json:"id"`
	CircularNo   string `json:"circularNo"`
	CircularDate string `json:"circularDate"`
	Year         int    `json:"year"`
	Category     string `json:"category"`
	Subject      string `json:"subject"`
	PDFName      string `json:"pdfName"`
	PDFURL       string `json:"pdfUrl"`
}
