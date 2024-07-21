package main

import (
	"archive/zip"

	"encoding/csv"

	"fmt"
)

func main() {

	r, _ := zip.OpenReader("E:\\Projects\\GoLang\\Study\\src\\Files\\task.zip") //открываем файл архива

	defer r.Close() // после выполнения всего закрываем файл архива

	for _, file := range r.File { //перебираем все файлы

		if !file.FileInfo().IsDir() { // если файл есть, то

			ff, _ := file.Open() // открываем каждый файл в цикле

			if data, _ := csv.NewReader(ff).ReadAll(); len(data) == 10 { //проверка текущего файла в цикле, если ок то ---->

				fmt.Println(data[4][2]) // печатаем значения

			}

			ff.Close() // закрываем текущий файл в цикле и переходим к следующему

		}

	}

}
