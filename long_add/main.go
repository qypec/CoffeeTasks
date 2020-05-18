package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const ascii0 = int('0') // 48

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func toInt(a []byte, i int) int {
	if i < len(a) {
		return int(a[len(a) - i - 1]) - ascii0
	}
	return 0
}

// const digits = "0123456789"

// func smallItoa(n int) string {
// 	return digits[n : n+1]
// }

// func LongAddv2(a, b []byte) string {
// 	var res string

// 	carry := 0
// 	for i := 0; i < max(len(a), len(b)); i++ {
// 		aInt := toInt(a, i)
// 		bInt := toInt(b, i)
// 		res = smallItoa((aInt + bInt + carry) % 10) + res;
// 		carry = (aInt + bInt + carry) / 10
// 	}
// 	if carry != 0 {
// 		res = strconv.Itoa(carry) + res;
// 	}
// 	return res
// }

func LongAdd(a, b []byte) string {
	var res string

	carry := 0
	for i := 0; i < max(len(a), len(b)); i++ {
		aInt := toInt(a, i)
		bInt := toInt(b, i)
		res = strconv.Itoa((aInt + bInt + carry) % 10) + res;
		carry = (aInt + bInt + carry) / 10
	}
	if carry != 0 {
		res = strconv.Itoa(carry) + res;
	}
	return res
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	a := scanner.Bytes()
	scanner.Scan()
	b := scanner.Bytes()

	res := LongAdd(a, b)
	fmt.Println(res)
}