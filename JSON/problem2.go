package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"os"
)

type GlobalIds []struct {
	Id int `json:"global_id"`
}

/*type Item struct {
	Properties Prop `json:"properties"`
}

type Codex struct {

}*/

func main() {
	file, err := os.Open("E:\\Projects\\GoLang\\Study\\src\\JSON\\data-20190514T0100.json")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	var sum int
	rd := bufio.NewReader(file)
	reader := json.NewDecoder(rd)
	gl := new(GlobalIds)
	err = reader.Decode(&gl)
	if err != nil {
		panic(err)
	}
	for _, v := range *gl {
		sum += v.Id
	}
	fmt.Println(sum)
}
