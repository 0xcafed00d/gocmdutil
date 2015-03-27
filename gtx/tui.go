package main

import (
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

	rootNode := &treeNode{nil, nil, true, 0, nodeData{rootpath, rootInfo}}

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

		node = node.nextNode()
		if node == nil {
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
			switch ev.Key {
			case termbox.KeyEsc:
				return
			case termbox.KeyArrowUp:
				if node := currentNode.prevNode(); node != nil {
					currentNode = node
				}

			case termbox.KeyArrowDown:
				if node := currentNode.nextNode(); node != nil {
					currentNode = node
				}

			case termbox.KeyPgdn:
				node, _ := currentNode.advanceNodes(10)
				currentNode = node

			case termbox.KeyPgup:
				node, _ := currentNode.retreatNodes(10)
				currentNode = node
			}

			termbox.Flush()

		case termbox.EventResize:
			termx, termy = ev.Width, ev.Height
			termbox.Flush()
		}
	}

	termbox.Flush()
}
