package main

import (
    "fmt"
    "io/ioutil"
    "log"
)

func main() {
    conteudo, err := ioutil.ReadFile("arquivos/numeros_aleatorios.txt")
    if err != nil {
        log.Fatal(err)
    }

	

	for i := 0; i < 5; i++ {
    fmt.Println(string(conteudo[i]))
	}
}
