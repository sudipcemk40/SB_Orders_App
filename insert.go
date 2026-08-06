package main

import "database/sql"

func InsertCircular(db *sql.DB, c Circular) error {
	query := `
	INSERT INTO circulars
	(circular_no,circular_date,year,category,subject,pdf_name,pdf_url)
	VALUES($1,$2,$3,$4,$5,$6,$7)
	`
	_, err := db.Exec(
		query,
		c.CircularNo,
		c.CircularDate,
		c.Year,
		c.Category,
		c.Subject,
		c.PDFName,
		c.PDFURL,
	)
	return err
}
