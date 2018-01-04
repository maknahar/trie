package trie

type Node struct {
	// list of all the children
	Children      []*Node
	childIndexMap map[rune]int
	Val           rune
}
