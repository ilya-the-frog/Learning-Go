package main

import (
	"fmt"
	"math/big"
)

/*
Даны два числа неотрицательных числа A и B(числа могут содержать до 1000 цифр). Вам нужно вычислить их сумму.
*/

func main() {
	var v0, v1 string
	fmt.Scanf("%s %s", &v0, &v1)

	var big0, _ = new(big.Int).SetString(v0, 0)
	var big1, _ = new(big.Int).SetString(v1, 0)

	count := big.NewInt(0)

	count.Add(count, big0)
	count.Add(count, big1)

	fmt.Printf("%v", count)
}
