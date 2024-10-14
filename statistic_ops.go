package main

import (
	"math"
)

func avg(conv_list []float64) float64 {
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
	ucl := (list_avg + (3 * std_dev(orig_list))) / math.Sqrt(float64(len(orig_list)))
	return ucl
}
func lcl(orig_list []float64) float64 {
	list_avg := avg(orig_list)
	lcl := (list_avg - (3 * std_dev(orig_list))) / math.Sqrt(float64(len(orig_list)))
	return lcl
}
