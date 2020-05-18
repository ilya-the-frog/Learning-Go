package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
)

/*
Даны два числа A и B. Вам нужно вычислить их сумму A+B. В этой задаче вам нужно
читать из файла и выводить ответ в файл
*/

func main() {
	input, _ := ioutil.ReadFile("input.txt")
	string := strings.Split(string(input), " ")
	v0, _ := strconv.Atoi(string[0])
	v1, _ := strconv.Atoi(string[1])
	res := strconv.Itoa(v0 + v1)
	output, _ := os.OpenFile("output.txt", os.O_APPEND|os.O_WRONLY, 0644)
	fmt.Fprintf(output, res)
	output.Close()
}
