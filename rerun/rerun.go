package main

import (
	"fmt"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"
)

func checkArgCount(n int) {
	if len(os.Args) < n {
		fmt.Fprintln(os.Stderr, "Invalid Args: usage: rerun [-restart delay] <command> [comamnd args]")
		os.Exit(1)
	}
}

func main() {
	checkArgCount(2)

	sleepytime := 1
	commandindex := 1

	if strings.HasPrefix(os.Args[1], "-") {
		i, err := strconv.Atoi(os.Args[1][1:])
		if err != nil {
			fmt.Fprintln(os.Stderr, "Invalid restart delay: ", os.Args[1], err)
			os.Exit(1)
		}
		sleepytime = i
		commandindex++
		checkArgCount(3)
	}

	for {
		cmd := exec.Command(os.Args[commandindex])
		cmd.Stdin = os.Stdin
		cmd.Stdout = os.Stdout
		cmd.Stderr = os.Stderr

		if len(os.Args) > 2 {
			cmd.Args = append(cmd.Args, os.Args[commandindex+1:]...)
		}

		err := cmd.Run()
		if err != nil {
			fmt.Fprintf(os.Stderr, "%s\n", err)
		}

		time.Sleep(time.Duration(sleepytime) * time.Second)
	}
}
