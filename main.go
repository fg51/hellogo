package main

import "fmt"

func main() {
	// inn-out
	fmt.Print("Enter a number: ")
	var inn float64
	fmt.Scanf("%f", &inn)
	out := inn * 2
	fmt.Println(inn, " * 2 = ", out)

	// for-loop
	for i := 1; i <= 10; i++ {
		if i % 2 == 0 {
			fmt.Println(i, "even")
		} else {
			fmt.Println(i, "odd")
		}
	}
}
