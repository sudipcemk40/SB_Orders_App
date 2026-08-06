package main

import (
	"database/sql"
	"fmt"
)

func GetPDFName(db *sql.DB, no string) (string, error) {

	var pdfName string
	fmt.Println("Searching for:", no)
	err := db.QueryRow(
		"SELECT pdf_name FROM circulars WHERE circular_no=$1",
		no,
	).Scan(&pdfName)

	fmt.Println("Error:", err)
	fmt.Println("PDF Name:", pdfName)
}
