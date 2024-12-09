package main

import "fmt"

var i int = 10

func main() {
	slc := [10]int{10, 21, 30, 41, 50, 61, 70, 81, 90, 101}
	for _, v := range slc {
		if v%2 == 0 {
			fmt.Println("number is even", v)
		} else {
			fmt.Println("number is odd", v)
		}
	}
}
