package trie

import "sync"

type Node struct {
	// list of all the children
	Children      []*Node
	m             *sync.RWMutex
	childIndexMap map[rune]int
	Val           rune

	// isEndOfWord is true if the node
	// represents end of a word
	IsEndOfWord bool
}

// CreateNode returns an initialized trie node
func CreateNode(v rune) *Node {
	return &Node{
		Children:      make([]*Node, 0),
		childIndexMap: make(map[rune]int),
		Val:           v,
		m:             new(sync.RWMutex),
	}
}

// AddChildNode add child node to current node with value v
func (n *Node) AddChildNode(v rune) *Node {
	node, exist := n.GetChildNode(v)
	if exist {
		return node
	}
	n.m.Lock()
	n.childIndexMap[v] = len(n.Children)
	n.m.Unlock()
	node = CreateNode(v)
	n.Children = append(n.Children, node)
	return node
}

// Len returns the number of children of node
func (n *Node) Len() int {
	n.m.RLock()
	l := len(n.childIndexMap)
	n.m.RUnlock()
	return l
}

// IsLeafNode returns true if current node is a leaf node in Trie
func (n *Node) IsLeafNode() bool {
	return n.Len() == 0
}

// GetChildNode retrieve child node with value v.
func (n *Node) GetChildNode(v rune) (node *Node, exist bool) {
	n.m.RLock()
	defer n.m.RUnlock()
	if i, ok := n.childIndexMap[v]; !ok {
		return nil, false
	} else {
		return n.Children[i], true
	}
}

// DeleteChildNode Deletes the child node if it exist
func (n *Node) DeleteChildNode(v rune) {
	n.m.Lock()
	defer n.m.Unlock()
	n.Children[n.childIndexMap[v]] = nil
	delete(n.childIndexMap, v)
}
