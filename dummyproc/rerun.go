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
		}
	}
}
