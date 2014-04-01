// Author: sheppard(ysf1026@gmail.com) 2014-04-01

package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var (
	readme *os.File
)

func main() {
	paths, err := filepath.Glob("*")
	checkError(err)

	readme, err = os.OpenFile("./README.md", os.O_WRONLY, os.ModePerm)
	checkError(err)
	defer readme.Close()

	// Head of README.md
	head :=
`toolbox-go
===============
my small tools.

`
	write(head)

	for _, path := range paths {
		fi, err := os.Stat(path)
		checkError(err)

		if fi.IsDir() {
			gen(path)
		}
	}
}


func gen(path string) {
	egPath := path + "/example"
	fi, err := os.Stat(egPath)
	if err != nil || !fi.IsDir() {
		return
	}

	err = os.Chdir(egPath)
	defer os.Chdir("../..")
	checkError(err)
	runFile, err := os.Open("./run")
	checkError(err)

	runCmd, err := ioutil.ReadAll(runFile)
	checkError(err)
	cmdStr := strings.Trim(string(runCmd), "\n")

	cmdArgs := strings.Split(cmdStr, " ")
	cmd := exec.Command(cmdArgs[0], cmdArgs[1:]...)

	output, err := cmd.Output()
	checkError(err)

	writeExample(path, "", cmdStr, string(output))
}

func writeExample(name, input, cmd, output string) {
	write("### " + name + "\n")
	write("#### input\n")
	write("```\n")
	write(input)
	write("```\n")

	write("#### run\n")
	write("```\n")
	write(cmd + "\n")
	write("```\n")

	write("#### output\n")
	write("```\n")
	write(output)
	fmt.Println(output)
	write("```\n")
}

func write(str string) {
	_, err := readme.Write([]byte(str))
	checkError(err)
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
