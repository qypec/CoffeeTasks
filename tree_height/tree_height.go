/*
	Вычислить высоту данного дерева.

	Вход. Корневое дерево с вершинами {0, . . . , n−1}, заданное
		как последовательность parent0, . . . , parentn−1, где parenti — родитель i-й вершины.

	Выход. Высота дерева.
*/

package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func subTreeHeight(arr []int, Heights []float64, vertex int) float64 {
	height := float64(1)

	if arr[vertex] == -1 { // проверяем, не является ли родитель этой вершины корнем
		return height
	}
	if Heights[vertex] == 0 { // проверяем, не считали ли уже высоту для этой вершины
		Heights[vertex] = height + subTreeHeight(arr, Heights, arr[vertex])
	}
	return Heights[vertex]
}

func treeHeight(arr []int, N int) int {
	maxHeight := float64(1)

	Heights := make([]float64, N)
	for i := 0; i < N; i++ {
		maxHeight = math.Max(maxHeight, subTreeHeight(arr, Heights, i))
	}
	return int(maxHeight)
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	N, _ := strconv.Atoi(scanner.Text())

	arr := make([]int, N)
	for i := 0; scanner.Scan(); i++ {
		arr[i], _ = strconv.Atoi(scanner.Text())
	}
	fmt.Println(treeHeight(arr, N))
}
