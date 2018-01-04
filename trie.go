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

func (r *Root) Search(keyword string) (found bool) {
	node := &r.Node
	for _, v := range []rune(keyword) {
		if j, ok := node.childIndexMap[v]; ok {
			node = node.Children[j]
			found = ok
			continue
		}
		return false
	}

	return found
}

func (r *Root) SearchWithWordBoundary(keyword string) (found bool) {
	node := &r.Node
	for _, v := range []rune(keyword) {
		if j, ok := node.childIndexMap[v]; ok {
			node = node.Children[j]
			found = ok
			continue
		}
		return false
	}

	return found && node.IsEndOfWord
}

func (r *Root) Delete(keyword string) {
	node := &r.Node
	var breakNode *Node
	var breakRune rune
	for _, v := range []rune(keyword) {
		if j, ok := node.childIndexMap[v]; ok {
			if node.IsEndOfWord {
				breakNode = node
				breakRune = v
			}
			node = node.Children[j]
			continue
		}
		return
	}

	if !node.IsEndOfWord {
		return
	}

	if !node.IsLeafNode() {
		node.IsEndOfWord = false
		return
	}

	if breakNode == nil {
		breakNode = &r.Node
		breakRune = []rune(keyword)[0]
	}

	breakNode.Children[breakNode.childIndexMap[breakRune]] = nil
	delete(breakNode.childIndexMap, breakRune)
}
