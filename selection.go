package main

import "time"

func selection (arr[]int)(time.Duration, int, int){
	inicio:= time.Now()

	mov := 0
	comp := 0
	n := len(arr)

	for i := 0; i < n-1; i++ {
		chave := i
		for j := i + 1; j < n; j++ {
			comp++
			if arr[j] < arr[chave] {
				chave = j
			}
		}
		if chave != i {
			arr[i], arr[chave] = arr[chave], arr[i]
			mov++
		}
	}
	duracao := time.Since(inicio)
	return duracao, comp, mov


}