package main

import (
	"bytes"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

const tenMinllionTestPath = "./testdata/10_000_000_test.txt"

func TestUniq(t *testing.T) {

	// Test 01 | 1.000.000 pairs
	{
		file, err := os.Open(tenMinllionTestPath)
		if err != nil {
			panic(err)
		}

		fileContent, err := ioutil.ReadAll(file)
		if err != nil {
			panic(err)
		}
		file.Close()

		require.Equal(t, 0, Uniq(bytes.NewReader(fileContent)))
	}

	testCases := []string{
		/* 01 */ "1 2 3 4 5 6 7 8 9 8 7 6 5 4 3 2 1", // basic
		/* 02 */ "11111111111 11111111111 898899899889 898899899889 33333333", // big numbers
		/* 03 */ "11 22 33 33 33 22 11", // more then two equal numbers
		/* 04 */ "1 11 22 33 33 33 11 33 22 22 33 22 33 1 1 1 1 22 22",
		/* 05 */ "0 0 0 0 0 0 0 0 0 0 0 0 0", // zero
		/* 06 */ "-100 -500 100 100 -500", // negative
		/* 07 */ "-100 -0 +0 -100 +0",
		/* 08 */ "-100 -0 +0 -100 -0",
		/* 08 */ "-100 -0 +0 -100 -0 0 +0",
	}

	testExpected := []interface{}{
		/* 01 */ 9,
		/* 02 */ 33333333,
		/* 03 */ 33,
		/* 04 */ 1,
		/* 05 */ 0,
		/* 06 */ -100,
		/* 07 */ 0,
		/* 08 */ 0,
		/* 09 */ 0,
	}

	for i, testCase := range testCases {
		require.Equal(t, testExpected[i], Uniq(strings.NewReader(testCase)))
	}
}

// go test -bench . -benchmem
func BenchmarkFast(b *testing.B) {
	file, err := os.Open(tenMinllionTestPath)
	if err != nil {
		panic(err)
	}

	fileContent, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	file.Close()

	fileContentReader := bytes.NewReader(fileContent)
	for i := 0; i < b.N; i++ {
		Uniq(fileContentReader)
	}
}
