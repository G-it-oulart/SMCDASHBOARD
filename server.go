package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func handleRequests() {
	http.HandleFunc("/Filter", returnData)
	log.Fatal(http.ListenAndServe(":10000", nil))
}

func set_headers(w http.ResponseWriter) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.Header().Set("Access-Control-Allow-Methods", "GET, OPTIONS")
	w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
	w.Header().Set("Content-Type", "application/json")
}

func returnData(w http.ResponseWriter, r *http.Request) {
	set_headers(w)
	dataInit := r.URL.Query().Get("date_init")
	dataEnd := r.URL.Query().Get("date_end")
	materials := r.URL.Query().Get("materials")
	color := r.URL.Query().Get("color")
	responseData, _ := filt_dados_pesagens(dataInit, dataEnd, materials, color)
	sample_std_dev := std_dev(responseData)
	log.Printf("sample_std_dev: %+v\n", sample_std_dev)
	form_result := fmt.Sprintf("The standard deviation of the chosen frametime is: %.2f", sample_std_dev)
	if boca := json.NewEncoder(w).Encode(form_result); boca != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
