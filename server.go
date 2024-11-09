package main

import (
	"encoding/json"
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

type array_encoder struct {
	Val []float64
	Avg []float64
	Ucl []float64
	Lcl []float64
}

func returnData(w http.ResponseWriter, r *http.Request) {
	set_headers(w)
	dataInit := r.URL.Query().Get("date_init")
	dataEnd := r.URL.Query().Get("date_end")
	materials := r.URL.Query().Get("materials")
	color := r.URL.Query().Get("color")
	responseData, _ := filt_dados_pesagens(dataInit, dataEnd, materials, color)
	data_ucl := ucl(responseData)
	data_lcl := lcl(responseData)
	data_avg := avg(responseData)
	var ucl_array []float64
	var lcl_array []float64
	var avg_array []float64
	for range responseData {
		ucl_array = append(ucl_array, data_ucl)
	}
	for range responseData {
		lcl_array = append(lcl_array, data_lcl)
	}
	for range responseData {
		avg_array = append(avg_array, data_avg)
	}
	encode_pack := array_encoder{responseData, avg_array, ucl_array, lcl_array}
	if boca := json.NewEncoder(w).Encode(encode_pack); boca != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
