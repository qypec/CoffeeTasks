/* 
	Допустим у нас есть функция loop(), которая выполняется 1-2000ms.
	Напишите такой код, который будет вызывать эту функцию строго раз в секунду. Цикл должен выполняться, 
	пока не придет context.Done. Если время выполнения loop больше одной секунды, то следующий вызов игнорируется.

	Задача со звездочкой: нельзя использовать time.Sleep.
	Задача с другой звездочкой: можно использовать time.Sleep, но нельзя создавать горутины.
*/

package main

import (
	"context"
	"fmt"
	"math/rand"
	"time"

	"go.uber.org/atomic"
)

func main() {
	ctx, cancel := context.WithCancel(context.Background())
	go func() {
		time.Sleep(1 * time.Minute)
		cancel()
	}()

	var loopIsFinished atomic.Bool
	loopIsFinished.Store(true)

	watcherCh := make(chan bool, 1)
	ticker := time.NewTicker(1 * time.Second)
	for {
		start := time.Now()
		fmt.Printf("Start at %v\n", start)

		select {
		case <-ticker.C:
			if loopIsFinished.Swap(false) {
				go func() {
					flag := <-watcherCh
					loopIsFinished.Store(flag)
				}()

				go func() {
					defer func() {
						watcherCh <- true
					}()

					loop()
				}()
			}
		case <-ctx.Done():
			fmt.Println("context.Done()")
			return
		}

		end := time.Now()
		fmt.Printf("End at %v\n", end)
		fmt.Printf("Step duration: %v\n\n", end.Sub(start))
	}
}

func loop() {
	randomNum := rand.Intn(2000)

	fmt.Println("Sleep period:", randomNum, "ms")

	time.Sleep(time.Millisecond * time.Duration(randomNum))
}
