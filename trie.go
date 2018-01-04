package trie

// Root Represent a node of trie
// Not advised to create the node directly. Use helper function CreateNode() instead
type Root struct {
	Node
}

// CreateNode returns an initialized trie node
func CreateNode(val rune) *Node {
	return &Node{
		Children:      make([]*Node, 0),
		childIndexMap: make(map[rune]int),
		Val:           val,
	}
}

// Create returns an initialized trie root node
func New() *Root {
	return &Root{
		*CreateNode(0),
	}
}

// Insert allow one or more keyword to be inserted in trie
// keyword can be any string
func (r *Root) Insert(keywords ...string) *Root {
	for _, v := range keywords {
		r.insert(v)
	}

	return r
}

func (r *Root) insert(keyword string) {
	node := &r.Node
	for _, v := range []rune(keyword) {
		if j, ok := node.childIndexMap[v]; ok {
			node = node.Children[j]
			continue
		}
		node = node.AddChildNode(v)
	}
	node.IsEndOfWord = true
}

func (r *Root) Search(keyword string) (isFound bool) {
	node := &r.Node
	for _, v := range []rune(keyword) {
		if j, ok := node.childIndexMap[v]; ok {
			node = node.Children[j]
			isFound = ok
			continue
		}
		return false
	}

	return isFound
}

func (r *Root) SearchWithWordBoundary(keyword string) (isFound bool) {
	node := &r.Node
	for _, v := range []rune(keyword) {
		if j, ok := node.childIndexMap[v]; ok {
			node = node.Children[j]
			isFound = ok
			continue
		}
		return false
	}

	return isFound && node.IsEndOfWord
}
