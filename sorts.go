package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	//создаем слайс и заполняем его рандомными циферками
	slice := make([]int, 5)
	rand.Seed(time.Now().UnixNano())
	for i := range slice {
    slice[i] = rand.Intn(100)
	}
	//делаем сортировки
	fmt.Println("\nSlice by default:", slice,)
	fmt.Println("\nSlice bubleSorted", bubleSort(slice))
}

func bubleSort(s []int) []int {
	for n := 1; n < len(s); n++ {
	for i := 0; i < len(s) - 1; i++ {

		if s[i] > s[i+1] {
			//fmt.Println("Swapping", s[i], "&", s[i+1])
			s[i+1], s[i] = s[i], s[i+1]
			//fmt.Println("Swapped", s[i], "&", s[i+1])
		}
		}
	}
	return s
}
