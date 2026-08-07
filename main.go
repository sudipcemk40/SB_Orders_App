package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/circular", AddCircular)
	http.HandleFunc("/circulars", enableCORS(GetCirculars))
	http.HandleFunc("/search", enableCORS(SearchCircular))
	http.HandleFunc("/category", GetCircularsByCategoryHandler)
	http.HandleFunc("/subject", GetCircularsBySubjectHandler)
	http.HandleFunc("/pdf", OpenPDF)
	http.HandleFunc("/dashboard", enableCORS(DashboardHandler))
	fmt.Println("Server running on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
func enableCORS(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		if r.Method == http.MethodOptions {
			return
		}

		next(w, r)
	}
}
