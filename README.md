# trie
Go Implementation of a Thread Safe Trie Data Dtructure and it's *Applications*

# Supported Operation

## Trie

- New() *Trie
- Insert(keywords ...string)
- Search(keyword string) (found bool)
- SearchWithWordBoundary(keyword string) (found bool)
- Delete(keyword string)


## Trie Node

- CreateNode(val rune) *Node
- AddChildNode(v rune) *Node
- Len() int
- IsLeafNode() bool
- GetChildNode(v rune) (node *Node, exist bool)
- DeleteChildNode(v rune)