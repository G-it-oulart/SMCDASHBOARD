package main

import (
	"encoding/json"
	"log"
	"net/http"
)

func handleRequests() {
	http.HandleFunc("/Filter", returnData)
	http.HandleFunc("/dropdown", return_colors)
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
	Lab []int
	Std []float64
}

var ucl_array []float64
var lcl_array []float64
var avg_array []float64
var standards_array []float64
var label_array []int
var counter int

func return_colors(w http.ResponseWriter, r *http.Request) {
	set_headers(w)
	dropdown_colors:= conv_array_string(return_color_names(""))
	if boca := json.NewEncoder(w).Encode(dropdown_colors); boca != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
	
func returnData(w http.ResponseWriter, r *http.Request) {
	set_headers(w)
	dataInit := r.URL.Query().Get("date_init")
	dataEnd := r.URL.Query().Get("date_end")
	materials := r.URL.Query().Get("materials")
	color := r.URL.Query().Get("color")
	data_response, _ := filt_dados_pesagens(dataInit, dataEnd, materials, color)
	standards_response, _ := return_standards(materials, color)
	data_ucl := ucl(conv_array_float64(data_response))
	data_lcl := lcl(conv_array_float64(data_response))
	data_avg := avg(conv_array_float64(data_response))

	if ucl_array !=nil {
		ucl_array = nil
		for range five_avg(conv_array_float64((data_response))) {
			ucl_array = append(ucl_array, data_ucl)
		}
	} else {
		for range five_avg(conv_array_float64((data_response))) {
			ucl_array = append(ucl_array, data_ucl)
		}
	}
	
	if lcl_array !=nil{
		lcl_array = nil
		for range five_avg(conv_array_float64((data_response))) {
			lcl_array = append(lcl_array, data_lcl)
		}
	} else {
		for range five_avg(conv_array_float64((data_response))) {
			lcl_array = append(lcl_array, data_lcl)
		}
		}
	
	if avg_array !=nil {
		avg_array = nil
		for range five_avg(conv_array_float64((data_response))) {
		avg_array = append(avg_array, data_avg)
		}
	} else {
		for range five_avg(conv_array_float64((data_response))) {
		avg_array = append(avg_array, data_avg)
	}
	}
	if label_array !=nil {
		label_array = nil
		counter = 0
		for range five_avg(conv_array_float64((data_response))) {
		label_array = append(label_array, counter)
		counter += 1
		}
	} else {
		for range five_avg(conv_array_float64((data_response))) {
		label_array = append(label_array, counter)
		counter += 1
		}
	}
	
	if standards_array != nil {
		standards_array = nil
		for range five_avg(conv_array_float64((data_response))) {
		standards_array = append(standards_array, conv_value_float64(standards_response))
		}
	} else {
		for range five_avg(conv_array_float64((data_response))) {
		standards_array = append(standards_array, conv_value_float64(standards_response))
		counter += 1
		}
	}
	
	encode_pack := array_encoder{five_avg(conv_array_float64(data_response)), avg_array, ucl_array, lcl_array, label_array, standards_array}
	if boca := json.NewEncoder(w).Encode(encode_pack); boca != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}