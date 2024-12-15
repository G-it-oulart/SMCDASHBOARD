package main

import (
	"reflect"

	"github.com/jackc/pgx/v5/pgtype"
)

func conv_array_float64(conv_list []any) []float64 {
	var new_list []float64
	var floatType = reflect.TypeOf(float64(0))
	for _, element := range conv_list {
		new_element := reflect.ValueOf(element)
		new_element = reflect.Indirect(new_element)
		conv_element := new_element.Convert(floatType)
		new_list = append(new_list, conv_element.Float())
	}
	return new_list
}

func conv_value_float64(orig_value any) float64 {
	var floatType = reflect.TypeOf(float64(0))
	new_value := reflect.ValueOf(orig_value)
	new_value = reflect.Indirect(new_value)
	conv_value := new_value.Convert(floatType)
	return conv_value.Float()
}

func convert_numeric (orig_list []pgtype.Numeric) []float64 {
	for _, element := range orig_list {
		if element.Valid {
			floatValue := element.Float64
		}
}
}