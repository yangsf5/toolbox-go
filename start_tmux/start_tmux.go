// Author: sheppard(ysf1026@gmail.com) 2014-04-28

package main

import (
	"flag"
	"fmt"
	"os/exec"
)

var (
	name = flag.String("name", "", "tmux session name")
	path = flag.String("path", "", "init path of windows")
)

func main() {
	flag.Parse()

	run(*name)
}

func run(name string) {
	steps := [][]string{
		{"tmux", "new-session", "-d", "-s", name, "-n", "git"},
		{"tmux", "new-window", "-t", name+":2", "-n", "src"},
		{"tmux", "new-window", "-t", name+":3", "-n", "bin"},

		{"tmux", "send-keys", "-t", name+":1", "cd "+*path, "C-m"},
		{"tmux", "send-keys", "-t", name+":2", "cd "+*path, "C-m"},
		{"tmux", "send-keys", "-t", name+":3", "cd "+*path, "C-m"},

		//{"tmux", "attach-session", "-t", name},
	}

	for _, step := range steps {
		cmd := exec.Command(step[0], step[1:]...)
		output, err := cmd.CombinedOutput()
		if err != nil {
			panic(fmt.Sprintf("Command error, err=%s\noutput=%s\n", err.Error(), output))
		}
		if len(output) != 0 {
			fmt.Println(output)
		}
	}
}
