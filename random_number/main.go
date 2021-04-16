package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// create and send the generater
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	fmt.Print(r.Int63n(10000))
}
