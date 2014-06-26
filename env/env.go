package main

import (
	"fmt"
	"os"
	"strings"
)

// lists envionment vars
func main() {
	var match string

	if len(os.Args) > 1 {
		match = strings.ToLower(os.Args[1])
	}

	env := os.Environ()
	for _, v := range env {
		if len(match) > 0 {
			parts := strings.SplitN(v, "=", 2)
			if strings.Contains(strings.ToLower(parts[0]), match) {
				fmt.Println(v)
			}
		} else {
			fmt.Println(v)
		}
	}
}
