package main

import (
	"bufio"
	"os"
	// "io"
	"strconv"
	"fmt"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	sum := 0;
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		sum ^= num;
	}
	fmt.Println(sum)
}
