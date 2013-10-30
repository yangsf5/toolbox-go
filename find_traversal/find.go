// Author: sheppard(ysf1026@gmail.com) 2013-10-28

package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
)

func main() {
	if len(os.Args) != 3 {
		fmt.Println("Usage:", os.Args[0], "filetype", "string")
		return
	}

	fileSuffix := "." + os.Args[1]
	targetStr := os.Args[2]

	visit := func(path string, info os.FileInfo, err error) error {
		name := info.Name()
		if !info.IsDir() && strings.HasSuffix(name, fileSuffix) {
			file, err := os.Open(path)
			if err != nil {
				fmt.Println(err)
				return nil
			}
			scanner := bufio.NewScanner(file)
			for i := 1; scanner.Scan(); i++ {
				if strings.Contains(scanner.Text(), targetStr) {
					fmt.Println(path, i, scanner.Text())
				}
			}
			if err = scanner.Err(); err != nil {
				fmt.Println(err)
				return nil
			}
		}
		return nil
	}
	filepath.Walk(".", visit)
}

