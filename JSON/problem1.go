package main

import (
	"encoding/json"
	"io"
	"os"
)

type Group struct {
	ID       int
	Number   string
	Year     int
	Students []Student
}

type Student struct {
	LastName   string
	FirstName  string
	MiddleName string
	Birthday   string
	Adress     string
	Phone      string
	Rating     []int
}

type Avg struct {
	Average float64
}

func main() {
	data, err := io.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	var cnt float64
	gr := new(Group)
	err = json.Unmarshal(data, &gr)
	for _, student := range gr.Students {
		cnt += float64(len(student.Rating))
	}
	avg := Avg{cnt / float64(len(gr.Students))}
	datOut, _ := json.MarshalIndent(avg, "", "    ")
	_, err = os.Stdout.Write(datOut)

}
