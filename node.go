package trie

type Node struct {
	// list of all the children
	Children      []*Node
	childIndexMap map[rune]int
	Val           rune
}

func (n *Node) AddChildNode(v rune) *Node {
	n.childIndexMap[v] = len(n.Children)
	n.Children = append(n.Children, CreateNode(v))
	return n.Children[n.childIndexMap[v]]
}
