/*
**	Даны два числа неотрицательных числа A и B(числа могут содержать до 1000 цифр).
**	Вам нужно вычислить их сумму.
 */

package main

import (
	"bufio"
	"fmt"
	"os"
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
		return int(a[len(a)-i-1]) - ascii0
	}
	return 0
}

// BigAdd adds two big numbers
// test -> github.com/qypec/coffee-tasks/tree/master/long_add
func BigAdd(a, b []byte) string {
	carry := 0
	res := make([]rune, max(len(a), len(b))+1)
	for i := 0; i < max(len(a), len(b)); i++ {
		aInt := toInt(a, i)
		bInt := toInt(b, i)
		res[len(res)-i-1] = rune((aInt+bInt+carry)%10 + ascii0)
		carry = (aInt + bInt + carry) / 10
	}
	if carry != 0 {
		res[0] = rune(carry + ascii0)
	} else {
		res = res[1:]
	}
	return string(res)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	a := scanner.Bytes()
	scanner.Scan()
	b := scanner.Bytes()

	res := BigAdd(a, b)
	fmt.Println(res)
}
