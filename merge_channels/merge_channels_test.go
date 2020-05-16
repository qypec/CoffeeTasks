package main

import (

	"testing"
	"github.com/stretchr/testify/require"
	"time"
	// "sync"
)

func worker(x int) int {
	time.Sleep(1 * time.Millisecond)
	return x * 10
}

const blockingTestNumbers = 1000

func TestMergeChannels(t *testing.T) {
	ch1 := make(chan int, blockingTestNumbers)
	ch2 := make(chan int, blockingTestNumbers)
	out := make(chan int, blockingTestNumbers)
	defer close(ch1)
	defer close(ch2)
	defer close(out)

	for i := 0; i < blockingTestNumbers; i++ {
		ch1 <- i
		ch2 <- i
	}

	start := time.Now()
	Merge2Channels(worker, ch1, ch2, out, blockingTestNumbers)
	end := int64(time.Since(start))

	expectedTime := int64(time.Millisecond)
	require.GreaterOrEqual(t, expectedTime, end, "blocking Merge2Channels")

	for i := 0; i < blockingTestNumbers; i++ {
		_ = <-out
	}
}