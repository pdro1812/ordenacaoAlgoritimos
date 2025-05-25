package main

import "time"

func boubble (arr[]int)(time.Duration, int, int){
	inicio:= time.Now()
	comp, mov := 0, 0
	n := len(arr)

	
	for i := 0; i < n-1; i++ {
		 for j:= 0 ; j<n-1-i ; j++{
			comp++
            if(arr[j] > arr[j+1]){
                temp := arr[j];
                arr[j] = arr[j+1];
                arr[j+1] = temp;
				mov++
            }
	}
}

	duracao := time.Since(inicio)
	return duracao, comp, mov

}