package main

import (
	"encoding/json"
	"fmt"
)

func main() {
	var name, address string
	fmt.Println("Enter Name:")
	fmt.Scanln(&name)
	fmt.Println("Enter Address:")
	fmt.Scanln(&address)
	maps := map[string]string{name: address}
	data, err := json.Marshal(maps)
	if err != nil {
		fmt.Println("Error", err.Error())
	}

	fmt.Println(string(data))
}
