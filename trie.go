package trie

type Node struct {
	Childrens []*Node

	// isEndOfWord is true if the node
	// represents end of a word
	IsEndOfWord bool
}
