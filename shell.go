package main

import "time"

func shell (arr[]int)(time.Duration, int, int){
	inicio:= time.Now()
	comp, mov := 0, 0
	n := len(arr)

	
	var intervalo, temp, j int

    for intervalo = n / 2; intervalo > 0; intervalo /= 2{
        for i := intervalo; i < n; i++ {
            temp = arr[i];
            for j = i; j >= intervalo && arr[j - intervalo] > temp; j -= intervalo {
                arr[j] = arr[j - intervalo];
				mov++
				comp++
            }
			if j != i {
				arr[j] = temp
				mov++
				comp++
			}
        }
    }


	duracao := time.Since(inicio)
	return duracao, comp, mov

}