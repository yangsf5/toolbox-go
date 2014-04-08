// Author: sheppard(ysf1026@gmail.com) 2014-04-08

package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	segment_column := flag.Int("segment_column", 1, "segment column number")
	calc_column := flag.Int("calc_column", 2, "calculate column number")
	segment := flag.Int("segment", 0, "segment size")
	sep := flag.String("sep", " ", "separate string")
	flag.Parse()

	if *segment == 0 {
		panic("please set segment param")
	}

	segs := make(map[int]int)

	scan := bufio.NewScanner(os.Stdin)
	for scan.Scan() {
		line := scan.Text()
		strs := strings.Split(line, *sep)
		if *segment_column > len(strs) || *calc_column > len(strs) {
			panic("count of column err, at line:" + line)
		}
		intSegColumn, err := strconv.Atoi(strs[*segment_column-1])
		if err != nil {
			panic(err)
		}
		intCalcColumn, err := strconv.Atoi(strs[*calc_column-1])
		if err != nil {
			panic(err)
		}
		index :=  intSegColumn / *segment
		segs[index] += intCalcColumn
	}

	if err := scan.Err(); err != nil {
		panic(err)
	}

	for i, data := range segs {
		fmt.Printf("%d~%d %d\n", *segment*i, *segment*(i+1), data)
	}
}
