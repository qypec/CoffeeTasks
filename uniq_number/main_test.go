package main

import (
	"testing"
	"github.com/stretchr/testify/require"
	"io/ioutil"
	"os"
	"bytes"
)

func TestUniq(t *testing.T) {

// Test 01 | 1.000.000 pairs
	file, err := os.Open("./testdata/1_000_000_test.txt")
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