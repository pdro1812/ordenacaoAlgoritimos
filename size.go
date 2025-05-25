package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
)

func size(size int, typ int) []int {
	switch typ {
	case 1:
		file, err := os.Open("arquivos/numeros_aleatorios.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		result := make([]int, 0, size)

		for scanner.Scan() && len(result) < size {
			line := scanner.Text()
			num, err := strconv.Atoi(line)
			if err != nil {
				log.Fatal(err)
			}
			result = append(result, num)
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}

		return result

	case 2:
		file, err := os.Open("arquivos/numeros_ordem_decrescente.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		result := make([]int, 0, size)

		for scanner.Scan() && len(result) < size {
			line := scanner.Text()
			num, err := strconv.Atoi(line)
			if err != nil {
				log.Fatal(err)
			}
			result = append(result, num)
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}

		return result

	case 3:
		file, err := os.Open("arquivos/numeros_ordenados.txt")
		if err != nil {
			log.Fatal(err)
		}
		defer file.Close()

		scanner := bufio.NewScanner(file)
		result := make([]int, 0, size)

		for scanner.Scan() && len(result) < size {
			line := scanner.Text()
			num, err := strconv.Atoi(line)
			if err != nil {
				log.Fatal(err)
			}
			result = append(result, num)
		}

		if err := scanner.Err(); err != nil {
			log.Fatal(err)
		}

		return result

	default:
		return nil
	}	
}