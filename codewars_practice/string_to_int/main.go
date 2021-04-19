package main

import (
	"fmt"
	"math/rand"
	"strconv"
	"time"
)

func main() {
	i, err := strconv.Atoi("23")
	fmt.Print(i, err)

	str := strconv.Itoa(getRandomNumber())
	fmt.Print(str)
}

func getRandomNumber() int {
	var r = rand.New(rand.NewSource(time.Now().UnixNano()))
	return r.Intn(10000)
}
