/*
	1. Реализуйте функцию Walk.
	2. Протестируйте функцию Walk.

	Функция tree.New(k) создает структурированное случайным образом двоичное дерево,
	которое содержит значения k, 2k, 3k, ..., 10k.
	Создайте новый канал ch и запустите обход дерева: go Walk(tree.New(1), ch)
	Затем получите и напечатайте 10 значений из канала.
	Это должны быть цифры 1, 2, 3, ..., 10.

	3. Реализуйте функцию Same, используя Walk для определения, хранят ли t1 и t2 одинаковые значения.
	4. Протестируйте функцию Same.
	
	Same(tree.New(1), tree.New(1)) должна вернуть true, а Same(tree.New(1), tree.New(2)) - false.
*/

package main

import (
	"fmt"
	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	ch <- t.Value
	walk(t.Right, ch)
	walk(t.Left, ch)
}

func Walk(t *tree.Tree, ch chan int) {
	walk(t, ch)
	close(ch)
}

func search(arr []int, toFind int, quit chan bool) {
	i := 0
	for _, elem := range arr {
		if elem == toFind {
			break
		}
		i++
	}
	if i == len(arr) {
		quit <- false
	} else {
		quit <- true
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	values1 := make([]int, 0)

	go Walk(t1, ch1)
	go Walk(t2, ch2)
	for i := range ch1 {
		values1 = append(values1, i)
	}

	quit := make(chan bool, len(values1))
	for elem := range ch2 {
		go search(values1, elem, quit)
	}
	for i := 0; i < len(values1); i++ { // search waiting
		if ok := <-quit; !ok {
			return false
		}
	}
	return true
}

func main() {
	fmt.Println(Same(tree.New(7), tree.New(7)))
	fmt.Println(Same(tree.New(2), tree.New(3)))
}
