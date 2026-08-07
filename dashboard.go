package main

import "database/sql"

type Dashboard struct {
	TotalCirculars int    `json:"totalCirculars"`
	TotalYears     int    `json:"totalYears"`
	LatestCircular string `json:"latestCircular"`
}

func GetDashboard(db *sql.DB) (Dashboard, error) {
	var d Dashboard
	err := db.QueryRow(`
		SELECT 
		COUNT(*),
		COUNT (DISTINCT year),
		MAX(circular_no) 
		from circulars
		`).Scan(
		&d.TotalCirculars,
		&d.TotalYears,
		&d.LatestCircular,
	)
	return d, err
}
