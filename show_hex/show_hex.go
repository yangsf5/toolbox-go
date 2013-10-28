// Author: sheppard(ysf1026@gmail.com) 2013-10-14

package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Usage:", os.Args[0], "figure")
		return
	}

	figure, _ := strconv.Atoi(os.Args[1])

	fmt.Printf("%d to hex: %x\n", figure, figure)
}

