package main

import (
	"encoding/json"
	"time"
) 	
type JSON_float64 struct {
		Value float64 `json:"material"`
	}
type JSON_string struct {
		Value string `json:"string_value"`
	}
	
func conv_array_float64(orig_list []byte) []float64 {
	var new_list []JSON_float64
	var new_float_list []float64
	json.Unmarshal(orig_list, &new_list)
	for _, element := range new_list {
		new_float_list=append(new_float_list,element.Value)
	}
	return new_float_list
}

func conv_value_float64(orig_list []byte) float64 {
	var value []JSON_float64
	var float_list []float64
	json.Unmarshal(orig_list,&value)
	for _, element := range value {
		float_list=append(float_list,element.Value)
	}
	return float_list[0]
}

func conv_array_string(orig_list []byte) []string {
	var new_list []JSON_string
	var new_string_list []string
	json.Unmarshal(orig_list, &new_list)
	for _, element := range new_list {
		new_string_list=append(new_string_list,element.Value)
	}
	return new_string_list
}

func format_json(orig_dates []string) []string {
	var formatted_dates []string
	for _,element:= range orig_dates{
	t,_ := time.Parse(time.RFC3339, element)
	dateWithoutHours := t.Format("2006-01-02")
	formatted_dates = append(formatted_dates,dateWithoutHours)
	}
	return formatted_dates
}