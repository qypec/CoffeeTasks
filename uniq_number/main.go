package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

// Uniq finds a unique integer among couples
// tests -> github.com/qypec/coffee-tasks/tree/master/uniq_number
func Uniq(in io.Reader) int {
	scanner := bufio.NewScanner(in) // MaxScanTokenSize = 64 * 1024
	scanner.Split(bufio.ScanWords)

	sum := 0
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		sum ^= num
	}
	return sum
}

func main() {
	fmt.Println(Uniq(os.Stdin))
}
