/*
	Реализовать обработчик сетевых пакетов.

	У компьютера, обрабатывающего пакеты, имеется сетевой буфер
	размера size. До начала обработки пакеты хранятся в буфере.
	Если буфер полностью заполнен в момент поступления пакета (есть size пакетов,
	поступивших ранее, которые до сих пор не обработаны), этот пакет
	отбрасывается и уже не будет обработан. Если несколько пакетов поступает
	в одно и то же время, они все будут сперва сохранены в буфер.

	Вход. Размер буфера size и число пакетов n, а также
		две последовательности arrival1, . . . , arrivaln и
		duration1, . . . , durationn, обозначающих время поступления
		и длительность обработки n пакетов.
	Выход. Для каждого из данных n пакетов необходимо
		вывести время начала его обработки или −1, если пакет
		не был обработан (это происходит в случае, когда пакет
		поступает в момент, когда в буфере компьютера уже
		находится size пакетов).
*/

package main

import (
	"bufio"
	"container/list"
	"fmt"
	"os"
	"strconv"
)

type buffer struct {
	queue *list.List
	size int
}

func (b *buffer) isFull() bool {
	if b.queue.Len() == b.size {
		return true
	}
	return false
}

func getPackage(time int, arrival []int, counter int) bool {
	if counter != len(arrival) && time >= arrival[counter] {
		return true
	}
	return false
}

func packageProcess(bufferSize int, arrival, duration []int) []int {
	if len(arrival) == 0 {
		return nil
	}

	processStartTime := make([]int, len(arrival))
	for i, _ := range processStartTime {
		processStartTime[i] = -1
	}

	buffer := &buffer{list.New(), bufferSize}
	time := arrival[0]
	for counter := 0; ; {
		if received := getPackage(time, arrival, counter); received {
			if buffer.isFull() {
				processStartTime[counter] = -1
			} else {
				buffer.queue.PushBack(counter)
			}
			counter++
		}

		if packet := buffer.queue.Front(); packet != nil {
			processingPackage := packet.Value.(int)
			if processStartTime[processingPackage] == -1 {
				processStartTime[processingPackage] = time
			}

			currentProcessTime := processStartTime[processingPackage] + duration[processingPackage]
			if counter != len(arrival) && currentProcessTime > arrival[counter] {
				time = arrival[counter]
			} else {
				time = currentProcessTime
				buffer.queue.Remove(buffer.queue.Front())
			}
		}
		if counter != len(arrival) && buffer.queue.Len() == 0 && time < arrival[counter] {
			time = arrival[counter]
		}
		if counter == len(arrival) && buffer.queue.Len() == 0 {
			break
		}
	}
	return processStartTime
}

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Split(bufio.ScanWords)

	scanner.Scan()
	bufferSize, _ := strconv.Atoi(scanner.Text())
	scanner.Scan()
	n, _ := strconv.Atoi(scanner.Text())

	arrival := make([]int, n)
	duration := make([]int, n)
	for i := 0; i < n; i++ {
		scanner.Scan()
		arrival[i], _ = strconv.Atoi(scanner.Text())
		scanner.Scan()
		duration[i], _ = strconv.Atoi(scanner.Text())
	}

	time := packageProcess(bufferSize, arrival, duration)
	for _, t := range time {
		fmt.Println(t)
	}
}
