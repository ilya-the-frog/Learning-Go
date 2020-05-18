package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	readFile, _ := os.Open("input.txt")
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileTextLines []string
	for fileScanner.Scan() {
		fileTextLines = append(fileTextLines, fileScanner.Text())
	}
	readFile.Close()

	target, _ := strconv.Atoi(fileTextLines[0])   //получили исходное число
	array := strings.Split(fileTextLines[1], " ") //превращем последовательность в слайс

	//переделать слайс строк в слайс интов
	var intArr []int

	for i, _ := range array {
		int, _ := strconv.Atoi(array[i])
		intArr = append(intArr, int)
	}

	var res int

	for i, v := range intArr {
		for n, value := range intArr {
			if i != n && v+value == target {
				res = 1
			} else {
				continue
			}
		}
	}

	result := strconv.Itoa(res)
	output, _ := os.OpenFile("output.txt", os.O_APPEND|os.O_WRONLY, 0644)
	fmt.Fprintf(output, result)
	output.Close()
}
