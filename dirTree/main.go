package main

import (
	"fmt"
	"os"
	"sort"
)

func isFileCloser(index, dirSize int) bool {
	if index == dirSize - 1 {
		return true
	}
	return false
}

func printBranches(out *os.File, fileCloser bool, nestingLvl int) {
	for i := 0; i < nestingLvl; i++ {
		fmt.Print("\t")
		if i + 1 < nestingLvl {
			fmt.Print("|")
		}
	}
	if fileCloser {
		fmt.Print("└")
	} else {
		fmt.Print("|")
	}
	fmt.Print("───")
}

func printFile(out *os.File, file os.FileInfo, fileCloser bool, nestingLvl int) {
	printBranches(out, fileCloser, nestingLvl)
	fmt.Printf("%v (", file.Name())
	if file.Size() == 0 {
		fmt.Printf("empty)\n")
	} else {
		fmt.Printf("%vb)\n", file.Size())
	}
}

func printDir(out *os.File, file os.FileInfo, fileCloser bool, nestingLvl int) {
	printBranches(out, fileCloser, nestingLvl)
	fmt.Printf("%v\n", file.Name())
}

func readPath(path string) ([]os.FileInfo, error) {
	dirFile, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	files, err := dirFile.Readdir(0)
	if err != nil {
		return nil, err
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

func printDirectoryFiles(out *os.File, path string, printFiles bool, nestingLvl int) error {
	files, err := readPath(path)
	if err != nil {
		return err
	}

	sort.Slice(files, func(i, j int) bool {
		return files[i].Name() < files[j].Name()
	})

	for i, file := range files {
		fileCloser := isFileCloser(i, len(files))
		if !file.IsDir() && printFiles {
			printFile(out, file, fileCloser, nestingLvl)
		}

		if file.IsDir() {
			printDir(out, file, fileCloser, nestingLvl)
			err := printDirectoryFiles(out, ConcatenatePaths(path, file.Name()), printFiles, nestingLvl + 1)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func dirTree(out *os.File, path string, printFiles bool) error {
	return printDirectoryFiles(out, path, printFiles, 0)
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
