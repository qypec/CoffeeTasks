package main

import (
	"fmt"
	"math/rand"
	"os"
)

const limit = 1000000
const maxNum = 999999999

const filePath = "testdata/"

// фиксированное количество различных чисел
// пример: 1 3 2 1 3 1 1 2 3 3 3 2 2 1, алфавит состоит из 3 чисел
func test1() {
	file, err := os.Create(filePath + "test_01.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	alphabetSize := 1000 // количество разных чисел
	alphabet := make([]int, 0, alphabetSize)
	alphabetStep := int(maxNum / alphabetSize) // шаг, на который числа должны отличаться друг от друга
	prevValue := 0
	for i := 0; i < alphabetSize; i++ { // получаем фиксированное количество рандомных чисел (алфавит)
		alphabet = append(alphabet, prevValue + rand.Intn(alphabetStep))
		prevValue += alphabetStep
	}

	fmt.Fprintf(file, "%v\n", alphabet[len(alphabet) - 1]) // ищем сумму самого большого числа из алфавита
	fmt.Fprintf(file, "%v ", alphabet[len(alphabet) - 1])
	for i := 0; i < limit; i++ { // заполняем файл в рандомном порядке фиксированными числами
		fmt.Println("1 ", i)
		fmt.Fprintf(file, "%v ", alphabet[rand.Intn(alphabetSize - 1)])
	}
	fmt.Fprintf(file, "0")
}

// одинаковые числа maxNum, ищем сумму maxNum
func test2() {
	file, err := os.Create(filePath + "test_02.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fmt.Fprintf(file, "%v\n", maxNum)
	for i := 0; i < limit; i++ {
		fmt.Println("2 ", i)
		fmt.Fprintf(file, "%v ", maxNum)
	}
	fmt.Fprintf(file, "0")
}

// полностью случайные числа
func test3() {
	file, err := os.Create(filePath + "test_03.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	fmt.Fprintf(file, "%v\n", maxNum)
	for i := 0; i < limit; i++ {
		fmt.Println("3 ", i)
		fmt.Fprintf(file, "%v ", rand.Intn(maxNum - 1))
	}
	fmt.Fprintf(file, "0")
}

// go run testdata/generator.go
func main() {
	test1() // фиксированное количество различных чисел
	test2() // одинаковые числа
	test3() // полностью случайные числа
}
