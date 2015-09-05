package main

import (
	"fmt"
	"os"
	"os/signal"
	"strconv"
	"time"
)

// dummy process used for testing programs that spawn processes

func main() {

	for i, a := range os.Args {
		fmt.Printf("arg: %d: %s\n", i, a)
	}

	var err error
	var exitCode int
	var sleepyTime time.Duration

	if len(os.Args) >= 2 {
		sleeptime, err := strconv.Atoi(os.Args[1])
		if err == nil {
			fmt.Fprintf(os.Stderr, "Sleeping for %d Seconds\n", sleeptime)
			sleepyTime = time.Duration(sleeptime) * time.Second
		} else {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(-1)
		}
	}

	if len(os.Args) >= 3 {
		exitCode, err = strconv.Atoi(os.Args[2])
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			os.Exit(-1)
		}
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, os.Kill)
	select {
	case s := <-c:
		fmt.Fprintln(os.Stderr, "Signal Received: ", s)

	case <-time.After(sleepyTime):
		fmt.Println("timed out")
	}

	fmt.Printf("Exiting with code: %d\n", exitCode)
	os.Exit(exitCode)
}
