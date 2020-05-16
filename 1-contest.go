package main

import (
	"fmt"
)

/*
Даны два числа A и B. Вам нужно вычислить их сумму A+B. В этой задаче вам нужно
читать из стандартного ввода и выводить ответ в стандартный вывод
*/

func main() {
	var input1, input2 int
	fmt.Scan(&input1, &input2)
	fmt.Println(input1 + input2)
}
