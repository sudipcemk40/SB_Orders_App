package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
)

func AddCircular(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Only POST method allowed", http.StatusMethodNotAllowed)
		return
	}
	var c Circular
	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	db := ConnectDB()
	defer db.Close()
	err = InsertCircular(db, c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(map[string]string{
		"status":  "success",
		"message": "Circular received successfully",
	})
}
func GetCircularsOld(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Only GET method allowed", http.StatusMethodNotAllowed)
		return
	}

	yearStr := r.URL.Query().Get("year")

	year, err := strconv.Atoi(yearStr)
	if err != nil {
		http.Error(w, "Invalid year", http.StatusBadRequest)
		return
	}

	db := ConnectDB()
	defer db.Close()

	circulars, err := GetCircularsByYear(db, year)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(circulars)
}
func SearchCircular(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Only GET method allowed", http.StatusMethodNotAllowed)
		return
	}

	no := r.URL.Query().Get("no")

	db := ConnectDB()
	defer db.Close()

	circular, err := GetCircularByNumber(db, no)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(circular)
}
func GetCircularsByCategoryHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Only GET method allowed", http.StatusMethodNotAllowed)
		return
	}

	category := r.URL.Query().Get("category")

	db := ConnectDB()
	defer db.Close()

	circulars, err := GetCircularsByCategory(db, category)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(circulars)
}
func GetCircularsBySubjectHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		http.Error(w, "Only GET method allowed", http.StatusMethodNotAllowed)
		return
	}

	keyword := r.URL.Query().Get("keyword")

	db := ConnectDB()
	defer db.Close()

	circulars, err := GetCircularsBySubject(db, keyword)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(circulars)
}
func OpenPDF(w http.ResponseWriter, r *http.Request) {
	fmt.Println("OpenPDF handler called")

	no := r.URL.Query().Get("no")
	if no == "" {
		http.Error(w, "Circular number required", http.StatusBadRequest)
		return
	}

	db := ConnectDB()
	defer db.Close()

	pdfName := GetPDFName(no)
	if pdfName == "" {
		http.Error(w, "Invalid circular number", http.StatusBadRequest)
		return
	}
	fmt.Println(pdfName)
	fmt.Println("./pdfs/" + pdfName)

	path := "D:/Sb Orders App/pdfs/" + pdfName
	fmt.Println(path)
	http.ServeFile(w, r, path)
}
func DashboardHandler(w http.ResponseWriter, r *http.Request) {
	db := ConnectDB()
	defer db.Close()
	dashboard, err := GetDashboard(db)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(dashboard)
}
func GetCirculars(w http.ResponseWriter, r *http.Request) {
	db := ConnectDB()
	defer db.Close()

	rows, err := db.Query(`
        SELECT circular_no, pdf_name
        FROM circulars
        ORDER BY circular_no DESC
    `)
	if err != nil {
		http.Error(w, err.Error(), 500)
		return
	}
	defer rows.Close()

	type Circular struct {
		CircularNo string `json:"circularNo"`
		PDFName    string `json:"pdfName"`
	}

	var circulars []Circular

	for rows.Next() {
		var c Circular
		rows.Scan(&c.CircularNo, &c.PDFName)
		circulars = append(circulars, c)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(circulars)
}
