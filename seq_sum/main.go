package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
)

const inputFileName = "input.txt"
const outputFileName = "output.txt"

func partition3(arr []int, l int, r int) (int, int) {
	if r-l <= 1 { // len(arr) = 2
		if arr[r] < arr[l] {
			arr[r], arr[l] = arr[l], arr[r]
		}
		return l, r
	}

	mid := l
	pivot := arr[r]
	for mid <= r {
		if arr[mid] < pivot {
			arr[l], arr[mid] = arr[mid], arr[l]
			l++
			mid++
		} else if arr[mid] == pivot {
			mid++
		} else if arr[mid] > pivot {
			arr[mid], arr[r] = arr[r], arr[mid]
			r--
		}
	}
	return l - 1, mid
}

func quickSort3(arr []int, l, r int) {
	if l < 0 || r < 0 {
		return
	}
	if l >= r {
		return
	}
	m1, m2 := partition3(arr, l, r)
	quickSort3(arr, l, m1)
	quickSort3(arr, m2, r)
}

func readSequence(in io.Reader) (int, []int) {
	scanner := bufio.NewScanner(in)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	target, _ := strconv.Atoi(scanner.Text())
	seq := make([]int, 0)
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		if num <= target {
			seq = append(seq, num)
		}
	}
	return target, seq
}

// SeqSum searches in a seq for two numbers that in total give a target.
// If successful, it will write to out 1, otherwise 0
// tests -> github.com/qypec/coffee-tasks/tree/master/seq_sum
func SeqSum(in io.Reader, out io.Writer) {
	target, seq := readSequence(in)
	quickSort3(seq, 0, len(seq)-1) // nlog(n)

	ans := 0
	lower := 0
	upper := len(seq) - 1
	for lower >= 0 && lower < upper {
		sum := seq[lower] + seq[upper]
		if sum == target {
			ans = 1
			break
		} else if sum < target {
			lower++
		} else {
			upper--
		}
	}
	fmt.Fprintf(out, "%v", ans)
}

func main() {
	input, err := os.Open(inputFileName)
	if err != nil {
		panic(err)
	}
	defer input.Close()
	output, err := os.Create(outputFileName)
	if err != nil {
		panic(err)
	}
	defer output.Close()

	SeqSum(input, output)
}

/* finds the first element that is larger or equal than the given(toFind) */
/* array a must be sorted */
/* log(n) */
func upperBound(a []int, toFind int, l, r int) int {
	fix := -1
	for l <= r {
		m := l + int((r-l)/2)
		if a[m] >= toFind {
			fix = m
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return fix
}

/* BinarySearch finds the first element that is equal than the given (toFind) */
/* array a must be sorted */
/* Complexity: log(n) */
func binarySearch(a []int, toFind int, l, r int) int {
	fix := -1
	for l <= r {
		m := l + int((r-l)/2)
		if a[m] >= toFind {
			if a[m] == toFind {
				fix = m
			}
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return fix
}

// // my implementation of quick sort (github.com/qypec/basic-algorithms/blob/master/quick_sort/)
// func partition(arr []int, l, r int) int {
// 	pivot := arr[l]
// 	j := l
// 	for i := l + 1; i <= r; i++ {
// 		if arr[i] <= pivot {
// 			j++
// 			arr[j], arr[i] = arr[i], arr[j]
// 		}
// 	}
// 	arr[l], arr[j] = arr[j], arr[l]
// 	return j
// }

// // quickSort sorts an array
// // Complexity: nlog(n)
// func quickSort(arr []int, l, r int) {
// 	if l < 0 || r < 0 {
// 		return
// 	}
// 	if l >= r {
// 		return
// 	}
// 	m := partition(arr, l, r)
// 	quickSort(arr, l, m-1)
// 	quickSort(arr, m+1, r)
// }

// SeqSum searches in a seq for two numbers that in total give a target.
// If successful, it will return 1, otherwise it will return 0
// Complexity: nlog(n)
// tests -> github.com/qypec/coffee-tasks/tree/master/seq_sum
func SeqSumPrev(in io.Reader, out io.Writer) {
	target, seq := readSequence(in)
	if len(seq) == 0 {
		fmt.Fprintf(out, "%v", 0)
		return
	}
	quickSort3(seq, 0, len(seq)-1)

	// Removes small numbers
	i := upperBound(seq, target-seq[len(seq)-1], 0, len(seq)-1)

	ans := 0
	// Searches for a number equal to `target - seq[i]` in a sequence
	for ; i >= 0 && i < len(seq); i++ {
		fix := binarySearch(seq, target-seq[i], i+1, len(seq)-1)
		if fix != -1 {
			ans = 1
			break
		}
	}
	fmt.Fprintf(out, "%v", ans)
}

func SeqSumMap(in io.Reader, out io.Writer) {
	scanner := bufio.NewScanner(in)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	target, _ := strconv.Atoi(scanner.Text())
	seq := make([]int, 0)
	for scanner.Scan() {
		num, _ := strconv.Atoi(scanner.Text())
		if num <= target {
			seq = append(seq, num)
		}
	}
	seqMap := make(map[int]int, len(seq))
	for _, item := range seq {
		seqMap[item]++
	}
	ans := 0
	for num := range seqMap {
		seqMap[num]--
		if seqMap[num] == 0 {
			delete(seqMap, num)
		}
		if _, ok := seqMap[target-num]; ok {
			ans = 1
			break
		}
	}
	fmt.Fprintf(out, "%v", ans)
}
