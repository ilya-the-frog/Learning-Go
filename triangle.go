package kata

//is that a triangle kata from codewars
//https://www.codewars.com/kata/56606694ec01347ce800001b

func IsTriangle(a, b, c int) bool {
	if a+b > c {
		if a+c > b {
			if b+c > a {
				return true
			}
		}
	}
	return false
}
