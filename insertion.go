package main

import "time"

func insertion (arr[]int)(time.Duration, int, int){
	inicio:= time.Now()
	comp, mov := 0, 0

	for i := 1; i < len(arr); i++ {
	chave := arr[i]
	j := i - 1
		for j >= 0 {
			comp++
			if arr[j] > chave {
				arr[j+1] = arr[j]
				mov++
				j--
			} else {
				break
			}
		}
		arr[j+1] = chave
		mov++
	}	

	duracao := time.Since(inicio)
	return duracao, comp, mov


}