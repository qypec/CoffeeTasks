package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

const inputFileName = "input.txt"
const outputFileName = "output.txt"

// my implementation of caching binary search
type cacheT struct {
	toFind int
	m      int
}

// caching remembers values until the next binarySearch call
func (prev *cacheT) caching(toFind, m int) {
	prev.toFind = toFind
	prev.m = m
}

// prev.m - индекс, на котором завершился binarySearch в прошлый раз
// toFind (следующее по очереди число) <= prev.toFind, т.к. toFind = target - num (см. функцию SeqSum), a num увеличивается с каждой итерацией
// отсюда следует, что toFind может находится в пределах одной ячейки массива от prev.m
// эти условия и проверяет checkCache
func (prev *cacheT) checkCache(a []int, toFind, l, r int) (bool, int) {
	if prev.m < l || prev.m > r {
		return false, -1
	} // prev.m out of range
	if prev.toFind == toFind {
		return true, -1
	} // If the same `toFind`, then the last time it was not found
	if a[prev.m] == toFind {
		return true, prev.m
	}
	if a[prev.m-1] == toFind {
		return true, prev.m - 1
	}
	if a[prev.m-1] < toFind && a[prev.m] > toFind {
		return true, -1
	}
	return false, -1
}

func ascendingMoves(a []int, m int, toFind int, fix, l, r *int) {
	if a[m] >= toFind {
		if a[m] == toFind {
			*fix = m
		}
		*r = m - 1
	} else {
		*l = m + 1
	}
}

/* BinarySearch finds the element that is equal than the given (toFind) */
/* array a must be sorted */
/* Complexity: log(n) */
func binarySearch(a []int, toFind int, l, r int, prev cacheT) int {
	if use, fix := prev.checkCache(a, toFind, l, r); use {
		return fix
	}

	fix := -1
	m := 0
	for l <= r {
		m = l + int((r-l)/2)
		ascendingMoves(a, m, toFind, &fix, &l, &r)
		if fix != -1 {
			break
		}
	}
	prev.caching(toFind, m)
	return fix
}

// my implementation of quick sort (github.com/qypec/basic-algorithms/blob/master/quick_sort/)
func partition(arr []int, l, r int) int {
	pivot := arr[l]
	j := l
	for i := l + 1; i <= r; i++ {
		if arr[i] <= pivot {
			j++
			arr[j], arr[i] = arr[i], arr[j]
		}
	}
	arr[l], arr[j] = arr[j], arr[l]
	return j
}

func recQuickSort(arr []int, l, r int) {
	if l < 0 || r < 0 {
		return
	}
	if l >= r {
		return
	}
	m := partition(arr, l, r)
	recQuickSort(arr, l, m-1)
	recQuickSort(arr, m+1, r)
}

// QuickSort sorts an array
// Complexity: nlog(n)
func quickSort(arr []int) { recQuickSort(arr, 0, len(arr)-1) }

// SeqSum searches in a seq for two numbers that in total give a target.
// If successful, it will return 1, otherwise it will return 0
// Complexity: nlog(n)
// tests -> github.com/qypec/coffee-tasks/tree/master/seq_sum
func SeqSum(target int, seq []int) int {
	quickSort(seq)
	cache := cacheT{-1, -1}

	// Searches for a number equal to `target - num` in a sequence
	// nlog(n)
	for i, num := range seq {
		if num > target/2 {
			break
		}
		fix := binarySearch(seq, target-num, i+1, len(seq)-1, cache)
		if fix != -1 && fix != i {
			return 1
		}
	}
	return 0
}

func main() {
	// openning input file
	input, err := os.Open(inputFileName)
	if err != nil {
		panic(err)
	}
	defer input.Close()
	scanner := bufio.NewScanner(input)
	scanner.Split(bufio.ScanWords)

	// reading sequence
	scanner.Scan()
	target, _ := strconv.Atoi(scanner.Text())
	seq := make([]int, 0)
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		if num <= target {
			seq = append(seq, num)
		}
	}

	res := SeqSum(target, seq)

	// writing the response to the output file
	output, err := os.Create(outputFileName)
	if err != nil {
		panic(err)
	}
	defer output.Close()
	fmt.Fprintf(output, "%v", res)
}
