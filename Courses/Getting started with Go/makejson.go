package main

import (
	"encoding/json"
	"fmt"
)

/*
Write a program which prompts the user to first enter a name, and then enter an address.
Your program should create a map and add the name and address to the map using the keys
“name” and “address”, respectively. Your program should use Marshal() to create a JSON object
from the map, and then your program should print the JSON object.
*/

func main() {
	person := make(map[string]string)
	fmt.Println("Input name:")
	var name string
	fmt.Scanln(&name)
	person["name"] = name

	fmt.Println("Input address:")
	var address string
	fmt.Scanln(&address)
	person["address"] = address

	jsoned, _ := json.Marshal(person)
	fmt.Println(string(jsoned))
}
