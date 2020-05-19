package main

import (
	"github.com/stretchr/testify/require"
	"testing"
)

type seqSumTest struct {
	target int
	seq    []int
}

func TestSeqSum(t *testing.T) {
	testCases := []seqSumTest{
		/* 1 */ {5, []int{1, 7, 3, 4, 7, 9}},
		/* 2 */ {1, []int{1}},
		/* 3 */ {999999998, []int{999999997, 1}},
		/* 4 */ {2, []int{1, 1, 1, 1, 1, 1, 1}},
		/* 5 */ {3, []int{1, 1, 1, 1, 1, 1, 1}},
		/* 6 */ {0, []int{9, 7, 10, 0, 5, 5, 0}},
		/* 7 */ {1, []int{9, 7, 1, 2, 5, 5, 2, 0, 3, 5}},
		/* 8 */ {5, []int{9, 7, 1, 2, 5, 6, 2, 0, 3, 2, 4, 3, 3, 1, 0, 5, 5}},
		/* 9 */ {99, []int{9, 7, 1, 2, 5, 6, 2, 0, 3, 2, 4, 3, 3, 1, 0, 5, 5}},
		/* 10 */ {0, []int{0}},
		/* 11 */ {0, []int{0, 0}},
		/* 12 */ {0, []int{0, 0, 0}},
		/* 13 */ {22, []int{1, 3, 4, 12, 23, 24, 6, 3, 7, 10}},
		/* 14 */ {6, []int{1, 7, 3, 4, 7, 9}},
	}
	expectedResult := []int{
		/* 1 */ 1,
		/* 2 */ 0,
		/* 3 */ 1,
		/* 4 */ 1,
		/* 5 */ 0,
		/* 6 */ 1,
		/* 7 */ 1,
		/* 8 */ 1,
		/* 9 */ 0,
		/* 10 */ 0,
		/* 11 */ 1,
		/* 12 */ 1,
		/* 13 */ 1,
		/* 14 */ 0,
	}

	for i, testCase := range testCases {
		require.Equal(t, expectedResult[i], SeqSum(testCase.target, testCase.seq), "test number %v\ntarget %v\nseq %v\n", i+1, testCase.target, testCase.seq)
	}
}
