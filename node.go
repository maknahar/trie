package trie

type Node struct {
	// list of all the children
	Children      []*Node
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
	}
}

func (n *Node) AddChildNode(v rune) *Node {
	n.childIndexMap[v] = len(n.Children)
	n.Children = append(n.Children, CreateNode(v))
	return n.Children[n.childIndexMap[v]]
}

func (n *Node) Len() int {
	return len(n.childIndexMap)
}

func (n *Node) IsLeafNode() bool {
	return len(n.childIndexMap) == 0
}
