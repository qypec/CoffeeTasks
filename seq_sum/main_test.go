package main

import (
	// "fmt"
	"bytes"
	"strings"
	"testing"
	"os"
	"github.com/stretchr/testify/require"
)

func TestSeqSumBasic(t *testing.T) {
	testCases := []string{
		/* 1 */ "5\n1 3 4",
		/* 2 */ "1\n1",
		/* 3 */ "999999998\n999999997 1",
		/* 4 */ "2\n1 1 1 1 1 1 1",
		/* 5 */ "3\n1 1 1 1 1 1 1 2",
		/* 6 */ "0\n0 0",
		/* 7 */ "1\n1 0",
		/* 8 */ "5\n1 2 5 2 0 3 42 2 4 12 3 3 1 0 5 5 8",
		/* 9 */ "99\n9 7 1 2 5 6 2 0 3 2 4 3 3 1 0 5 5",
		/* 10 */ "0\n0",
		/* 11 */ "0\n0 1 3 0 2",
		/* 12 */ "22\n1 3 4 12 6 3 7 10 27 12",
		/* 13 */ "6\n1 4 3",
		/* 14 */ "25\n1 3 3 4 6 7 10 11 11 11 20",
	}
	expectedResult := []string{
		/* 1 */ "1",
		/* 2 */ "0",
		/* 3 */ "1",
		/* 4 */ "1",
		/* 5 */ "1",
		/* 6 */ "1",
		/* 7 */ "1",
		/* 8 */ "1",
		/* 9 */ "0",
		/* 10 */ "0",
		/* 11 */ "1",
		/* 12 */ "1",
		/* 13 */ "0",
		/* 14 */ "0",
	}

	for i, testCase := range testCases {
		in := strings.NewReader(testCase)
		out := new(bytes.Buffer)
		SeqSum(in, out)
		require.Equal(t, expectedResult[i], out.String(), "test number %v\n", i+1)
	}
}

func BenchmarkSeqSumUpgradeTest01(b *testing.B) {
	in, err := os.Open("testdata/test_01.txt")
	if err != nil {
		panic(err)
	}
	defer in.Close()

	out := new(bytes.Buffer)	
	for i := 0; i < b.N; i++ {
		SeqSum(in, out)
	}
}

func BenchmarkSeqSumUpgradeTest02(b *testing.B) {
	in, err := os.Open("testdata/test_02.txt")
	if err != nil {
		panic(err)
	}
	defer in.Close()

	out := new(bytes.Buffer)
	for i := 0; i < b.N; i++ {
		SeqSum(in, out)
	}
}

func BenchmarkSeqSumUpgradeTest03(b *testing.B) {
	in, err := os.Open("testdata/test_03.txt")
	if err != nil {
		panic(err)
	}
	defer in.Close()

	out := new(bytes.Buffer)	
	for i := 0; i < b.N; i++ {
		SeqSum(in, out)
	}
}

func BenchmarkSeqSumTest01(b *testing.B) {
	in, err := os.Open("testdata/test_01.txt")
	if err != nil {
		panic(err)
	}
	defer in.Close()

	out := new(bytes.Buffer)	
	for i := 0; i < b.N; i++ {
		SeqSumPrev(in, out)
	}
}

func BenchmarkSeqSumTest02(b *testing.B) {
	in, err := os.Open("testdata/test_02.txt")
	if err != nil {
		panic(err)
	}
	defer in.Close()

	out := new(bytes.Buffer)
	for i := 0; i < b.N; i++ {
		SeqSumPrev(in, out)
	}
}

func BenchmarkSeqSumTest03(b *testing.B) {
	in, err := os.Open("testdata/test_03.txt")
	if err != nil {
		panic(err)
	}
	defer in.Close()

	out := new(bytes.Buffer)	
	for i := 0; i < b.N; i++ {
		SeqSumPrev(in, out)
	}
}

func BenchmarkSeqSumMapTest01(b *testing.B) {
	in, err := os.Open("testdata/test_01.txt")
	if err != nil {
		panic(err)
	}
	defer in.Close()

	out := new(bytes.Buffer)	
	for i := 0; i < b.N; i++ {
		SeqSumMap(in, out)
	}
}

func BenchmarkSeqSumMapTest02(b *testing.B) {
	in, err := os.Open("testdata/test_02.txt")
	if err != nil {
		panic(err)
	}
	defer in.Close()

	out := new(bytes.Buffer)
	for i := 0; i < b.N; i++ {
		SeqSumMap(in, out)
	}
}

func BenchmarkSeqSumMapTest03(b *testing.B) {
	in, err := os.Open("testdata/test_03.txt")
	if err != nil {
		panic(err)
	}
	defer in.Close()

	out := new(bytes.Buffer)	
	for i := 0; i < b.N; i++ {
		SeqSumMap(in, out)
	}
}
