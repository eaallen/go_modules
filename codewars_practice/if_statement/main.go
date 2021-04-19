package main

import (
	"fmt"
	"os"
)

func print(a ...interface{}) {
	fmt.Println(a...)
}

func main() {

	if args := os.Args[1:]; len(args) > 1 {
		fmt.Println("You have passed two arguments into this function. Only the first argument wll be used")
	} else if args[0] == "hello" {
		fmt.Println("good-bye")
	} else {
		print("Here is you first arguemnt:", args[0])
	}

	// fmt.Println(args[0])
}
