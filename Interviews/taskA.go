package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

/*
На вход программе подается большое количество целых чисел. Все числа, кроме одного, имеют пару,
причем может быть несколько одинаковых пар. Найдите число без пары.
*/

func main() {
	readFile, _ := os.Open("input-201.txt")
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileTextLines []string

	for fileScanner.Scan() {
		fileTextLines = append(fileTextLines, fileScanner.Text())
	}
	readFile.Close()

	res := strings.Join(RemoveDuplicates(fileTextLines), "")
	output, _ := os.OpenFile("input-201.a.txt", os.O_APPEND|os.O_WRONLY, 0644)
	fmt.Fprintf(output, res)
	output.Close()
}

func RemoveDuplicates(strSlice []string) []string {
	var result []string
	for i, value := range strSlice {
		for n, val := range strSlice {

			if i != n && value != val && !checkDuplicates(strSlice, value) {
				result = append(result, value)
				break
			}
		}
	}
	return result
}

func checkDuplicates(s []string, item string) bool {
	n := 0
	for _, a := range s {
		if a == item {
			n++
		}
	}
	if n == 1 {
		return false
	} else {
		return true
	}

}
