package main

import (
	"encoding/json"
	"log"
	"net/http"
)

// Struct para representar um objeto de exemplo
type Todo struct {
	ID    string `json:"id"`
	Title string `json:"title"`
}

var todos []Todo

func main() {
	// Rotas da API
	http.HandleFunc("/todos", getTodos)
	http.HandleFunc("/todos/add", addTodo)

	// Inicia o servidor na porta 8080
	log.Fatal(http.ListenAndServe(":8080", nil))
}

// Manipulador para a rota GET /todos
func getTodos(w http.ResponseWriter, r *http.Request) {
	// Define o cabeçalho Content-Type como application/json
	w.Header().Set("Content-Type", "application/json")

	// Codifica o slice de todos em JSON
	json.NewEncoder(w).Encode(todos)
}

// Manipulador para a rota POST /todos/add
func addTodo(w http.ResponseWriter, r *http.Request) {
	// Define o cabeçalho Content-Type como application/json
	w.Header().Set("Content-Type", "application/json")

	// Decodifica o JSON do corpo da requisição em um objeto Todo
	var newTodo Todo
	err := json.NewDecoder(r.Body).Decode(&newTodo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	// Adiciona o novo todo ao slice
	todos = append(todos, newTodo)

	// Retorna o todo adicionado como resposta
	json.NewEncoder(w).Encode(newTodo)
}
