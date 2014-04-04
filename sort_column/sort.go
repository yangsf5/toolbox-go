// Author: sheppard(ysf1026@gmail.com) 2014-04-04

package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

var (
	column = flag.Int("column", 1, "key column, number of column based on 1")
	sep = flag.String("sep", " ", "separate string")
	sortType = flag.String("type", "int", "type of sort, int or string")
)

func main() {
	flag.Parse()

	if *sortType == "int" {
		data := make(map[int]string)

		inputCb := func(key, line string) {
			intKey, err := strconv.Atoi(key)
			if err != nil {
				panic(err)
			}
			data[intKey] = line
		}
		input(inputCb)

		keys := []int{}
		for k, _ := range data {
			keys = append(keys, k)
		}
		sort.Ints(keys)
		for _, k := range keys {
			fmt.Println(data[k])
		}
	} else if *sortType == "string" {
		data := make(map[string]string)

		inputCb := func(key, line string) {
			data[key] = line
		}
		input(inputCb)

		keys := []string{}
		for k, _ := range data {
			keys = append(keys, k)
		}
		sort.Strings(keys)
		for _, k := range keys {
			fmt.Println(data[k])
		}
	} else {
		panic("sort type is error")
	}
}

func input(cb func(key, line string)) {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		if strs := strings.Split(line, *sep); *column <= len(strs) {
			key := strs[*column - 1]
			cb(key, line)
		}
	}
	if err:= scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "reading standard input:", err)
	}
}

