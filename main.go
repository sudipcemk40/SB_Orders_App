package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/circular", AddCircular)
	http.HandleFunc("/circulars", GetCirculars)
	http.HandleFunc("/search", SearchCircular)
	http.HandleFunc("/category", GetCircularsByCategoryHandler)
	http.HandleFunc("/subject", GetCircularsBySubjectHandler)
	http.HandleFunc("/pdf", OpenPDF)
	fmt.Println("Server running on http://localhost:8080")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err)
	}
}
