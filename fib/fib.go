package main

import (
	"fmt"
)

func fib(num int) int {
	arr := []int{0, 1}
	
	for i := 2; i <= num; i++ {
		arr = append(arr, arr[i - 1] + arr[i - 2])
	}
	return arr[num]
}

func main() {
	fmt.Println("1 = ", fib(1))
	fmt.Println("2 = ", fib(2))
	fmt.Println("3 = ", fib(0))
	fmt.Println("4 = ", fib(10))
	fmt.Println("5 = ", fib(20))
	fmt.Println("6 = ", fib(50))
}
