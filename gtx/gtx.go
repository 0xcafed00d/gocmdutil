package main

import (
	"fmt"
	"os"
	"path/filepath"
	"github.com/davecgh/go-spew/spew"
	"github.com/simulatedsimian/neo"
)

type treeNode struct {
	path     string
	info     os.FileInfo
	children []*treeNode
	parent   *treeNode
	expanded bool
	index    int
}

func (n *treeNode) isLast() bool {
	if n.parent == nil {
		return true
	} else {
		return n.index == len(n.parent.children)-1
	}
}

func (n *treeNode) nextSibling() (*treeNode, bool) {
	if n.isLast() {
		return nil, false
	} else {
		return n.parent.children[n.index+1], true
	}
}

func (n *treeNode) prevSibling() (*treeNode, bool) {
	if n.index == 0 {
		return nil, false
	} else {
		return n.parent.children[n.index-1], true
	}
}

func createNodes(rootPath string, parent *treeNode) ([]*treeNode, error) {

	var res []*treeNode
	err := filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {

		if err == nil && path != rootPath {
			res = append(res, &treeNode{path, info, nil, parent, false, len(res)})
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

func nodeToStrings(node *treeNode) (string, string) {

	preamble := ""
	for n := node; n.parent != nil; n = n.parent {
		if n.parent.isLast() {
			preamble = "   " + preamble
		} else {
			preamble = "│  " + preamble
		}
	}

	if node.isLast() {
		preamble += "└─"
	} else {
		preamble += "├─"
	}

	if node.info.IsDir() {
		if node.expanded {
			preamble += "[-]"
		} else {
			preamble += "[+]"
		}
	} else {
			preamble += "── "
	}

	return preamble, node.info.Name()
}

func drawNodes(nodes []*treeNode) {

	for _, node := range nodes {
		fmt.Println(nodeToStrings(node))
		if node.expanded && len(node.children) > 0 {
			drawNodes(node.children)
		}
	}
}

func nextNode(node *treeNode) (*treeNode, bool) {

	if node.children != nil && len(node.children) > 0 {
		return node.children[0], true
	}

	if node.isLast() {
		for node.parent != nil {
			node = node.parent
			if ps, ok := node.nextSibling(); ok {
				return ps, ok
			}
		}
		return nil, false
	}
	return node.nextSibling()
}

func advanceNodes(node *treeNode, advanceCount int) (*treeNode, bool) {
	for n := 0; n < advanceCount; n++ {
		next, ok := nextNode(node)
		if !ok {
			return node, false
		}
		node = next
	}
	return node, true
}

func prevNode(node *treeNode) *treeNode {
	return nil
}

func drawNodesFrom(node *treeNode, count int) {
	siblings := node.parent.children
	for i, max := node.index, len(siblings); i < max; i++ {
		fmt.Println(nodeToStrings(siblings[i]))
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

	rootInfo, err := os.Stat(rootpath)
	neo.PanicOnError(err)

	rootNode := &treeNode{rootpath, rootInfo, nil, nil, true, 0}

	nodes, err := createNodes(rootpath, rootNode)
	neo.PanicOnError(err)
	filltree(nodes)
	rootNode.children = nodes

	var root []*treeNode
	root = append(root, rootNode)

	//	nodes[0].expanded = false

	//drawNodes(root)

	start, _ := advanceNodes(rootNode, 20)

	for n, ok := start, true; ok; n, ok = nextNode(n) {
		fmt.Println(nodeToStrings(n))
	}

	spew.Dump(err)
	fmt.Println(err)
}

func mainx() {
	test()
}
