package kata

//is that a triangle kata from codewars

func IsTriangle(a, b, c int) bool {
  if a + b > c {
    if a + c > b {
      if b + c > a {
        return true
       }
    }
  }
  return false
}
