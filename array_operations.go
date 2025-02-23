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

func order_dev_rank(orig_array []dev_rank) []dev_rank{
	var store_smaller float64
	var store_configurado string
	ordered:=false
	for !ordered {
		counter:=0
		for i:=0; i<len(orig_array);i+=1 {	
			if i == len(orig_array)-1 {
				break
			}
			if orig_array[i].Value < orig_array[i+1].Value {
				store_smaller = orig_array[i].Value
				store_configurado = orig_array[i].Configurado
				orig_array[i].Value=orig_array[i+1].Value
				orig_array[i].Configurado = orig_array[i+1].Configurado
				orig_array[i+1].Value=store_smaller
				orig_array[i+1].Configurado = store_configurado
				counter +=1
			}
		}
		if counter == 0 {
			ordered=true
		}
	}
	return orig_array
}

func separate_db_array (db_list []dev_rank) [][]dev_rank {
	var new_list [][]dev_rank
	var separation_list []dev_rank
	for c:=0;c <len(db_list);c++ {
		if c==0{
			separation_list=append(separation_list,db_list[c])
			c+=1
		}
		if db_list[c].Configurado !=db_list[c-1].Configurado{
		new_list = append(new_list,separation_list)
		separation_list= nil
		}
		separation_list=append(separation_list,db_list[c])
	}
	return new_list
}

func rank_list_avg (orig_list []dev_rank) []dev_rank {
	list:=separate_db_array(orig_list)
	var sum float64
	var avg float64
	var count float64
	var configurado_avg dev_rank
	var dev_rank_list []dev_rank
	i:=0
	for c:=0;c <len(list);c++{
 		for i=0;i<len((list[c]))-1;i++ {
		sum+=list[c][i].Value
		count+=1
		}
		avg = sum/count
		sum=0
		count=0
		configurado_avg.Configurado=list[c][i].Configurado
		configurado_avg.Value=avg
		dev_rank_list=append(dev_rank_list,configurado_avg)
	}
	return dev_rank_list
}

func rank_list_devs (orig_list []dev_rank) []dev_rank {
	list:=separate_db_array(orig_list)
	avg_list:=rank_list_avg(orig_list)
	var cp_list []dev_rank
	var dev dev_rank
	var dev_sum float64
	var std_dev float64
	var ucl float64
	var lcl float64
	var cp float64
	i:=0
	for c:=0;c <len(list);c++{
 		for i=0;i<len((list[c]));i++ {
			if len(list[c])==1{
				c+=1
			}
			if i==len(list[c]){
				break
			}
			dev:= (list[c][i].Value -avg_list[c].Value)*(list[c][i].Value -avg_list[c].Value)
			dev_sum += dev
		}
		std_dev= math.Sqrt((dev_sum) / (float64(i-1)))
		dev_sum = 0
		ucl = (avg_list[c].Value + (3 * std_dev))
		lcl = (avg_list[c].Value - (3* std_dev))
		cp = (ucl-lcl)/(6*std_dev)
		dev.Configurado = list[c][i-1].Configurado
		dev.Value = cp
		cp_list = append(cp_list, dev)
}	 
		return cp_list
}