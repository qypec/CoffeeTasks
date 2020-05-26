package main

import (
	// "bufio"
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	// "io/ioutil"
	"os"
	"strings"
	// "log"
)


type user struct {
	Browsers []interface{}
	Email string
	Name string
}

func FastSearch(out io.Writer) {
	seenBrowsers := make(map[string]bool)
	// foundUsers := ""

	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	fmt.Fprintln(out, "found users:")
	for i := 0; scanner.Scan(); i++ {
		line := scanner.Bytes()
		isAndroid := false
		isMSIE := false

		user := user{}
		err := json.Unmarshal(line, &user)
		if err != nil {
			panic(err)
		}

		browsers := user.Browsers
		for _, browserRaw := range browsers {
			browser, ok := browserRaw.(string)
			if !ok {
				// log.Println("cant cast browser to string")
				continue
			}
			if ok := strings.Contains(browser, "Android"); ok {
				isAndroid = true
				seenBrowsers[browser] = true;
			}
			if ok := strings.Contains(browser, "MSIE"); ok {
				isMSIE = true
				seenBrowsers[browser] = true;
			}
		}

		if !(isAndroid && isMSIE) {
			continue
		}

		// log.Println("Android and MSIE user:", user["name"], user["email"])
		email := strings.ReplaceAll(user.Email, "@", " [at] ");
		fmt.Fprintf(out, "[%d] %s <%s>\n", i, user.Name, email)
	}
	fmt.Fprintln(out, "\nTotal unique browsers", len(seenBrowsers))
}