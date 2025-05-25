package main

import "fmt"

func main() {
	arr := []int{23, 1, 10, 5, 2}
	for _, v := range arr {
		println(v)
	}

	duracao, comp, mov := selection(arr)

    fmt.Printf("Tempo de execução: %v\n", duracao)
    fmt.Printf("comparações: %d, movimentacoes: %d\n", comp, mov)

}
