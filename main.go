package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"log"
	"net/http"
	"time"
)

// templates armazena os templates HTML pré-carregados para otimização.
var templates *template.Template

// RequestData define a estrutura dos dados recebidos do frontend.
type RequestData struct {
	Algoritmo1 int
	Algoritmo2 int
	Quantidade int
	Disposicao int
}

// Structs formatadas para a resposta JSON enviada ao cliente.
type ResultadoJSON struct {
	Duracao      int64 `json:"Duracao"`
	Comparacoes  int   `json:"Comparacoes"`
	Movimentacoes int   `json:"Movimentacoes"`
}

type ResponseDataJSON struct {
	Algoritmo1 ResultadoJSON        `json:"Algoritmo1"`
	Algoritmo2 ResultadoJSON        `json:"Algoritmo2"`
	Relatorio  RelatorioComparativo `json:"Relatorio"`
}

type RelatorioComparativo struct {
	MelhorTempo        string `json:"MelhorTempo"`
	MenosComparacoes   string `json:"MenosComparacoes"`
	MenosMovimentacoes string `json:"MenosMovimentacoes"`
}

func main() {
	// Carrega os templates na inicialização. O programa falhará se os arquivos não forem encontrados.
	templates = template.Must(template.ParseFiles("templates/index.html"))

	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/executar", executarHandler)

	fmt.Println("Servidor rodando em http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// indexHandler serve a página HTML principal.
func indexHandler(w http.ResponseWriter, r *http.Request) {
	err := templates.ExecuteTemplate(w, "index.html", nil)
	if err != nil {
		log.Printf("ERRO: Falha ao executar o template: %v", err)
		http.Error(w, "Erro interno ao renderizar a página", http.StatusInternalServerError)
	}
}

// executarHandler recebe os dados, executa os algoritmos e retorna o resultado em JSON.
func executarHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	var dados RequestData
	err := json.NewDecoder(r.Body).Decode(&dados)
	if err != nil {
		log.Printf("ERRO: Falha ao decodificar JSON do request: %v", err)
		http.Error(w, "Erro ao decodificar JSON", http.StatusBadRequest)
		return
	}

	// Validação básica dos dados de entrada.
	if dados.Algoritmo1 < 1 || dados.Algoritmo1 > 5 || dados.Algoritmo2 < 1 || dados.Algoritmo2 > 5 {
		http.Error(w, "ID de algoritmo inválido. Deve ser entre 1 e 5.", http.StatusBadRequest)
		return
	}
	if dados.Disposicao < 1 || dados.Disposicao > 3 {
		http.Error(w, "ID de disposição de dados inválido. Deve ser entre 1 e 3.", http.StatusBadRequest)
		return
	}

	arrayOriginal := size(dados.Quantidade, dados.Disposicao)
	array1 := make([]int, len(arrayOriginal))
	array2 := make([]int, len(arrayOriginal))
	copy(array1, arrayOriginal)
	copy(array2, arrayOriginal)

	dur1, comp1, mov1 := executarAlgoritmo(dados.Algoritmo1, array1)
	dur2, comp2, mov2 := executarAlgoritmo(dados.Algoritmo2, array2)

	relatorio := gerarRelatorio(dur1, dur2, comp1, comp2, mov1, mov2)

	respostaJSON := ResponseDataJSON{
		Algoritmo1: ResultadoJSON{
			Duracao:      dur1.Milliseconds(),
			Comparacoes:  comp1,
			Movimentacoes: mov1,
		},
		Algoritmo2: ResultadoJSON{
			Duracao:      dur2.Milliseconds(),
			Comparacoes:  comp2,
			Movimentacoes: mov2,
		},
		Relatorio: relatorio,
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(respostaJSON)
}

// gerarRelatorio cria o texto comparativo a partir dos resultados dos algoritmos.
func gerarRelatorio(dur1, dur2 time.Duration, comp1, comp2, mov1, mov2 int) RelatorioComparativo {
	relatorio := RelatorioComparativo{}

	if dur1 < dur2 {
		relatorio.MelhorTempo = "Algoritmo 1 foi mais eficiente em tempo."
	} else if dur2 < dur1 {
		relatorio.MelhorTempo = "Algoritmo 2 foi mais eficiente em tempo."
	} else {
		relatorio.MelhorTempo = "Ambos algoritmos tiveram a mesma eficiência em tempo."
	}

	if comp1 < comp2 {
		relatorio.MenosComparacoes = "Algoritmo 1 realizou menos comparações."
	} else if comp2 < comp1 {
		relatorio.MenosComparacoes = "Algoritmo 2 realizou menos comparações."
	} else {
		relatorio.MenosComparacoes = "Ambos algoritmos realizaram a mesma quantidade de comparações."
	}

	if mov1 < mov2 {
		relatorio.MenosMovimentacoes = "Algoritmo 1 realizou menos movimentações."
	} else if mov2 < mov1 {
		relatorio.MenosMovimentacoes = "Algoritmo 2 realizou menos movimentações."
	} else {
		relatorio.MenosMovimentacoes = "Ambos algoritmos realizaram a mesma quantidade de movimentações."
	}
	return relatorio
}

// executarAlgoritmo atua como um 'roteador' para a função de ordenação correta.
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
		fim := len(arr) - 1
		return quick(arr, comeco, fim)
	default:
		return 0, 0, 0
	}
}

