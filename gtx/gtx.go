package main

import (
	"fmt"
	"github.com/davecgh/go-spew/spew"
	"github.com/nsf/termbox-go"
	"github.com/simulatedsimian/neo"
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
	children []*treeNode
	parent   *treeNode
}

func createNodes(path string, parent *treeNode) ([]*treeNode, error) {
	var res []*treeNode

	fmt.Println(path)

	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err == nil && path != "." {
			res = append(res, &treeNode{info, nil, parent})
			if info.IsDir() {
				return filepath.SkipDir
			}
		}
		return nil
	})

	return res, err
}

func populateChildren(node *treeNode) error {

	if node.info.IsDir() {
		children, err := createNodes(getPathFromNode(node), node)
		neo.PanicOnError(err)
		node.children = children
	}
	return nil
}

func getPathFromNode(node *treeNode) string {
	path := node.info.Name()
	for node.parent != nil {
		path = filepath.Join(node.info.Name(), path)
	}
	return path
}

func test() {
	nodes, err := createNodes(".", nil)
	neo.PanicOnError(err)

	for _, node := range nodes {
		populateChildren(node)
	}

	spew.Dump(nodes)
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
