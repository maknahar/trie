package trie

import (
	"testing"
)

func TestRoot_Insert(t *testing.T) {
	trie := New().Insert("mayank")

	ShouldBeEqual(t, 1, len(trie.Children))
	ShouldBeEqual(t, 1, len(trie.Insert("mayank").Children))
	ShouldBeEqual(t, 2, len(trie.Insert("mayank").Children))
	ShouldBeEqual(t, 3, len(trie.Insert("nahar").Children))
}

func TestRoot_Search(t *testing.T) {
	trie := New().Insert("mayank")

	ShouldBeEqual(t, 1, len(trie.Children))
	ShouldBeEqual(t, 1, len(trie.Insert("mayank").Children))
	ShouldBeEqual(t, 2, len(trie.Insert("mayank").Children))
	ShouldBeEqual(t, 3, len(trie.Insert("nahar").Children))

	ShouldBeEqual(t, trie.Search("maya"), true)
	ShouldBeEqual(t, trie.Search("foo"), false)
	ShouldBeEqual(t, trie.Search("mayank"), true)
	ShouldBeEqual(t, trie.Search("patel"), true)
}
