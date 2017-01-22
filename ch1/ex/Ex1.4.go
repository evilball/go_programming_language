//Exercise 1.4: Modify dup2 to print the names of all files in which each duplicated
//line occurs.
package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	counts := make(map[string]int)
	fileCounts := make(map[string]map[string]int)
	files := os.Args[1:]
	if len(files) == 0 {
		countLines(os.Stdin, counts, fileCounts)
	} else {
		for _, arg := range files {
			f, err := os.Open(arg)
			if err != nil {
				fmt.Fprintf(os.Stderr, "dup2: %v\n", err)
				continue
			}
			countLines(f, counts, fileCounts)
			f.Close()
		}
	}
	for line, n := range counts {
		if n > 1 && line != "" {
			fmt.Printf("%d:\t%s\n", n, line)
			for filename, occurences := range fileCounts[line] {
				fmt.Printf("%s: %d occurences\n", filename, occurences)
			}
			fmt.Println("--------------------------------")
		}
	}
}

func countLines(f *os.File, counts map[string]int, fileCount map[string]map[string]int) {
	input := bufio.NewScanner(f)
	for input.Scan() {
		text := input.Text()
		fileName := f.Name()
		counts[text]++
		if fileCount[text] == nil {
			fileCount[text] = make(map[string]int)
		}
		fileCount[text][fileName]++
	}
}
