package testdata

import (
	"os"
	"fmt"
)

const limit = 1000000

func main() {
	file, err := os.Create("testdata.txt")
     
    if err != nil{
        fmt.Println("Unable to create file:", err) 
        os.Exit(1) 
    }
    defer file.Close() 

	for i := 1; i <= limit; i++ {
		fmt.Fprintf(file, "%v %v ", i, i);
		fmt.Println(i)
	}
	fmt.Fprintf(file, "0");
}
