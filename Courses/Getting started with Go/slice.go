package main

import (
	"fmt"
	"strconv"
)

/*
Write a program which prompts the user to enter integers and stores the integers in a sorted slice.
The program should be written as a loop. Before entering the loop, the program should create an empty integer
slice of size (length) 3. During each pass through the loop, the program prompts the user to enter an integer
to be added to the slice. The program adds the integer to the slice, sorts the slice, and prints the contents of
the slice in sorted order. The slice must grow in size to accommodate any number of integers which the user decides
to enter. The program should only quit (exiting the loop) when the user enters the character ‘X’ instead of an integer.
*/

func main() {
	x := make([]int, 3)
	fmt.Scanln(&x[0], &x[1], &x[2])
	bubbleSort(x)
	fmt.Println("Starting slice:")
	fmt.Printf("%v\n", x)

	var isWorking bool = true
	var v string
	var i int

	for isWorking {
		fmt.Println("Add 1 more int")
		fmt.Scanln(&v)
		if v == "x" {
			isWorking = false
			break
		}
		i, _ = strconv.Atoi(v)
		x = append(x, i)
		bubbleSort(x)
		fmt.Printf("%v\n", x)
	}
}

func bubbleSort(s []int) []int {
	for n := 1; n < len(s); n++ {
		for i := 0; i < len(s)-1; i++ {
			if s[i] > s[i+1] {
				s[i+1], s[i] = s[i], s[i+1]
			}
		}
	}
	return s
}
