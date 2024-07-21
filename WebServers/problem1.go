package main

import (
	"fmt"
	"net/http"
)

// Обработчик HTTP-запросов
func handler0(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello, web!"))
}

func main() {
	http.HandleFunc("/get", handler0)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Println("Ошибка запуска сервера:", err)
	}
}
