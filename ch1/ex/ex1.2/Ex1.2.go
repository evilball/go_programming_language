//Exercise 1.2: Modify the echo program to print the index and value of each of its
//arguments, one per line.

package main

import (
	"os"
	"fmt"
	"strconv"
)

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Println(strconv.Itoa(i) + ": " + os.Args[i])
	}
}
