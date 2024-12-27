package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"
	"time"
)

func handleRequests() {
	http.HandleFunc("/Filter", returnData)
	http.HandleFunc("/dropdown", return_colors)
	http.HandleFunc("/Insert", insertData)
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
	Lab []string
	Std []float64
}


func return_colors(w http.ResponseWriter, r *http.Request) {
	set_headers(w)
	dropdown_colors:= conv_array_string(return_color_names(""))
	if boca := json.NewEncoder(w).Encode(dropdown_colors); boca != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
func returnData(w http.ResponseWriter, r *http.Request) {
	set_headers(w)
	var ucl_array []float64
	var lcl_array []float64
	var avg_array []float64
	var standards_array []float64
	var label_array []string
	var counter int
	var first_date string
	var last_date string
	var date_string string
	date_counter := 1
	dataInit:= r.URL.Query().Get("date_init")
	dataEnd:= r.URL.Query().Get("date_end")
	materials:= r.URL.Query().Get("materials")
	color:= r.URL.Query().Get("color")
	data_response, _:= filt_dados_pesagens(dataInit, dataEnd, materials, color)
	standards_response, _:= return_standards(materials, color)
	dates_response, _:= return_dates(dataInit, dataEnd, materials, color)
	data_ucl:= ucl(conv_array_float64(data_response))
	data_lcl:= lcl(conv_array_float64(data_response))
	data_avg := avg(conv_array_float64(data_response))

	for _,element := range format_json(conv_array_string(dates_response)) {
		if date_counter == 1 {
			first_date = element
		}
		if date_counter == 5{
			last_date = element
			date_counter = 0
			if last_date != first_date {
				date_string= fmt.Sprintf("%s - %s",first_date,last_date)
			} else {
				date_string = first_date
			}
			label_array = append(label_array,date_string)
		}
		date_counter+=1
	}
	for range five_avg(conv_array_float64((data_response))) {
		ucl_array = append(ucl_array, data_ucl)
		}
	for range five_avg(conv_array_float64((data_response))) {
		lcl_array = append(lcl_array, data_lcl)
		}
	for range five_avg(conv_array_float64((data_response))) {
		avg_array = append(avg_array, data_avg)
		}
		for range five_avg(conv_array_float64((data_response))) {
		standards_array = append(standards_array, conv_value_float64(standards_response))
		counter += 1
		}
	encode_pack := array_encoder{five_avg(conv_array_float64(data_response)), avg_array, ucl_array, lcl_array, label_array, standards_array}
	if boca := json.NewEncoder(w).Encode(encode_pack); boca != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}

func insertData (w http.ResponseWriter, r *http.Request) {
	set_headers(w)
	linhas:=r.URL.Query().Get("linhas")
	massa,_:= strconv.ParseFloat(r.URL.Query().Get("massa"),64) 
	primer,_:= strconv.ParseFloat(r.URL.Query().Get("primer"),64)
	verniz,_:= strconv.ParseFloat(r.URL.Query().Get("verniz"),64)
	esmalte,_:= strconv.ParseFloat(r.URL.Query().Get("esmalte"),64)
	tingidor,_:= strconv.ParseFloat(r.URL.Query().Get("tingidor"),64)
	cor:=r.URL.Query().Get("color_list")
	var current_date string
	var date = time.Now()
	current_date = date.Format("2006-01-02")
	insert_into_db(linhas,cor,current_date,massa,primer,verniz,esmalte,tingidor)
}