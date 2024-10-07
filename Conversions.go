package main

import (
	"fmt"
)

func listToFloat(orig_list []any) []float32 {
	new_list := make([]float32, len(orig_list))
	for i := 0; i <= len(orig_list); i++ {
		if float, err := orig_list[i].(float32); err {
			fmt.Printf("error converting")
		} else {
			new_list[i] = float
		}
	}
	return new_list
}
