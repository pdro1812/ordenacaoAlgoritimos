package main

import "fmt"

func main() {
	arr := []int{23, 1, 10, 5, 2}
	

	duracao, comp, mov := quick(arr, 0, len(arr)-1)
	for _, v := range arr {
			println(v)
		}

    fmt.Printf("Tempo de execução: %v\n", duracao)
    fmt.Printf("comparações: %d, movimentacoes: %d\n", comp, mov)

}
