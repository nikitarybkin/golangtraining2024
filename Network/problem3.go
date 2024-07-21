package main

// данные пакеты нужны для системы проверки
import (
	"fmt"
	"io"
	"net/http"
	"net/url"
)

func main() {
	var name, age string
	fmt.Scan(&name, &age)

	base, err := url.Parse("http://127.0.0.1:8080/hello")
	if err != nil {
		return
	}

	params := url.Values{}
	params.Add("name", name)
	params.Add("age", age)

	base.RawQuery = params.Encode()

	fmt.Println(base.String())

	resp, err := http.Get(base.String())

	defer resp.Body.Close()

	data, err := io.ReadAll(resp.Body)
	if err != nil {
		return
	}

	fmt.Printf("%s", data)
}
