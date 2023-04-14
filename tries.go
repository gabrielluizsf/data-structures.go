package datastructuresgo

import (
	"fmt"
)

type nodeT struct {
	children [26]*nodeT
	isWord   bool
}
type Trie struct {
	root *nodeT
}


func NewTrie() *Trie {
	return &Trie{
		root: &nodeT{},
	}
}

func (t *Trie) Insert(word string) {
	curr := t.root
	for i := 0; i < len(word); i++ {
		index := word[i] - 'a'
		if curr.children[index] == nil {
			curr.children[index] = &nodeT{}
		}
		curr = curr.children[index]
	}
	curr.isWord = true
}

func (t *Trie) Search(word string) bool {
	curr := t.root
	for i := 0; i < len(word); i++ {
		index := word[i] - 'a'
		if curr.children[index] == nil {
			return false
		}
		curr = curr.children[index]
	}
	return curr.isWord
}

func YesOrNo(trie *Trie, name string)string{
	if trie.Search(name) == true{
		return "Sim";
	}else{
		return "Não";
	}
}

func Tries() {

	PrintMessage("\nTries:")

	trie := NewTrie()
	trie.Insert("apple")
	trie.Insert("banana")
	trie.Insert("cat")

	fmt.Println("Tem Maçã?",YesOrNo(trie,"apple")) 
	fmt.Println("Tem Banana?",YesOrNo(trie,"banana")) 
	fmt.Println("Tem Gato?",YesOrNo(trie,"cat")) 
	fmt.Println("Tem Cachorro?",YesOrNo(trie,"dog")) 
}
