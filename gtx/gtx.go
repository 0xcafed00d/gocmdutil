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
	path     string
	info     os.FileInfo
	children []*treeNode
	parent   *treeNode
	expanded bool
	last     bool
}

func createNodes(rootPath string, parent *treeNode) ([]*treeNode, error) {

	var res []*treeNode
	err := filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {

		if err == nil && path != rootPath {
			if len(res) > 0 {
				res[len(res)-1].last = false
			}
			res = append(res, &treeNode{path, info, nil, parent, false, true})
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
		children, err := createNodes(node.path, node)
		neo.PanicOnError(err)
		node.children = children
		node.expanded = true
	}
	return nil
}

func drawNode(node *treeNode) {

	preamble := ""
	for n := node; n.parent != nil; n = n.parent {
		if n.parent.last {
			preamble = "   " + preamble
		} else {
			preamble = "│  " + preamble
		}
	}
	fmt.Print(preamble)

	if node.last {
		fmt.Print("└─")
	} else {
		fmt.Print("├─")
	}

	if node.info.IsDir() {
		if node.expanded {
			fmt.Print("[-]")
		} else {
			fmt.Print("[+]")
		}
	} else {
		fmt.Print("── ")
	}

	fmt.Println(node.info.Name())
}

func drawNodes(nodes []*treeNode) {

	for _, node := range nodes {
		drawNode(node)
		if node.expanded && len(node.children) > 0 {
			drawNodes(node.children)
		}
	}
}

func filltree(nodes []*treeNode) {

	for _, node := range nodes {
		populateChildren(node)
		if len(node.children) > 0 {
			filltree(node.children)
		}
	}
}

func test() {
	rootpath := "."
	if len(os.Args) > 1 {
		rootpath = os.Args[1]
	}

	nodes, err := createNodes(rootpath, nil)
	neo.PanicOnError(err)

	filltree(nodes)

	nodes[0].expanded = false

	drawNodes(nodes)

	spew.Dump(err)
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
