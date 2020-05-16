// solving http://golang-book.ru/chapter-06-arrays-slices-maps.html

package main

import "fmt"

func main() {
	fmt.Println("Ex1 Как обратиться к четвертому элементу массива или среза?")
	var mas1 [6]int
	fourth := mas1[3]   // при счете с 0, третий элемент будет 4ым
	fmt.Println(fourth) //чтобы работало
	fmt.Println("")

	fmt.Println("Ex2 Чему равна длина среза, созданного таким способом?")
	fmt.Println(len(make([]int, 3, 9))) // 3, срез не мб больше массива
	fmt.Println("")

	fmt.Println("Дан массив: x := [6]string{\"a\",\"b\",\"c\",\"d\",\"e\",\"f\"}")
	fmt.Println("Что вернет вам x[2:5]?")
	x := [6]string{"a", "b", "c", "d", "e", "f"}
	fmt.Println(x[2:5]) //cde, 2-3-4 элемент массива
	fmt.Println("")

	fmt.Println("Ex4 Напишите программу, которая находит наименьший элемент в этом списке:")
	z := []int{
		48, 96, 86, 68,
		57, 82, 63, 70,
		37, 34, 83, 27,
		19, 97, 9, 17,
	}
	var lowest int = z[0] //если присвоить ноль, то ноль так и останется.

	for i := 0; i < len(z); i++ {
		fmt.Print(z[i], " ")
		if lowest > z[i] {
			lowest = z[i]
		}
	}
	fmt.Println("\n", lowest)
}
