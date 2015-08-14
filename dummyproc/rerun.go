package main

import (
	"fmt"
	"os"
	"strconv"
	"time"
)

// dummy process used for testing programs that spawn processes

func main() {

	for i, a := range os.Args {
		fmt.Printf("arg: %d: %s\n", i, a)
	}

	if len(os.Args) >= 2 {
		sleeptime, err := strconv.Atoi(os.Args[1])
		if err == nil {
			fmt.Fprintf(os.Stderr, "Sleeping for %d Seconds\n", sleeptime)
			time.Sleep(time.Duration(sleeptime) * time.Second)
		} else {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(-1)
		}
	}

	if len(os.Args) >= 3 {
		exitcode, err := strconv.Atoi(os.Args[2])
		if err == nil {
			fmt.Fprintf(os.Stderr, "Exiting with code: %d\n", exitcode)
			os.Exit(exitcode)
		} else {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(-1)
		}
	}

}
