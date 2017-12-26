package main

import "fmt"

func main() {
	fmt.Print("Enter a number: ")
	var inn float64
	fmt.Scanf("%f", &inn)
	out := inn * 2
	fmt.Println(inn, " * 2 = ", out)
}
