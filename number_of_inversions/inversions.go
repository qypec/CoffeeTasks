/*
	Первая строка содержит число n, вторая — массив A[1…n], содержащий натуральные числа.
	Необходимо посчитать число пар индексов 1 < i < j <= n, для которых A[i]>A[j].
	(Такая пара элементов называется инверсией массива)
*/

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

var invertionsCounter int

func merge(first [][]int, second [][]int) [][]int {
	mergeSlice := make([][]int, 1)
	for len(first[0]) != 0 && len(second[0]) != 0 {
		if first[0][0] > second[0][0] {
			invertionsCounter += len(first[0])
			mergeSlice[0] = append(mergeSlice[0], second[0][0])
			second[0] = second[0][1:]
		} else {
			mergeSlice[0] = append(mergeSlice[0], first[0][0])
			first[0] = first[0][1:]
		}
	}
	mergeSlice[0] = append(mergeSlice[0], first[0]...)
	mergeSlice[0] = append(mergeSlice[0], second[0]...)
	return mergeSlice
}

func mergeSort(queue [][]int) [][]int {
	if len(queue) == 1 {
		return queue
	}
	m := int(math.Ceil(float64(len(queue) / 2)))
	first := queue[:m]
	second := queue[m:]
	return merge(mergeSort(first), mergeSort(second))
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()

	queue := make([][]int, 0)
	for scanner.Scan() {
		newArray := make([]int, 0)
		elem, _ := strconv.Atoi(scanner.Text())
		newArray = append(newArray, elem)
		queue = append(queue, newArray)
	}

	queue = mergeSort(queue)
	fmt.Println(invertionsCounter)
}
