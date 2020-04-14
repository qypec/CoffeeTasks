package main

import (
	"bufio"
	"fmt"
	"github.com/AlgoAndStruct/myContainers/sort"
	"os"
	"strconv"
)

func customBound(a []int, toFind int, l, r int) int {
	//if isDescending(a) {
	//	if a[0] > toFind {
	//		return 0
	//	} else {
	//		return -1
	//	}
	//}

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

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

/* scanning n and m */
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	m, _ := strconv.Atoi(scanner.Text())

/* scanning segments */
	segmL := make([]int, 0)
	segmR := make([]int, 0)
	for i := 0; i < n; i++ {
		scanner.Scan()
		a, _ := strconv.Atoi(scanner.Text())
		segmL = append(segmL, a)
		scanner.Scan()
		b, _ := strconv.Atoi(scanner.Text())
		segmR = append(segmR, b)
	}

/* scanning dots */
	dots := make([]int, 0)
	for i := 0; i < m; i++ {
		scanner.Scan()
		x, _ := strconv.Atoi(scanner.Text())
		dots = append(dots, x)
	}

	sort.QuickSort3(segmL)
	sort.QuickSort3(segmR)
	for _, x := range dots {
		countSegmL := sort.UpperBound(segmL, x, 0, len(segmL) - 1)
		countSegmR := customBound(segmR, x, 0, len(segmR) - 1)
		if countSegmL == -1 { countSegmL = len(segmL) }
		if countSegmR == -1 { countSegmR = len(segmR) }
		num := countSegmL - countSegmR

		fmt.Print(num)
		fmt.Print(" ")
	}
	fmt.Print("\n")
}
