// Author: sheppard(ysf1026@gmail.com) 2013-10-09

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage:", os.Args[0], "filetype")
		return
	}

	fileSuffix := "." + os.Args[1]

	var sum int
	visit := func(path string, info os.FileInfo, err error) error {
		name := info.Name()
		if !info.IsDir() && strings.HasSuffix(name, fileSuffix) {
			in, _ := ioutil.ReadFile(path)
			count := strings.Count(string(in), "\n")
			sum += count
		}
		return nil
	}
	filepath.Walk(".", visit)
	fmt.Printf("file_type=%s count_code_line=%d\n", os.Args[1], sum)
}

