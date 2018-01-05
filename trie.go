package trie

// Trie Represent a node of trie
// Not advised to create the node directly. Use helper function CreateNode() instead
type Trie struct {
	Node
}

// Create returns an initialized trie root node
func New() *Trie {
	return &Trie{
		*CreateNode(0),
	}
}

// Insert allow one or more keyword to be inserted in trie
// keyword can be any string
func (r *Trie) Insert(keywords ...string) *Trie {
	for _, v := range keywords {
		r.insert(v)
	}

	return r
}

func (r *Trie) insert(keyword string) {
	node := &r.Node
	for _, v := range []rune(keyword) {
		node = node.AddChildNode(v)
	}
	node.IsEndOfWord = true
}

func (r *Trie) Search(keyword string) (found bool) {
	node := &r.Node
	for _, v := range []rune(keyword) {
		if n, ok := node.GetChildNode(v); ok {
			node = n
			found = ok
			continue
		}
		return false
	}

	return found
}

func (r *Trie) SearchWithWordBoundary(keyword string) (found bool) {
	node := &r.Node
	for _, v := range []rune(keyword) {
		if n, ok := node.GetChildNode(v); ok {
			node = n
			found = ok
			continue
		}
		return false
	}

	return found && node.IsEndOfWord
}

func (r *Trie) Delete(keyword string) {
	node := &r.Node
	var breakNode *Node
	var breakRune rune
	for _, v := range []rune(keyword) {
		if n, ok := node.GetChildNode(v); ok {
			if node.IsEndOfWord {
				breakNode = node
				breakRune = v
			}
			node = n
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

	breakNode.DeleteChildNode(breakRune)
}
