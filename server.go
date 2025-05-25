package main

import (
    "net/http"
    "html/template"
	"log"

)

func minhaFuncao() string {
    return "Mensagem vinda do Go via template externo!"
}


// Handler para a página principal
func homeHandler(w http.ResponseWriter, r *http.Request) {
    // Carrega o template externo
    tmpl, err := template.ParseFiles("templates/index.html")
    if err != nil {
        http.Error(w, "Erro ao carregar template", http.StatusInternalServerError)
        return
    }

    dados := struct {
        Mensagem string
    }{
        Mensagem: minhaFuncao(),
    }

    // Executa o template passando os dados
    err = tmpl.Execute(w, dados)
    if err != nil {
        log.Println("Erro ao executar template:", err)
    }
}
