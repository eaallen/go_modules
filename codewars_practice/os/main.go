package main

import (
	"fmt"
	"os"
)

func main() {
	file, err := os.Create("./test.txt")
	if err == nil {
		// var str string = "test 1"
		var file_text string = file.Read(100)
		// file.WriteString(str)
		fmt.Println(file_text)
	} else {
		fmt.Println(err)
	}
}
