package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
)

/*
Ввод	input.txt
Вывод	output.txt
Дано целое положительное число "target". Также дана последовательность из целых положительных чисел. Необходимо за
писать в выходной файл "1", если в последовательности есть два числа сумма, которых равна значению "target" или "0"
если таких нет.
*/
const (
	inputFile  = "input.txt"
	outputFile = "output.txt"
)

func main() {
	readFile, err := os.Open(inputFile)
	if err != nil {
		log.Fatal("can't open input file")
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	fileTextLines := make([]string, 0, 2)
	for fileScanner.Scan() {
		fileTextLines = append(fileTextLines, fileScanner.Text())
	}
	_ = readFile.Close()

	if len(fileTextLines) != 2 {
		log.Fatal("incorrect input file")
	}

	res := 0

	target, err := strconv.Atoi(fileTextLines[0]) //получили исходное число
	array := strings.Split(fileTextLines[1], " ") //превращем последовательность в слайс
	intArr := make([]int, 0, len(array))

	for _, v := range array {
		intEntry, err := strconv.Atoi(v)
		if err != nil {
			log.Fatal("incorrect input format")
		}
		intArr = append(intArr, intEntry)
	}

	res = hasSum(intArr, target)

	output, err := os.OpenFile(outputFile, os.O_APPEND|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatal("can't open output file")
	}
	defer output.Close()

	_, err = fmt.Fprint(output, res)
	if err != nil {
		log.Fatal("can't write to output file")
	}
}

func hasSum(intArr []int, target int) int {
	for i, v := range intArr {
		for n, value := range intArr {
			if i != n && v+value == target {
				return 1
			}
		}
	}
	return 0
}
