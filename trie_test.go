package trie

import (
	"testing"
)

func TestRoot_Insert(t *testing.T) {
	trie := New().Insert("mayank")

	ShouldBeEqual(t, 1, len(trie.Children))
	ShouldBeEqual(t, 1, len(trie.Insert("mayank").Children))
	ShouldBeEqual(t, 2, len(trie.Insert("patel").Children))
	ShouldBeEqual(t, 3, len(trie.Insert("nahar").Children))
	ShouldBeEqual(t, 3, len(trie.Insert("mayank", "nahar").Children))
}

func TestRoot_Search(t *testing.T) {
	trie := New().Insert("mayank")

	ShouldBeEqual(t, 1, len(trie.Children))
	ShouldBeEqual(t, 1, len(trie.Insert("mayank").Children))
	ShouldBeEqual(t, 2, len(trie.Insert("patel").Children))
	ShouldBeEqual(t, 3, len(trie.Insert("nahar").Children))

	ShouldBeEqual(t, true, trie.Search("maya"))
	ShouldBeEqual(t, false, trie.Search("foo"))
	ShouldBeEqual(t, true, trie.Search("mayank"))
	ShouldBeEqual(t, true, trie.Search("patel"))
}

func TestRoot_SearchWithWordBoundary(t *testing.T) {
	trie := New().Insert("mayank")

	ShouldBeEqual(t, 1, len(trie.Children))
	ShouldBeEqual(t, 1, len(trie.Insert("mayank").Children))
	ShouldBeEqual(t, 2, len(trie.Insert("patel").Children))
	ShouldBeEqual(t, 3, len(trie.Insert("nahar").Children))
	ShouldBeEqual(t, 3, len(trie.Insert("mayankpatel").Children))

	ShouldBeEqual(t, false, trie.SearchWithWordBoundary("maya"))
	ShouldBeEqual(t, false, trie.SearchWithWordBoundary("foo"))
	ShouldBeEqual(t, true, trie.SearchWithWordBoundary("mayank"))
	ShouldBeEqual(t, true, trie.SearchWithWordBoundary("patel"))
	ShouldBeEqual(t, true, trie.SearchWithWordBoundary("mayankpatel"))

}

func TestRoot_Delete(t *testing.T) {
	trie := New().Insert("mayank")

	ShouldBeEqual(t, 1, len(trie.Children))
	ShouldBeEqual(t, 1, len(trie.Insert("mayank").Children))
	ShouldBeEqual(t, 2, len(trie.Insert("patel").Children))
	ShouldBeEqual(t, 3, len(trie.Insert("nahar").Children))

	ShouldBeEqual(t, true, trie.SearchWithWordBoundary("mayank"))
	ShouldBeEqual(t, true, trie.SearchWithWordBoundary("patel"))

	trie.Delete("maya")
	ShouldBeEqual(t, true, trie.Search("maya"))
	ShouldBeEqual(t, true, trie.Search("mayank"))
	ShouldBeEqual(t, true, trie.SearchWithWordBoundary("mayank"))

	trie.Delete("mayank")
	ShouldBeEqual(t, false, trie.Search("maya"))
	ShouldBeEqual(t, false, trie.Search("mayank"))
	ShouldBeEqual(t, false, trie.SearchWithWordBoundary("mayank"))
	ShouldBeEqual(t, true, trie.SearchWithWordBoundary("patel"))

	ShouldBeEqual(t, 4, len(trie.Insert("mayank").Children))
	ShouldBeEqual(t, 4, len(trie.Insert("mayankpatel").Children))
	ShouldBeEqual(t, true, trie.Search("maya"))
	ShouldBeEqual(t, true, trie.Search("mayank"))
	ShouldBeEqual(t, true, trie.SearchWithWordBoundary("mayank"))
	ShouldBeEqual(t, true, trie.SearchWithWordBoundary("mayankpatel"))
	trie.Delete("mayankpatel")
	ShouldBeEqual(t, true, trie.Search("maya"))
	ShouldBeEqual(t, true, trie.Search("mayank"))
	ShouldBeEqual(t, true, trie.SearchWithWordBoundary("mayank"))
	ShouldBeEqual(t, false, trie.SearchWithWordBoundary("mayankpatel"))

	ShouldBeEqual(t, 4, len(trie.Insert("mayank").Children))
	ShouldBeEqual(t, 4, len(trie.Insert("mayankpatel").Children))
	ShouldBeEqual(t, true, trie.Search("maya"))
	ShouldBeEqual(t, true, trie.Search("mayank"))
	ShouldBeEqual(t, true, trie.SearchWithWordBoundary("mayank"))
	ShouldBeEqual(t, true, trie.SearchWithWordBoundary("mayankpatel"))
	trie.Delete("mayank")
	ShouldBeEqual(t, true, trie.Search("maya"))
	ShouldBeEqual(t, true, trie.Search("mayank"))
	ShouldBeEqual(t, false, trie.SearchWithWordBoundary("mayank"))
	ShouldBeEqual(t, true, trie.SearchWithWordBoundary("mayankpatel"))

}
