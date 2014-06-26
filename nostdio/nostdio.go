package main

import (
	"fmt"
	"neomech/lib/neo"
	"os"
	"os/exec"
)

func checkArgCount(n int) {
	if len(os.Args) < n {
		fmt.Fprintln(os.Stderr, "Invalid Args: usage: nostdio <command> [comamnd args]")
		os.Exit(1)
	}
}

func main() {
	checkArgCount(2)

	cmd := exec.Command(os.Args[1])
	cmd.Stdin = neo.NullReaderWriterCloser{}
	cmd.Stdout = neo.NullReaderWriterCloser{}
	cmd.Stderr = neo.NullReaderWriterCloser{}

	if len(os.Args) > 2 {
		cmd.Args = append(cmd.Args, os.Args[2:]...)
	}

	cmd.Run()
}
