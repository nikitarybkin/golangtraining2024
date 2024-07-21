package main

import (
	"fmt"
	"net/http"
)

func handler1(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte(fmt.Sprintf("Hello, %s!", r.URL.Query().Get("name"))))
}

func main() {

	http.HandleFunc("/api/user", handler1)

	err := http.ListenAndServe(":9000", nil)
	if err != nil {
		fmt.Println("Ошибка запуска сервера:", err)
	}
}
