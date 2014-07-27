package main

import (
	"fmt"
	"github.com/nsf/termbox-go"
	//	"github.com/simulatedsimian/neo"
	"github.com/davecgh/go-spew/spew"
	"os"
	"path/filepath"
	"unicode/utf8"
)

func printAt(x, y int, s string, fg, bg termbox.Attribute) {
	for len(s) > 0 {
		r, rlen := utf8.DecodeRuneInString(s)
		termbox.SetCell(x, y, r, fg, bg)
		s = s[rlen:]
		x++
	}
}

func printAtDef(x, y int, s string) {
	printAt(x, y, s, termbox.ColorDefault, termbox.ColorDefault)
}

type treeNode struct {
	info     os.FileInfo
	children []treeNode
}

func createNodes(path string) ([]treeNode, error) {
	var res []treeNode

	walkFunc := func(path string, info os.FileInfo, err error) error {
		if err == nil {
			if info.IsDir() && path != "." {
				return filepath.SkipDir
			} else {
				res = append(res, treeNode{info, nil})			
			}
		}
		return nil
	}

	err := filepath.Walk(path, walkFunc)

	if err == nil {
		return res, nil
	} else {
		return nil, err
	}
}

func test() {
	
	nodes, err := createNodes (".")
	
	spew.Dump(&nodes)
	fmt.Println(err)

}

func termtest() {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	termbox.HideCursor()

	for {
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			printAt(0, 1, fmt.Sprint(ev), termbox.ColorDefault, termbox.ColorDefault)
			termbox.Flush()

			switch ev.Key {
			case termbox.KeyEsc:
				return
			case termbox.KeyArrowUp:
			case termbox.KeyArrowDown:
			}

			termbox.Flush()

		case termbox.EventResize:
			x, y := ev.Width, ev.Height
			printAt(0, 0, fmt.Sprintf("[%d, %d] ", x, y), termbox.ColorDefault, termbox.ColorDefault)
			termbox.Flush()
		}
	}

	termbox.Flush()

}

func main() {

	test()

}
