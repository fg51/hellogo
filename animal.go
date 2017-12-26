package main

import "fmt"


type Animal interface {
	Cry()
}


type Dog struct {}
func (this *Dog) Cry() {
	fmt.Println("bow bow")
}

type Cat struct {}
func (this *Cat) Cry() {
	fmt.Println("mew mew")
}

func LetAnimalCry(a Animal) {
	a.Cry()
}
