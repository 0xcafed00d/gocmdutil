package main

import (
	"github.com/simulatedsimian/neo"
	"os"
	"path/filepath"
)

type nodeData struct {
	path string
	info os.FileInfo
}

func createNodes(rootPath string, parent *treeNode) ([]*treeNode, error) {

	var res []*treeNode
	err := filepath.Walk(rootPath, func(path string, info os.FileInfo, err error) error {

		if err == nil && path != rootPath {
			res = append(res, newTreeNode(parent, len(res), nodeData{path, info}))
			if info.IsDir() {
				return filepath.SkipDir
			}
		}
		return nil
	})

	return res, err
}

func populateChildren(node *treeNode) error {

	if node.data.(nodeData).info.IsDir() {
		children, err := createNodes(node.data.(nodeData).path, node)
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

	if node.data.(nodeData).info.IsDir() {
		if node.expanded {
			preamble += "[-]"
		} else {
			preamble += "[+]"
		}
	} else {
		preamble += "── "
	}

	return preamble, node.data.(nodeData).info.Name()
}

func filltree(nodes []*treeNode) {

	for _, node := range nodes {
		populateChildren(node)
		if len(node.children) > 0 {
			filltree(node.children)
		}
	}
}
