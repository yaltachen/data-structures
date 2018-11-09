package trie

import (
	"fmt"
	"testing"
)

func TestTrie(t *testing.T) {
	trie := Trie()
	trie.Add("h")
	fmt.Println(trie.GetSize())
	trie.Add("hello")
	fmt.Println(trie.GetSize())
	trie.Add("hello")
	fmt.Println(trie.GetSize())
	trie.Add("hello1")
	fmt.Println(trie.GetSize())

	fmt.Println(trie.Contains("h"))
	fmt.Println(trie.Contains("he"))
	fmt.Println(trie.Contains("sad"))
	fmt.Println(trie.Contains("hello"))
}
