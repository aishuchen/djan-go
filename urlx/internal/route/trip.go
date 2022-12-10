package route

import (
	"strings"
)

type Node struct {
	root     *Node
	part     string
	children map[string]*Node
	isWild   bool
}

func NewRoot() *Node {
	root := &Node{children: make(map[string]*Node)}
	root.root = root
	return root
}

func (n *Node) Insert(path string) *Node {
	parts := parsePath(path)
	nn := n
	for _, part := range parts {
		if child, ok := nn.children[part]; ok {
			nn = child
			continue
		}
		child := &Node{root: n, part: part, children: make(map[string]*Node)}
		nn.children[part] = child
		nn = child
	}
	return nn
}

func (n *Node) Find(path string) *Node {
	parts := parsePath(path)
	nn := n
	for i, part := range parts {
		if len(n.children) == 0 && i != len(parts)-1 {
			return nil
		}
		nn = n.children[part]
	}
	return nn
}

func parsePath(path string) []string {
	if path[0] != '/' {
		panic(`path must start with "/"`)
	}
	parts := strings.Split(path[1:], "/")
	return parts
}
