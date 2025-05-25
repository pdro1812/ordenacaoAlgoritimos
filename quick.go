package main

import "time"

var comp int
var mov int

func trocar(arr[]int, i, j int){
	arr[i], arr[j] = arr[j], arr[i]
	mov++
}


func particionar(arr []int, inicio, fim int) int {
	pivo := arr[fim]
	i := inicio-1

	for j := inicio; j<fim; j++{
		comp++
		if arr[j] <= pivo{
			i++
			trocar(arr, i, j)
		}
	}

	trocar(arr, i+1,fim)
	return i+1
}

func quick (arr[]int, comeco, fim int)(time.Duration, int, int){
	inicio:= time.Now()
		

	if comeco<fim{
		p :=particionar(arr, comeco, fim)
		quick(arr, comeco, p-1)
		quick(arr, p+1, fim)
	}
	

	duracao := time.Since(inicio)
	return duracao, comp, mov

}