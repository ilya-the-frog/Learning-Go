package main

// https://www.codewars.com/kata/5667e8f4e3f572a8f2000039
import (
	"fmt"
	"strings"
)

func main() {
	s := "ZpglnRxqenU"

	var value string

	split := strings.Split(s, "")
	result := make([]string, 0, len(split))

	/* debug purpose
	for i, v := range split {
		fmt.Printf("Split[")
		fmt.Printf("%d", i)
		fmt.Printf("] value is = ")
		fmt.Printf("%s", v)
		fmt.Printf("\n")
	}

	fmt.Printf("\n")
	*/

	for i, v := range split {
		value = strings.ToUpper(v)
		for n := 0; n < i; n++ {
			v = strings.ToLower(v)
			value = value + v
		}
		result = append(result, value)
	}
	res := strings.Join(result[:], "-")
	//return res
	fmt.Println(res)
}
