package main

import "fmt"

// solving https://thecode.media/hr/
func main () {
  //office 1
  emp := 1

  for i := 1; i < 30; i++ {
    emp = emp + emp
  }
  fmt.Println("В старом офисе было", emp, "сотрудников")

  new_emp := 2
  days := 0

  for new_emp <= emp {
    new_emp = new_emp * 2
    days++
  }

  fmt.Println("На новый офис потребовалось", days, "дней")
}
