package main

import (
	"fmt"
	"os"

	"github.com/rwcarlsen/goexif/exif"
)

func exitOnError(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func main() {
	if len(os.Args) > 1 {
		for _, v := range os.Args[1:] {
			f, err := os.Open(v)
			exitOnError(err)
			defer f.Close()

			ex, err := exif.Decode(f)
			exitOnError(err)

			date, err := ex.DateTime()

			newname := fmt.Sprintf("IMG_%04d_%02d_%02d_%02d_%02d_%02d.jpg",
				date.Year(), int(date.Month()), date.Day(),
				date.Hour(), date.Minute(), date.Second())

			fmt.Printf("%s -> %s\n", v, newname)
		}
	}
}
