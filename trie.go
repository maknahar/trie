// Go Implementation of a Thread Safe Trie Data Structure and (some of) Trie Operations
package trie

// Trie Represent a node of trie
// Not advised to create the node directly. Use helper function CreateNode() instead
type Trie struct {
	Node
}

// New Creates an initialized trie data structure
func New() *Trie {
	return &Trie{
		*CreateNode(0),
	}
}

// Insert allow one or more keyword to be inserted in trie
// keyword can be any unicode string
func (t *Trie) Insert(keywords ...string) *Trie {
	for _, v := range keywords {
		t.insert(v)
	}

	return t
}

func (t *Trie) insert(keyword string) {
	node := &t.Node
	for _, v := range []rune(keyword) {
		node = node.AddChildNode(v)
	}
	node.IsEndOfWord = true
}

// PrefixSearch checks if keyword exist in trie as a keyword or prefix to a keyword
func (t *Trie) PrefixSearch(key string) (found bool) {
	node := &t.Node
	for _, v := range []rune(key) {
		if n, ok := node.GetChildNode(v); ok {
			node = n
			found = ok
			continue
		}
		return false
	}

	return found
}

// Search checks if keyword exist in trie as a fully qualified keyword.
func (t *Trie) Search(keyword string) (found bool) {
	node := &t.Node
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

// Delete deletes a keyword from a trie if keyword exist in trie
func (t *Trie) Delete(keyword string) {
	node := &t.Node
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
		breakNode = &t.Node
		breakRune = []rune(keyword)[0]
	}

	breakNode.DeleteChildNode(breakRune)
}

// DeleteBranch deletes all child after last letter of key if key exists in trie
// If key is found, key will be treated as a keyword after this operation
func (t *Trie) DeleteBranch(key string) {
	node := &t.Node
	for _, v := range []rune(key) {
		if n, ok := node.GetChildNode(v); ok {
			node = n
			continue
		}
		return
	}

	node.childIndexMap = make(map[rune]int)
	node.IsEndOfWord = true
	node.Children = make([]*Node, 0)
}
