# trie
Go Implementation of a Thread Safe Trie Data Dtructure and it's *Applications*

# Supported Operation

## Trie

- New() *Trie
- Insert(keywords ...string)
- PrefixSearch(Key string) (found bool)
- Search(keyword string) (found bool)
- Delete(keyword string)

## Trie Node

- CreateNode(v rune) *Node
- AddChildNode(v rune) *Node
- Len() int
- IsLeafNode() bool
- GetChildNode(v rune) (node *Node, exist bool)
- DeleteChildNode(v rune)