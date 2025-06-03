package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

type RequestData struct {
	Algoritmo1 int
	Algoritmo2 int
	Quantidade int
	Disposicao int
}

type Resultado struct {
	Duracao      time.Duration
	Comparacoes  int
	Movimentacoes int
}

type ResponseData struct {
	Algoritmo1 Resultado
	Algoritmo2 Resultado
	Relatorio  RelatorioComparativo
}

func executarAlgoritmo(num int, arr []int) (time.Duration, int, int) {
	switch num {
	case 1:
		return selection(arr)
	case 2:
		return insertion(arr)
	case 3:
		return boubble(arr)
	case 4:
		return shell(arr)
	case 5:
        comeco := 0 
        fim := len(arr)-1
		return quick(arr, comeco, fim)
	default:
		return 0, 0, 0
	}
}

type RelatorioComparativo struct {
	MelhorTempo           string
	MenosComparacoes      string
	MenosMovimentacoes    string
}




func executarHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	var dados RequestData
	err := json.NewDecoder(r.Body).Decode(&dados)
	if err != nil {
		http.Error(w, "Erro ao decodificar JSON", http.StatusBadRequest)
		return
	}

	fmt.Printf("Dados recebidos: Algoritmo1: %d, Algoritmo2: %d, Quantidade: %d, Disposicao: %d\n",
		dados.Algoritmo1, dados.Algoritmo2, dados.Quantidade, dados.Disposicao)

	arrayOriginal := size(dados.Quantidade, dados.Disposicao)
	array1 := make([]int, len(arrayOriginal))
	array2 := make([]int, len(arrayOriginal))
	copy(array1, arrayOriginal)
	copy(array2, arrayOriginal)

	dur1, comp1, mov1 := executarAlgoritmo(dados.Algoritmo1, array1)
	dur2, comp2, mov2 := executarAlgoritmo(dados.Algoritmo2, array2)

	fmt.Println("Resultado da comparação:")
	fmt.Printf("Algoritmo 1 - Duração: %v | Comparações: %d | Movimentações: %d\n", dur1, comp1, mov1)
	fmt.Printf("Algoritmo 2 - Duração: %v | Comparações: %d | Movimentações: %d\n", dur2, comp2, mov2)

	// Relatório comparativo
	relatorio := RelatorioComparativo{}

	// Tempo
	if dur1 < dur2 {
		relatorio.MelhorTempo = ("Algoritmo 1 foi mais eficiente em tempo.")
	} else if dur2 < dur1 {
		relatorio.MelhorTempo = ("Algoritmo 2 foi mais eficiente em tempo.")
	} else {
		relatorio.MelhorTempo = "Ambos algoritmos tiveram a mesma eficiência em tempo."
	}

	// Comparações
	if comp1 < comp2 {
		relatorio.MenosComparacoes = ("Algoritmo 1 realizou menos comparações.")
	} else if comp2 < comp1 {
		relatorio.MenosComparacoes = ("Algoritmo 2 realizou menos comparações.")
	} else {
		relatorio.MenosComparacoes = "Ambos algoritmos realizaram a mesma quantidade de comparações."
	}

	// Movimentações
	if mov1 < mov2 {
		relatorio.MenosMovimentacoes = ("Algoritmo 1 realizou menos movimentações.")
	} else if mov2 < mov1 {
		relatorio.MenosMovimentacoes = ("Algoritmo 2 realizou menos movimentações.")
	} else {
		relatorio.MenosMovimentacoes = "Ambos algoritmos realizaram a mesma quantidade de movimentações."
	}

	resposta := ResponseData{
		Algoritmo1: Resultado{Duracao: dur1, Comparacoes: comp1, Movimentacoes: mov1},
		Algoritmo2: Resultado{Duracao: dur2, Comparacoes: comp2, Movimentacoes: mov2},
		Relatorio:  relatorio,
	}

		fmt.Println("Relatório Comparativo:")
	fmt.Println(relatorio.MelhorTempo)
	fmt.Println(relatorio.MenosComparacoes)
	fmt.Println(relatorio.MenosMovimentacoes)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resposta)
}



func indexHandler(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("templates/index.html")
	if err != nil {
		http.Error(w, "Erro ao carregar template", http.StatusInternalServerError)
		return
	}
	tmpl.Execute(w, nil)
}

type TesteData struct {
	Algoritmo int
	Array     []int
}

func testarHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != "POST" {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	var dados TesteData
	err := json.NewDecoder(r.Body).Decode(&dados)
	if err != nil {
		http.Error(w, "Erro ao decodificar JSON", http.StatusBadRequest)
		return
	}

	fmt.Printf("Teste recebido: Algoritmo: %d, Array: %v\n", dados.Algoritmo, dados.Array)

	arrayCopia := make([]int, len(dados.Array))
	copy(arrayCopia, dados.Array)

	dur, comp, mov := executarAlgoritmo(dados.Algoritmo, arrayCopia)

	fmt.Printf("Resultado do teste - Duração: %v | Comparações: %d | Movimentações: %d\n", dur, comp, mov)

	resposta := Resultado{Duracao: dur, Comparacoes: comp, Movimentacoes: mov}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resposta)
}


func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/executar", executarHandler)
    //http.HandleFunc("/testar", testarHandler)


	fmt.Println("Servidor rodando em http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}
