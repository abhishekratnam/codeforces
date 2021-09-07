package main

import (
	"fmt"
	"io/ioutil"
	"strings"
)

type Name struct {
	fname string
	lname string
}

func main() {
	data, err := ioutil.ReadFile("file.txt")
	if err != nil {
		panic(err)
	}

	dat := string(data[:])
	res1 := strings.Split(dat, ",")
	model := []Name{}
	// var name string
	for i, v := range res1 {
		name := strings.Split(v, " ")
		model = append(model, Name{fname: name[i], lname: name[i+1]})
	}
	fmt.Println(model)
}
