# trie
Go Implementation of Trie Data Dtructure and it's *Applications*

# Supported Operation

## Trie

- New() *Trai
- Insert(keywords ...string)
- Search(keyword string) (found bool)
- SearchWithWordBoundary(keyword string) (found bool)
- Delete(keyword string)


## Trie Node

- CreateNode(val rune) *Node
- AddChildNode(v rune) *Node
- Len() int
- IsLeafNode() bool