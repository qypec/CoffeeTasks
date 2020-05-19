package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"sort"
)

const inputFileName = "input.txt"
const outputFileName = "output.txt"

// my implementation of upper bound (github.com/qypec/basic-algorithms/tree/master/upper_bound)
func isDescending(a []int) bool {
	if a[0] >= a[len(a)-1] {
		return true
	}
	return false
}

/* UpperBound finds the first element that is strictly larger than the given (toFind) */
/* array a must be sorted */
/* log(n) */
func upperBound(a []int, toFind int, l, r int) int {
	if isDescending(a) {
		if a[0] > toFind {
			return 0
		} else {
			return -1
		}
	}

	fix := -1
	for l <= r {
		m := l + int((r-l)/2)
		if a[m] > toFind {
			fix = m
			r = m - 1
		} else {
			l = m + 1
		}
	}
	return fix
}

// my implementation of binary search (github.com/qypec/basic-algorithms/tree/master/binary_search)
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

func descendingMoves(a []int, m int, toFind int, fix, l, r *int) {
	if a[m] > toFind {
		*l = m + 1
	} else {
		if a[m] == toFind {
			*fix = m
		}
		*r = m - 1
	}
}

/* BinarySearch finds the first element that is equal than the given (toFind) */
/* array a must be sorted */
/* Complexity: log(n) */
func binarySearch(a []int, toFind int, l, r int) int {
	moves := ascendingMoves
	if a[0] > a[len(a)-1] {
		moves = descendingMoves
	}

	fix := -1
	for l <= r {
		m := l + int((r-l)/2)
		moves(a, m, toFind, &fix, &l, &r)
	}
	return fix
}

// SeqSum searches in a seq for two numbers that in total give a target.
// If successful, it will return 1, otherwise it will return 0
// Complexity: nlog(n)
// tests -> !!!!
func SeqSum(target int, seq []int) int {
	sort.Slice(seq, func(i, j int) bool { // nlog(n)
			return seq[i] < seq[j]
	})

	// Removes from the sequence numbers that are larger than the target. 
	// These numbers cannot have a pair.
	// log(n)
	if largerTarget := upperBound(seq, target, 0, len(seq) - 1); largerTarget != -1 {
		seq = seq[0:largerTarget]
	}

	// Searches for a number equal to `target - num` in a sequence
	// nlog(n)
	for _, num := range seq {
		if binarySearch(seq, target - num, 0, len(seq) - 1) != -1 {
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
		seq = append(seq, num)
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
