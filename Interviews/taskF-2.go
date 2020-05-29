package main

import (
	"bufio"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

const (
	in  = "input.txt"
	out = "output.txt"
)

func main() {

	file, err := os.Open(in)
	if err != nil {
		writeResult(0)
		log.Fatalf("[Error] Not able to open file %v: %v\n", in, err)
	}
	defer file.Close()

	buff := make(map[int]int)
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	target, _ := strconv.Atoi(scanner.Text())
	for scanner.Scan() {
		number, _ := strconv.Atoi(scanner.Text())
		if number > target {
			continue
		}
		buff[number]++

		if repitTimes, ok := buff[target-number]; ok {
			if number == target-number && repitTimes <= 1 {
				continue
			}
			writeResult(1)
			return
		}
	}
	writeResult(0)
}

func writeResult(result int) {
	err := ioutil.WriteFile(out, []byte(strconv.Itoa(result)+"\n"), 0644)
	if err != nil {
		log.Fatalf("[Error] Not bale to write result in file %v: %v\n", out, err)
		return
	}
}
