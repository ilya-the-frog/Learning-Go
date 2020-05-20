package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
)

func main() {
	readFile, _ := os.Open("input-201.txt")
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	var fileTextLines []string

	for fileScanner.Scan() {
		fileTextLines = append(fileTextLines, fileScanner.Text())
	}
	readFile.Close()

	output := Unique(fileTextLines)

	out, _ := os.OpenFile("input-201.a.txt", os.O_APPEND|os.O_WRONLY, 0644)
	_, err := fmt.Fprint(out, output)
	if err != nil {
		log.Fatal("can't write to output file")
	}
	_ = out.Close()
}

func Unique(strSlice []string) []string {
	keys := make(map[string]bool)
	list := []string{}
	for _, entry := range strSlice {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			list = append(list, entry)
		}
	}
	return list
}
