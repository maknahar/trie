# Trie
Go Implementation of a Thread Safe Trie Data Structure and it's *Applications*

# Supported Operation

## Trie

- New() *Trie
- Insert(keywords ...string)
- PrefixSearch(Key string) (found bool)
- Search(keyword string) (found bool)
- Delete(keyword string)
- DeleteBranch(key string)

## Trie Node

- CreateNode(v rune) *Node
- AddChildNode(v rune) *Node
- Len() int
- IsLeafNode() bool
- GetChildNode(v rune) (node *Node, exist bool)
- DeleteChildNode(v rune)


Your ideas and suggestions are always welcome. Either in form of issues or PR.