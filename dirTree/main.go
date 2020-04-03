package main

import (
	"fmt"
	"io"
	"os"
	"sort"
)

const emptySymb = rune(' ')
const defaultSymb = rune('|')
const closerSymb = rune('└')

func isFileCloser(index, dirSize int) bool {
	if index == dirSize - 1 {
		return true
	}
	return false
}

func printFile(out *os.File, file os.FileInfo, dirNamePrefix []rune) {
	fmt.Printf("%c", dirNamePrefix[0])
	for i := 1; i < len(dirNamePrefix); i++ {
		fmt.Printf("\t")
		if dirNamePrefix[i] != emptySymb {
			fmt.Printf("%c", dirNamePrefix[i])
		}
	}
	fmt.Printf("───%v", file.Name())
	if !file.IsDir() {
		if file.Size() == 0 {
			fmt.Printf(" (empty)")
		} else {
			fmt.Printf(" (%vb)", file.Size())
		}
	}
	fmt.Printf("\n")
}

func readPath(path string, printFiles bool) ([]os.FileInfo, error) {
	dirFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	files := make([]os.FileInfo, 0)
	for file, err := dirFile.Readdir(1); err != io.EOF; file, err = dirFile.Readdir(1) {
		if err != nil {
			return nil, err
		}
		if file[0].IsDir() || printFiles {
			files = append(files, file[0])
		}
	}
	return files, nil
}

func ConcatenatePaths(first string, second string) string {
	newStr := make([]rune, len(first))
	newStr = []rune(first)
	newStr = append(newStr, '/')
	newStr = append(newStr, []rune(second)...)
	return string(newStr)
}

func printDirectoryFiles(out *os.File, path string, printFiles bool, dirNamePrefix []rune) error {
	files, err := readPath(path, printFiles) // creates an array of files in this directory
	if err != nil {
		return err
	}
	sort.Slice(files, func(i, j int) bool { // sorts alphabetically
		return files[i].Name() < files[j].Name()
	})

	for i, file := range files { // printing
		if isFileCloser(i, len(files)) {
			dirNamePrefix[len(dirNamePrefix) - 1] = closerSymb
		}
		printFile(out, file, dirNamePrefix)
		if isFileCloser(i, len(files)) {
			dirNamePrefix[len(dirNamePrefix) - 1] = emptySymb
		}
		if file.IsDir() {
			tmp := append(dirNamePrefix, defaultSymb)
			err := printDirectoryFiles(out, ConcatenatePaths(path, file.Name()), printFiles, tmp)
			if err != nil { return err }
		}
	}
	return nil
}

func dirTree(out *os.File, path string, printFiles bool) error {
	dirNamePrefix := []rune{defaultSymb}
	return printDirectoryFiles(out, path, printFiles, dirNamePrefix)
}

func main() {
	out := os.Stdout
	if !(len(os.Args) == 2 || len(os.Args) == 3) {
		panic("usage go run main.go . [-f]")
	}
	path := os.Args[1]
	printFiles := len(os.Args) == 3 && os.Args[2] == "-f"
	err := dirTree(out, path, printFiles)
	if err != nil {
		panic(err.Error())
	}
}
