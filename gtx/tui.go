package main

import (
	"fmt"
	"github.com/nsf/termbox-go"
	"unicode/utf8"
	"os"
	"github.com/simulatedsimian/neo"

)

func main() {
	
	rootpath := "."
	if len(os.Args) > 1 {
		rootpath = os.Args[1]
	}

	rootInfo, err := os.Stat(rootpath)
	neo.PanicOnError(err)

	rootNode := &treeNode{rootpath, rootInfo, nil, nil, true, 0}

	nodes, err := createNodes(rootpath, rootNode)
	neo.PanicOnError(err)
	filltree(nodes)
	rootNode.children = nodes

	var root []*treeNode
	root = append(root, rootNode)

	termtest(rootNode)
}

func drawFromNode (node *treeNode, count int) {
	for i := 0; i < count; i++ {
		p, n := nodeToStrings(node)
		printAtDef (0, i, p + n + "             ")
		var ok bool
		node, ok = nextNode(node)
		if !ok {
			break
		}
	}
}

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

func termtest(root *treeNode) {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	termbox.HideCursor()

	currentNode := root

	for {
	
		drawFromNode (currentNode, 10)	
		termbox.Flush()		
	
		switch ev := termbox.PollEvent(); ev.Type {
		case termbox.EventKey:
			printAt(0, 20, fmt.Sprint(ev), termbox.ColorDefault, termbox.ColorDefault)
			termbox.Flush()

			switch ev.Key {
			case termbox.KeyEsc:
				return
			case termbox.KeyArrowUp:
			case termbox.KeyArrowDown:
				n, ok := nextNode (currentNode)
				if ok {
					currentNode = n
				}			
			}

			termbox.Flush()

		case termbox.EventResize:
			x, y := ev.Width, ev.Height
			printAt(0, 21, fmt.Sprintf("[%d, %d] ", x, y), termbox.ColorDefault, termbox.ColorDefault)
			termbox.Flush()
		}
	}

	termbox.Flush()
}
