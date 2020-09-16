/*
	Параллельная обработка
	По данным n процессорам и m задач определите, для каждой из задач,
	каким процессором она будет обработана.

	Формат входа. Первая строка входа содержит числа n и m. Вторая
	содержит числа t0, . . . , tm−1, где ti — время, необходимое на обработку i-й задачи.
	Считаем, что и процессоры, и задачи нумеруются с нуля.

	Формат выхода. Выход должен содержать ровно m строк: i-я (считая с нуля) строка должна
	содержать номер процессора, который получит i-ю задачу на обработку, и время,
	когда это произойдёт.
*/

package main

import (
	"bufio"
	"container/list"
	"fmt"
	"github.com/qypec/basic-data-structures/tree/master/heap"
	"os"
	"strconv"
)

type Processor struct {
	id        int
	taskId    int
	startTime int
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	m, _ := strconv.Atoi(scanner.Text())

	duration := make([]int, m)
	for i := 0; i < m; i++ {
		scanner.Scan()
		duration[i], _ = strconv.Atoi(scanner.Text())
	}

	working := heap.New(heap.MinHeap)
	resting := list.New()
	for i := 0; i < n && i < m; i++ {
		resting.PushBack(Processor{id: i})
	}

	time := 0
	for i := 0; ; {
		for e := resting.Front(); i < m && e != nil; e = resting.Front() {
			processor := e.Value.(Processor)
			processor.taskId = i
			processor.startTime = time
			fmt.Printf("%v %v\n", processor.id, processor.startTime)
			working.Insert(processor.startTime+duration[processor.taskId], processor)
			resting.Remove(resting.Front())
			i++
		}
		if i >= m {
			break
		}

		time = working.Front().Value.(Processor).startTime + duration[working.Front().Value.(Processor).taskId]
		for e := working.Front(); e != nil; e = working.Front() {
			if time != e.Value.(Processor).startTime+duration[e.Value.(Processor).taskId] {
				break
			}
			resting.PushBack(Processor{id: e.Value.(Processor).id})
			working.ExtractMin()
		}
	}
}
