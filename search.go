package main

import "database/sql"

func GetCircularsByYear(db *sql.DB, year int) ([]Circular, error) {

	rows, err := db.Query(`
SELECT id,
       circular_no,
       circular_date,
       year,
       category,
       subject,
       pdf_name,
       pdf_url
FROM circulars
WHERE year = $1
ORDER BY circular_date DESC
`, year)

	if err != nil {
		return nil, err
	}
	defer rows.Close()

	circulars := make([]Circular, 0)

	for rows.Next() {
		var c Circular

		err := rows.Scan(
			&c.ID,
			&c.CircularNo,
			&c.CircularDate,
			&c.Year,
			&c.Category,
			&c.Subject,
			&c.PDFName,
			&c.PDFURL,
		)

		if err != nil {
			return nil, err
		}

		circulars = append(circulars, c)
	}

	return circulars, nil
}
func GetCircularByNumber(db *sql.DB, no string) (Circular, error) {

	var c Circular

	query := `
SELECT
id,
circular_no,
circular_date,
year,
category,
subject,
pdf_name,
pdf_url
FROM circulars
WHERE circular_no = $1
`

	err := db.QueryRow(query, no).Scan(
		&c.ID,
		&c.CircularNo,
		&c.CircularDate,
		&c.Year,
		&c.Category,
		&c.Subject,
		&c.PDFName,
		&c.PDFURL,
	)

	return c, err
}
func GetCircularsByCategory(db *sql.DB, category string) ([]Circular, error) {

	query := `
SELECT id,
       circular_no,
       circular_date,
       year,
       category,
       subject,
       pdf_name,
       pdf_url
FROM circulars
WHERE category = $1
ORDER BY circular_date DESC
`

	rows, err := db.Query(query, category)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var circulars []Circular

	for rows.Next() {

		var c Circular

		err := rows.Scan(
			&c.ID,
			&c.CircularNo,
			&c.CircularDate,
			&c.Year,
			&c.Category,
			&c.Subject,
			&c.PDFName,
			&c.PDFURL,
		)

		if err != nil {
			return nil, err
		}

		circulars = append(circulars, c)
	}

	return circulars, nil
}
func GetCircularsBySubject(db *sql.DB, keyword string) ([]Circular, error) {

	query := `
SELECT id,
       circular_no,
       circular_date,
       year,
       category,
       subject,
       pdf_name,
       pdf_url
FROM circulars
WHERE subject ILIKE '%' || $1 || '%'
ORDER BY circular_date DESC
`

	rows, err := db.Query(query, keyword)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var circulars []Circular

	for rows.Next() {
		var c Circular

		err := rows.Scan(
			&c.ID,
			&c.CircularNo,
			&c.CircularDate,
			&c.Year,
			&c.Category,
			&c.Subject,
			&c.PDFName,
			&c.PDFURL,
		)

		if err != nil {
			return nil, err
		}

		circulars = append(circulars, c)
	}

	return circulars, nil
}
