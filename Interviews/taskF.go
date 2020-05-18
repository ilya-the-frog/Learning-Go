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

	var res int
	var intArr []int

	target, _ := strconv.Atoi(fileTextLines[0]) //получили исходное число
	if len(fileTextLines) == 2 {
		array := strings.Split(fileTextLines[1], " ") //превращем последовательность в слайс

		for i, _ := range array {
			intEntry, _ := strconv.Atoi(array[i])
			intArr = append(intArr, intEntry)
		}

		for i, v := range intArr {
			for n, value := range intArr {
				if i != n && v+value == target {
					res = 1
				} else {
					continue
				}
			}
		}
	} else {
		res = 0
	}

	result := strconv.Itoa(res)
	output, _ := os.OpenFile("output.txt", os.O_APPEND|os.O_WRONLY, 0644)
	fmt.Fprintf(output, result)
	output.Close()
}
