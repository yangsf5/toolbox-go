// Author: sheppard(ysf1026@gmail.com) 2013-11-01

package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strings"
)

var (
	column = flag.Int("column", 1, "which column, number of column")
	sep = flag.String("sep", " ", "separate string")
)

func main() {
	flag.Parse()
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		if strs := strings.Split(scanner.Text(), *sep); *column <= len(strs) {
			fmt.Println(strs[*column - 1])
		}
	}
	if err:= scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}
