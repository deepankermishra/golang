package main

import "fmt"

type trie struct {
	alphabet map[rune]*trie
	isTerminal bool
}

func newTrie() *trie {
	tr := trie{}
	tr.alphabet = make(map[rune]*trie)
	return &tr
}

func (tr *trie) insert(word string) {
	cur := tr
	for _, ch := range word {
		if cur.alphabet[ch] == nil {
			cur.alphabet[ch] = newTrie()			
		}
		cur = cur.alphabet[ch]
	}
	cur.isTerminal = true
}

func (tr *trie) printTrie() {
	for ch, subTr := range tr.alphabet {
		fmt.Println(string(ch))
		subTr.printTrie()
	}
}


func main() {
	tr := newTrie()
	tr.insert("word")
	tr.printTrie()
	fmt.Println("Done")
}
