package main

import "fmt"
import    "net/http"


func main() {
    http.HandleFunc("/", homeHandler)
    fmt.Println("Servidor rodando em http://localhost:8080")
    http.ListenAndServe(":8080", nil)
}
