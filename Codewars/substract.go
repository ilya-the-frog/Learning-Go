package main

/*
Subtract the sum
NOTE! This kata can be more difficult than regular 8-kyu katas

Complete the function which get an input number n such that n >= 10 and n < 10000, then:

Sum all the digits of n.
Subtract the sum from n, and it is your new n.
If the new n is in the list below return the associated fruit, otherwise return back to task 1.
Example
n = 325
sum = 3+2+5 = 10
n = 325-10 = 315 (not in the list)
sum = 3+1+5 = 9
n = 315-9 = 306 (not in the list)
sum = 3+0+6 = 9
n =306-9 = 297 (not in the list)
. .
. ...until you find the first n in the list below.

https://www.codewars.com/kata/56c5847f27be2c3db20009c3/go
*/
import (
	"fmt"
)

func main() {
	n := 4711
	fmt.Println(SubtractSum(n))
}

func SubtractSum(n int) string {
	sum := 0
	baseNum := n
	var res string

	for baseNum > 100 {
		for n > 0 {
			sum += n % 10
			n = n / 10
		}
		baseNum = baseNum - sum
	}

	switch baseNum {
	case 1:
		res = "kiwi"
	case 2:
		res = "pear"
	case 3:
		res = "kiwi"
	case 4:
		res = "banana"
	case 5:
		res = "melon"
	case 6:
		res = "banana"
	case 7:
		res = "melon"
	case 8:
		res = "pineapple"
	case 9:
		res = "apple"
	case 10:
		res = "pineapple"
	case 11:
		res = "cucumber"
	case 12:
		res = "pineapple"
	case 13:
		res = "cucumber"
	case 14:
		res = "orange"
	case 15:
		res = "grape"
	case 16:
		res = "orange"
	case 17:
		res = "grape"
	case 18:
		res = "apple"
	case 19:
		res = "grape"
	case 20:
		res = "cherry"
	case 21:
		res = "pear"
	case 22:
		res = "cherry"
	case 23:
		res = "pear"
	case 24:
		res = "kiwi"
	case 25:
		res = "banana"
	case 26:
		res = "kiwi"
	case 27:
		res = "apple"
	case 28:
		res = "melon"
	case 29:
		res = "banana"
	case 30:
		res = "melon"
	case 31:
		res = "pineapple"
	case 32:
		res = "melon"
	case 33:
		res = "pineapple"
	case 34:
		res = "cucumber"
	case 35:
		res = "orange"
	case 36:
		res = "apple"
	case 37:
		res = "orange"
	case 38:
		res = "grape"
	case 39:
		res = "orange"
	case 40:
		res = "grape"
	case 41:
		res = "cherry"
	case 42:
		res = "pear"
	case 43:
		res = "cherry"
	case 44:
		res = "pear"
	case 45:
		res = "apple"
	case 46:
		res = "pear"
	case 47:
		res = "kiwi"
	case 48:
		res = "banana"
	case 49:
		res = "kiwi"
	case 50:
		res = "banana"
	case 51:
		res = "melon"
	case 52:
		res = "pineapple"
	case 53:
		res = "melon"
	case 54:
		res = "apple"
	case 55:
		res = "cucumber"
	case 56:
		res = "pineapple"
	case 57:
		res = "cucumber"
	case 58:
		res = "orange"
	case 59:
		res = "cucumber"
	case 60:
		res = "orange"
	case 61:
		res = "grape"
	case 62:
		res = "cherry"
	case 63:
		res = "apple"
	case 64:
		res = "cherry"
	case 65:
		res = "pear"
	case 66:
		res = "cherry"
	case 67:
		res = "pear"
	case 68:
		res = "kiwi"
	case 69:
		res = "pear"
	case 70:
		res = "kiwi"
	case 71:
		res = "banana"
	case 72:
		res = "apple"
	case 73:
		res = "banana"
	case 74:
		res = "melon"
	case 75:
		res = "pineapple"
	case 76:
		res = "melon"
	case 77:
		res = "pineapple"
	case 78:
		res = "cucumber"
	case 79:
		res = "pineapple"
	case 80:
		res = "cucumber"
	case 81:
		res = "apple"
	case 82:
		res = "grape"
	case 83:
		res = "orange"
	case 84:
		res = "grape"
	case 85:
		res = "cherry"
	case 86:
		res = "grape"
	case 87:
		res = "cherry"
	case 88:
		res = "pear"
	case 89:
		res = "cherry"
	case 90:
		res = "apple"
	case 91:
		res = "kiwi"
	case 92:
		res = "banana"
	case 93:
		res = "kiwi"
	case 94:
		res = "banana"
	case 95:
		res = "melon"
	case 96:
		res = "banana"
	case 97:
		res = "melon"
	case 98:
		res = "pineapple"
	case 99:
		res = "apple"
	case 100:
		res = "pineapple"
	default:
		SubtractSum(baseNum)
	}

	return res
}
