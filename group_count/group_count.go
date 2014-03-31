// Author: sheppard(ysf1026@gmail.com) 2014-03-31

package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"regexp"
)

func main() {
	fileName := flag.String("file", "", "file name")
	reg := flag.String("reg", "", "regexp")
	flag.Parse()

	if *fileName == "" || *reg == "" {
		panic("file or reg is empty")
	}

	re := regexp.MustCompile(*reg)

	file, err := os.Open(*fileName)
	if err != nil {
		panic(err)
	}

	groups := make(map[string]int)
	scan := bufio.NewScanner(file)
	for scan.Scan() {
		group := re.FindString(scan.Text())
		groups[group]++
	}

	if err := scan.Err(); err != nil {
		panic(err)
	}

	for name, count := range groups {
		fmt.Printf("count of [%s]: %d\n", name, count)
	}
}
