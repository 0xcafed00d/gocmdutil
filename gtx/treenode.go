package main

type treeNode struct {
	children []*treeNode
	parent   *treeNode
	expanded bool
	index    int
	data     interface{}
}

func (n *treeNode) isLast() bool {
	if n.parent == nil {
		return true
	} else {
		return n.index == len(n.parent.children)-1
	}
}

func (n *treeNode) nextSibling() *treeNode {
	if n.isLast() {
		return nil
	} else {
		return n.parent.children[n.index+1]
	}
}

func (n *treeNode) prevSibling() *treeNode {
	if n.index == 0 {
		return nil
	} else {
		return n.parent.children[n.index-1]
	}
}

func (n *treeNode) deepestNode() *treeNode {
	if n.children != nil {
		lastChild := n.children[len(n.children)-1]
		return lastChild.deepestNode()
	}
	return n
}

func (node *treeNode) nextNode() *treeNode {

	// this node has children? then next node is first child.
	// TODO check to see if node is expanded.
	if node.children != nil && len(node.children) > 0 {
		return node.children[0]
	}

	// if we are last sibling - taverse up tree till we find a
	// parent with a sibling.
	if node.isLast() {
		for node.parent != nil {
			node = node.parent
			if ps := node.nextSibling(); ps != nil {
				return ps
			}
		}
		// no parents with siblings
		return nil
	}

	return node.nextSibling()
}

func (node *treeNode) prevNode() *treeNode {

	if node.index == 0 {
		if node.parent != nil {
			return node.parent
		}

		return nil
	}

	prevSib := node.prevSibling()

	return prevSib.deepestNode()
}

func iabs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func (node *treeNode) traverseNodes(count int) (*treeNode, int) {
	for n := 0; n < iabs(count); n++ {
		var next *treeNode
		if count < 0 {
			next = node.prevNode()
		} else {
			next = node.nextNode()
		}
		if next == nil {
			return node, n
		}
		node = next
	}
	return node, count
}
