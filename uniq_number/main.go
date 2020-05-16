package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

func Uniq(in io.Reader) int {
	scanner := bufio.NewScanner(in)
	scanner.Split(bufio.ScanWords)

	sum := 0
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		sum ^= num;
	}
	return sum
}

func main() {
	fmt.Println(Uniq(os.Stdin))
}
