package main
import (
  "fmt"
)

struct person {
  name string
  balance float64
}


func main() {

  p1 := person{
    name : "Joseph Smith"
    balance : 10000
  }

  fmt.Println(p1.name, p1.balance)

}