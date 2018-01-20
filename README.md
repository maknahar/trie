# Trie
Go Implementation of a Thread Safe Trie Data Structure and (some of) Trie *Operations*

Code documentation can be found [here][1].

[1]: https://godoc.org/github.com/maknahar/trie

# Supported Operation

## Operations on Trie Data Structure Level

- **New() \*Trie**

    Creates an initialized trie data structure

- **Insert(keywords ...string)**

    Insert allow one or more keyword to be inserted in trie.
    keyword can be any valid Unicode string

- **PrefixSearch(Key string) (found bool)**

    Search checks if keyword exist in trie as a keyword or prefix to
     keyword

- **Search(keyword string) (found bool)**

    PrefixSearch checks if keyword exist in trie as a fully qualified
     keyword.

- **Delete(keyword string)**

    Delete deletes a keyword from a trie if keyword exist in trie

- DeleteBranch(key string)

    DeleteBranch deletes all child after last letter of key if key
     exists in trie. If key is found, key will be treated as a keyword after this operation

## Operations on Node Level


- **CreateNode(v rune) \*Node**

    CreateNode returns an initialized trie node

- **AddChildNode(v rune) \*Node**

    AddChildNode add child node to current node with value v

- **Len() int**

    Len returns the number of children for node

- **IsLeafNode() bool**

    IsLeafNode returns true if current node is a leaf node in Trie

- **GetChildNode(v rune) (node \*Node, exist bool)**

    GetChildNode retrieve child node with value v.

- **DeleteChildNode(v rune)**

    DeleteChildNode Deletes the child node if it exist

Your ideas and suggestions are always welcome. Either in form of issues or Pull Requests.