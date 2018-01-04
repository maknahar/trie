package trie

import (
	"testing"
)

func TestRoot_Insert(t *testing.T) {
	trie := New().Insert("mayank")
	ShouldBeEqual(t, 1, len(trie.Children))
	trie = New().Insert("mayank")
	ShouldBeEqual(t, 1, len(trie.Children))
}
