package main

import (
	"math"
)

func avg(conv_list []float64) float64{
	var sum float64
	var count float64
	for _, element := range conv_list {
		sum += element
		count += 1
	}
	avg := sum / count
	return avg
}

func std_dev(conv_list []float64) float64 {
	dev_avg := avg(conv_list)
	var dev_sum float64
	var dev float64
	for _, element := range conv_list {
		dev = (element - dev_avg) * (element - dev_avg)
		dev_sum += (dev)
	}
	return math.Sqrt((dev_sum) / (float64(len(conv_list)) - 1))
}
func ucl(orig_list []float64) float64 {
	list_avg := avg(orig_list)
	ucl := (list_avg + (3 * std_dev(orig_list)))
	return ucl
}

func lcl(orig_list []float64) float64 {
	list_avg := avg(orig_list)
	lcl := (list_avg - (3 * std_dev(orig_list)))
	return lcl
}
func five_avg(sql_list []float64) []float64 {
	var list []float64
	var avg_list []float64
	counter := 0
	for _, element := range sql_list {
		avg_list = append(avg_list, element)
		if counter == 4 {
			list = append(list, avg(avg_list))
		}
		counter += 1
		if counter == 5 {
			counter = 0
			avg_list = nil
		}
	}
	return list
}
