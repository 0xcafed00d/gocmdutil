package main

import (
	"fmt"
	"github.com/nsf/termbox-go"
	"github.com/simulatedsimian/go_sandbox/geom"
	"github.com/simulatedsimian/neo"
	"os"
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

func drawFromNode(node *treeNode, count int) {
	for i := 0; i < count; i++ {
		p, n := nodeToStrings(node)
		printAtDef(0, i, p+n+"             ")
		var ok bool
		node, ok = nextNode(node)
		if !ok {
			break
		}
	}
}

func termtest(root *treeNode) {
	err := termbox.Init()
	if err != nil {
		panic(err)
	}
	defer termbox.Close()

	termbox.HideCursor()

	termx, termy := termbox.Size()

	currentNode := root

	for {
		clearRectDef(geom.RectangleFromSize(geom.Coord{termx, termy}))
		drawFromNode(currentNode, termy)
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
				n, ok := nextNode(currentNode)
				if ok {
					currentNode = n
				}
			}

			termbox.Flush()

		case termbox.EventResize:
			termx, termy = ev.Width, ev.Height
			printAt(0, 21, fmt.Sprintf("[%d, %d] ", termx, termy), termbox.ColorDefault, termbox.ColorDefault)
			termbox.Flush()
		}
	}

	termbox.Flush()
}
