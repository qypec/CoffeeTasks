package main

import (
	"fmt"
	"math/rand"
	"os"
)

const limit = 1000

// go run testdata/generator.go
func main() {
	file, err := os.Create("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fmt.Fprintf(file, "999999998 1 ")
	for i := 1; i <= limit; i++ {
		fmt.Fprintf(file, "99999999%v ", rand.Intn(8))
		fmt.Println(i)
	}
}
