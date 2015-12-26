package main

import (
	"fmt"
	"os"
	"path/filepath"

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
		for _, path := range os.Args[1:] {
			f, err := os.Open(path)
			exitOnError(err)
			defer f.Close()

			ex, err := exif.Decode(f)
			exitOnError(err)

			date, err := ex.DateTime()
			if err != nil {
				fmt.Fprintln(os.Stderr, "Error processing:", path, err)
				continue
			}

			dir, file := filepath.Split(path)
			ext := filepath.Ext(file)

			newname := fmt.Sprintf("%04d_%02d_%02d_%02d_%02d_%02d%s",
				date.Year(), int(date.Month()), date.Day(),
				date.Hour(), date.Minute(), date.Second(), ext)

			newpath := filepath.Join(dir, newname)

			fmt.Printf("%s -> %s\n", path, newpath)
			os.Rename(path, newpath)
			if err != nil {
				fmt.Fprintln(os.Stderr, "Error processing:", path, err)
				continue
			}
		}
	}
}
