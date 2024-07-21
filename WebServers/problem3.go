package main

import (
	"fmt"
	"net/http"
	"strconv"
)

var counter int

func handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {

		num, err1 := strconv.Atoi(r.FormValue("count"))
		if err1 != nil {
			w.Write([]byte("это не число"))
			w.WriteHeader(http.StatusBadRequest)
			return
		}

		counter += num

		return
	} else if r.Method == "GET" {
		w.Write([]byte(strconv.Itoa(counter)))

		return
	}
	w.Write([]byte("Отправлен не POST и не GET!"))
	w.Write([]byte(fmt.Sprintf("Hello, %s!", r.URL.Query().Get("name"))))
}

func main() {

	http.HandleFunc("/count", handler)

	err := http.ListenAndServe(":3333", nil)
	if err != nil {
		fmt.Println("Ошибка запуска сервера:", err)
	}
}
