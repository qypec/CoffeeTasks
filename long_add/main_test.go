package main

import (
	"math/rand"
	"testing"
	"github.com/stretchr/testify/require"
	"math/big"
	"strconv"
	// "fmt"
)

type test struct {
	a []byte
	b []byte
}

func TestLongAddBasic(t *testing.T) {
	testCases := []test{
		/* 1 */ test{
			[]byte{'1', '2', '3', '4'},
			[]byte{'1', '2', '3', '4'},
		},
		/* 2 */ test{
			[]byte{'0'},
			[]byte{'1', '2', '3', '4'},
		},
		/* 3 */ test{
			[]byte{'9', '9', '9'},
			[]byte{'1'},
		},
		/* 4 */ test{
			[]byte{'9', '9', '0', '1'},
			[]byte{'1', '2', '3', '4', '5', '6', '7', '0'},
		},
	}
	testExpected := []string {
		/* 1 */ "2468",
		/* 2 */ "1234",
		/* 3 */ "1000",
		/* 4 */ "12355571",
	}

	for i, testCase := range testCases {
		require.Equal(t, testExpected[i], LongAdd(testCase.a, testCase.b), "test %v", i)
	}
}

func generator(size int) string {
	longNum := ""

	longNum += strconv.Itoa(rand.Intn(8) + 1)
	for i := 0; i < size - 1; i++ {
		longNum += strconv.Itoa(rand.Intn(9))
	}
	return longNum
}

func TestLongAddLONG(t *testing.T) {
	seedsNumber := 10
	testsNumber := 1000

	for seed := 1; seed <= seedsNumber; seed++ {
		rand.Seed(int64(seed))
		for i := 0; i < testsNumber; i++ {
			a := generator(rand.Intn(1000))
			b := generator(rand.Intn(1000))

			aBig := new(big.Int)
			aBig.SetString(a, 10)
			bBig := new(big.Int)
			bBig.SetString(b, 10)
			aBig.Add(aBig, bBig)

			require.Equal(t, aBig.String(), LongAdd([]byte(a), []byte(b)), "test %v", i)
		}
	}
}

func BenchmarkLongAdd(b *testing.B) {
	ar := []byte(generator(1000))
	br := []byte(generator(1000))
	for i := 0; i < b.N; i++ {
		LongAdd(ar, br)
	}
}

func BenchmarkBigAdd(b *testing.B) {
	ar := generator(1000)
	br := generator(1000)
	for i := 0; i < b.N; i++ {
		aBig := new(big.Int)
		aBig.SetString(ar, 10)
		bBig := new(big.Int)
		bBig.SetString(br, 10)
		aBig.Add(aBig, bBig)
	}
}