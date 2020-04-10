/**
 *
 *	Ваша цель в данной задаче — реализовать cхему хеширования цепочками,
 *	используя таблицу с m ячейками и полиномиальной хеш-функцией на строках.
 *	Ваша программа должна поддерживать следующие
 *	типы запросов:
 *		• add string: добавить строку string в таблицу. Если такая
 *		строка уже есть, проигнорировать запрос;
 *		• del string: удалить строку string из таблицы. Если такой
 *		строки нет, проигнорировать запрос;
 *		• find string: вывести «yes» или «no» в зависимости от того,
 *		есть в таблице строка string или нет;
 *		• check i: вывести i-й список (используя пробел в качестве разделителя);
 *		если i-й список пуст, вывести пустую строку.
 *
 *	Все строки имеют длину от одного до пятнадцати и содержат только буквы латинского
 *	алфавита.
 *
**/

package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strconv"
)

type HashTable struct {
	arr  []*list.List
	size int // number of elements
	x    []int
}

const x = 263
const p = 1000000007
const strSize = 15

/* All strings have a length of one to fifteen and contain only letters of the Latin alphabet. */
func (h *HashTable) polynomialHashFunction(str string) int {
	hash := 0
	for i, ch := range str {
		hash += int(ch) * h.x[i]
	}
	return (hash % p) % h.size
}

func (h *HashTable) Init(m int) {
	h.arr = make([]*list.List, m)
	h.size = m

	h.x = make([]int, strSize)
	h.x[0] = 1
	for i := 1; i < strSize; i++ {
		h.x[i] = (x * h.x[i-1]) % p
	}
}

func (h *HashTable) Add(str string) {
	hash := h.polynomialHashFunction(str)
	if h.arr[hash] == nil {
		h.arr[hash] = list.New()
	}
	for e := h.arr[hash].Front(); e != nil; e = e.Next() {
		if e.Value.(string) == str {
			return
		}
	}
	h.arr[hash].PushFront(str)
}

func (h *HashTable) Delete(str string) {
	hash := h.polynomialHashFunction(str)
	if h.arr[hash] == nil {
		return
	}
	for e := h.arr[hash].Front(); e != nil; e = e.Next() {
		if e.Value.(string) == str {
			h.arr[hash].Remove(e)
		}
	}
}

func (h *HashTable) Find(str string) bool {
	hash := h.polynomialHashFunction(str)
	if h.arr[hash] == nil {
		return false
	}
	for e := h.arr[hash].Front(); e != nil; e = e.Next() {
		if e.Value.(string) == str {
			return true
		}
	}
	return false
}

func main() {
	var table HashTable

	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	m, _ := strconv.Atoi(scanner.Text())
	scanner.Scan() // skip n

	table.Init(m)
	for scanner.Scan() {
		switch scanner.Text() {
		case "add":
			scanner.Scan()
			table.Add(scanner.Text())
		case "find":
			scanner.Scan()
			if table.Find(scanner.Text()) {
				fmt.Println("yes")
			} else {
				fmt.Println("no")
			}
		case "del":
			scanner.Scan()
			table.Delete(scanner.Text())
		case "check":
			scanner.Scan()
			i, _ := strconv.Atoi(scanner.Text())
			if table.arr[i] == nil {
				fmt.Print("\n")
			} else {
				for e := table.arr[i].Front(); e != nil; e = e.Next() {
					fmt.Printf("%s ", e.Value.(string))
				}
				fmt.Print("\n")
			}
		}
	}
}
