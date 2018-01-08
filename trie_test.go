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

	//Duplicate entry should not be created
	ShouldBeEqual(t, 3, len(trie.Insert("mayank", "nahar").Children))
}

func TestRoot_Search(t *testing.T) {
	trie := New().Insert("mayank")

	ShouldBeEqual(t, 1, len(trie.Children))
	ShouldBeEqual(t, 1, len(trie.Insert("mayank").Children))
	ShouldBeEqual(t, 2, len(trie.Insert("patel").Children))
	ShouldBeEqual(t, 3, len(trie.Insert("nahar").Children))

	ShouldBeEqual(t, true, trie.PrefixSearch("maya"))
	ShouldBeEqual(t, false, trie.PrefixSearch("foo"))
	ShouldBeEqual(t, true, trie.PrefixSearch("mayank"))
	ShouldBeEqual(t, true, trie.PrefixSearch("patel"))
}

func TestRoot_SearchWithWordBoundary(t *testing.T) {
	trie := New().Insert("mayank")

	ShouldBeEqual(t, 1, len(trie.Children))
	ShouldBeEqual(t, 1, len(trie.Insert("mayank").Children))
	ShouldBeEqual(t, 2, len(trie.Insert("patel").Children))
	ShouldBeEqual(t, 3, len(trie.Insert("nahar").Children))
	ShouldBeEqual(t, 3, len(trie.Insert("mayankpatel").Children))

	ShouldBeEqual(t, false, trie.Search("maya"))
	ShouldBeEqual(t, false, trie.Search("foo"))
	ShouldBeEqual(t, true, trie.Search("mayank"))
	ShouldBeEqual(t, true, trie.Search("patel"))
	ShouldBeEqual(t, true, trie.Search("mayankpatel"))

}

func TestRoot_Delete(t *testing.T) {
	trie := New().Insert("mayank")

	ShouldBeEqual(t, 1, len(trie.Children))
	ShouldBeEqual(t, 1, len(trie.Insert("mayank").Children))
	ShouldBeEqual(t, 2, len(trie.Insert("patel").Children))
	ShouldBeEqual(t, 3, len(trie.Insert("nahar").Children))

	ShouldBeEqual(t, true, trie.Search("mayank"))
	ShouldBeEqual(t, true, trie.Search("patel"))

	trie.Delete("maya")
	ShouldBeEqual(t, true, trie.PrefixSearch("maya"))
	ShouldBeEqual(t, true, trie.PrefixSearch("mayank"))
	ShouldBeEqual(t, true, trie.Search("mayank"))

	trie.Delete("mayank")
	ShouldBeEqual(t, false, trie.PrefixSearch("maya"))
	ShouldBeEqual(t, false, trie.PrefixSearch("mayank"))
	ShouldBeEqual(t, false, trie.Search("mayank"))
	ShouldBeEqual(t, true, trie.Search("patel"))

	ShouldBeEqual(t, 4, len(trie.Insert("mayank").Children))
	ShouldBeEqual(t, 4, len(trie.Insert("mayankpatel").Children))
	ShouldBeEqual(t, true, trie.PrefixSearch("maya"))
	ShouldBeEqual(t, true, trie.PrefixSearch("mayank"))
	ShouldBeEqual(t, true, trie.Search("mayank"))
	ShouldBeEqual(t, true, trie.Search("mayankpatel"))
	trie.Delete("mayankpatel")
	ShouldBeEqual(t, true, trie.PrefixSearch("maya"))
	ShouldBeEqual(t, true, trie.PrefixSearch("mayank"))
	ShouldBeEqual(t, true, trie.Search("mayank"))
	ShouldBeEqual(t, false, trie.Search("mayankpatel"))

	ShouldBeEqual(t, 4, len(trie.Insert("mayank").Children))
	ShouldBeEqual(t, 4, len(trie.Insert("mayankpatel").Children))
	ShouldBeEqual(t, true, trie.PrefixSearch("maya"))
	ShouldBeEqual(t, true, trie.PrefixSearch("mayank"))
	ShouldBeEqual(t, true, trie.Search("mayank"))
	ShouldBeEqual(t, true, trie.Search("mayankpatel"))
	trie.Delete("mayank")
	ShouldBeEqual(t, true, trie.PrefixSearch("maya"))
	ShouldBeEqual(t, true, trie.PrefixSearch("mayank"))
	ShouldBeEqual(t, false, trie.Search("mayank"))
	ShouldBeEqual(t, true, trie.Search("mayankpatel"))

}

func TestRoot_DeleteBranch(t *testing.T) {
	trie := New().Insert("mayank")
	ShouldBeEqual(t, 1, len(trie.Children))

	ShouldBeEqual(t, true, trie.Search("mayank"))
	ShouldBeEqual(t, true, trie.PrefixSearch("maya"))
	ShouldBeEqual(t, false, trie.Search("maya"))

	trie.DeleteBranch("maya")
	ShouldBeEqual(t, false, trie.Search("mayank"))
	ShouldBeEqual(t, true, trie.Search("maya"))
	ShouldBeEqual(t, true, trie.PrefixSearch("maya"))
}
