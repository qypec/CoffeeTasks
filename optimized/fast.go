package main

import (
	// "bufio"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"strings"
	// "log"
)


type User struct {
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

	fileContents, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(fileContents), "\n")
	fmt.Fprintln(out, "found users:")
	for i, line := range lines {
		isAndroid := false
		isMSIE := false

		user := User{}
		err := json.Unmarshal([]byte(line), &user)
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