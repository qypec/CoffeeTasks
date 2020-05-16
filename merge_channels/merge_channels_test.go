package main

import (
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func workerMillisecond(x int) int {
	time.Sleep(time.Millisecond)
	return x * 10
}

func workerSecond(x int) int {
	time.Sleep(time.Second)
	return x * 10
}

func TestMerge2Channels(t *testing.T) {
	testsNumber := 1000

	in1 := make(chan int, testsNumber)
	in2 := make(chan int, testsNumber)
	out := make(chan int, testsNumber)
	defer close(in1)
	defer close(in2)
	defer close(out)

	for i := 0; i < testsNumber; i++ {
		in1 <- i
		in2 <- i
	}

	// tests for non-blocking
	start := time.Now()
	Merge2Channels(workerMillisecond, in1, in2, out, testsNumber)
	end := int64(time.Since(start))

	expectedTime := int64(time.Millisecond)
	require.GreaterOrEqual(t, expectedTime, end, "blocking Merge2Channels")

	// waiting for the completion goroutines to avoid datarace
	for i := 0; i < testsNumber; i++ {
		res := <-out
		require.Equal(t, (i*10 + i*10), res)
	}

	// unbuffered cahnnels
	in1 = make(chan int)
	in2 = make(chan int)
	out = make(chan int)

	start = time.Now()
	Merge2Channels(workerMillisecond, in1, in2, out, testsNumber)
	end = int64(time.Since(start))

	expectedTime = int64(time.Millisecond)
	require.GreaterOrEqual(t, expectedTime, end, "blocking Merge2Channels unbuffered")

	for i := 0; i < testsNumber; i++ {
		in1 <- i
		in2 <- i
		_ = <-out
	}
}

func BenchmarkMerge2Channels(b *testing.B) {
	n := 100

	in1 := make(chan int, n)
	in2 := make(chan int, n)
	out := make(chan int, n)
	defer close(in1)
	defer close(in2)
	defer close(out)

	for i := 0; i < b.N; i++ {
		Merge2Channels(workerMillisecond, in1, in2, out, n)

		for j := 0; j < n; j++ {
			in1 <- j
			in2 <- j
			_ = <-out
		}
	}
}
