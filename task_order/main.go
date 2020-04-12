/*
	В первой строке вам дается число задач N. В следующей строке идет описание задач,
	для каждой задачи вам дана ее продолжительность - Ti (где i - это номер задачи, от 0 до N−1 не включительно).
	На выход вам требуется вывести номера задач (задачи нумеруются с 0) в порядке,
	который минимизирует среднее время ожидания завершения задачи.
 */

package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type task struct {
	time int
	index int
}

type ByTime []task

func (a ByTime) Len() int           { return len(a) }
func (a ByTime) Less(i, j int) bool { return a[i].time < a[j].time }
func (a ByTime) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	t := make([]task, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		t[i].time, _ = strconv.Atoi(scanner.Text())
		t[i].index = i
	}

	sort.Sort(ByTime(t))
	for i := 0; i < n; i++ {
		fmt.Printf("%v ", t[i].index)
	}
}
