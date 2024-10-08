package main

import (
	"math"
)

func std_dev(conv_list []float64) float64 {
	var sum float64
	var count float64
	var dev_sum float64
	var dev float64
	for _, element := range conv_list {
		sum += element
		count += 1
	}
	avg := sum / count
	for _, element := range conv_list {
		dev = (element - avg) * (element - avg)
		dev_sum += (dev)
	}
	return math.Sqrt((dev_sum) / (count - 1))
}
