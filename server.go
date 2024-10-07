package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func handleRequests() {
	http.HandleFunc("/Filter", returnDate)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func set_headers(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
}

func returnDate(w http.ResponseWriter, r *http.Request) {
	set_headers(w)
	dataInit := r.URL.Query().Get("date_init")
	dataEnd := r.URL.Query().Get("date_end")
	materials := r.URL.Query().Get("materials")
	responseData, _ := filt_dados_pesagens(dataInit, dataEnd, materials)
	log.Printf("Response data: %+v\n", responseData)
	if boca := json.NewEncoder(w).Encode(responseData); boca != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
