package main

import (
	"fmt"
	"os"
)

import "github.com/mitchellh/cli"

import a "github.com/kflange/hellogo"


func main() {
	c := cli.NewCLI("hellogo", "0.4.0")
	c.Args = os.Args[1:]
	c.Commands = map[string]cli.CommandFactory {
		"inout": func() (cli.Command, error) {
			return &Inout{}, nil
		},
		"loop": func() (cli.Command, error) {
			return &Loop{}, nil
		},
		"animal": func() (cli.Command, error) {
			return &AnimalInterface{}, nil
		},
	}
	exitStatus, err := c.Run()
	if err != nil {
		fmt.Println(err)
	}
	os.Exit(exitStatus)
}


type Inout struct {}

func (self *Inout) Help () string {
	// ./app --help Inout
	return "sample stdin stdout"
}

func (self *Inout) Synopsis() string {
	// ./app --help
	return "sample stdin stdout from synopsis"
}

func (self *Inout) Run(args []string) int {
	fmt.Print("Enter a number: ")
	var inn float64
	fmt.Scanf("%f", &inn)
	out := inn * 2
	fmt.Println(inn, " * 2 = ", out)
	return 0
}


type Loop struct {}

func (self *Loop) Help () string {
	// ./app --help loop
	return "sample for-loop"
}

func (self *Loop) Synopsis() string {
	// ./app --help
	return "sample for-loop from synopsis"
}

func (self *Loop) Run(args []string) int {
	for i := 1; i <= 10; i++ {
		if i % 2 == 0 {
			fmt.Println(i, "even")
		} else {
			fmt.Println(i, "odd")
		}
	}
	return 0
}


type AnimalInterface struct {}

func (self *AnimalInterface) Help () string {
	// ./app --help animal
	return "sample interface"
}

func (self *AnimalInterface) Synopsis() string {
	// ./app --help
	return "sample interface from synopsis"
}

func (self *AnimalInterface) Run(args []string) int {
	// interface
	var animal a.Animal = new(a.Dog)
	a.LetAnimalCry(animal)
	return 0
}
