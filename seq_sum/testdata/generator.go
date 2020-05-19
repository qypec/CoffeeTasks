package main

import (
	"fmt"
	"math/rand"
	"os"
)

const limit = 5000000

// go run testdata/generator.go
func main() {
	file, err := os.Create("input.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fmt.Fprintf(file, "678978", )
	for i := 1; i <= limit; i++ {
		fmt.Fprintf(file, "%v ", rand.Intn(999999995))
		fmt.Println(i)
	}
	fmt.Fprintf(file, "0")
}
