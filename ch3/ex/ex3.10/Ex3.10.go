//Exercise 3.10: Write a non-recursive version of comma, using bytes.Buffer
//instead of string concatenation.

package main

import (
	"bytes"
	"os"
	"fmt"
)

func comma(s string) string {
	var buf bytes.Buffer
	start := len(s)%3
	if start == 0 {
		start = 3
	}
	buf.WriteString(s[:start])
	for i := start ; i < len(s); i+=3 {
		buf.WriteString("," + s[i:i+3])
	}
	return buf.String()
}

func main() {
	for i := 1; i < len(os.Args); i++ {
		fmt.Printf("%s\n", comma(os.Args[i]))
	}
}