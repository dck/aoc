package main

import (
	"bufio"
	"fmt"
	"os"
)

type TrieNode struct {
	children map[rune]*TrieNode
	end      bool
	value    int
}

func NewTrieNode() *TrieNode {
	return &TrieNode{children: make(map[rune]*TrieNode)}
}

func (t *TrieNode) Insert(s string, value int) {
	node := t
	for _, c := range s {
		if node.children[c] == nil {
			node.children[c] = NewTrieNode()
		}
		node = node.children[c]
	}
	node.end = true
	node.value = value

	node = t
	for i := range s {
		c := rune(s[len(s)-1-i])
		if node.children[c] == nil {
			node.children[c] = NewTrieNode()
		}
		node = node.children[c]
	}
	node.end = true
	node.value = value
}
func (t *TrieNode) PrintTree(depth int) {
	for c, n := range t.children {
		for i := 0; i < depth; i++ {
			fmt.Print(" ")
		}
		fmt.Printf("%c : %d\n", c, n.value)
		if n.end {
			for i := 0; i < depth+1; i++ {
				fmt.Print(" ")
			}
			fmt.Println("-- end of word --")
		}
		n.PrintTree(depth + 1)
	}
}

var trie *TrieNode

func main() {
	trie = NewTrieNode()
	trie.Insert("1", 1)
	trie.Insert("one", 1)
	trie.Insert("2", 2)
	trie.Insert("two", 2)
	trie.Insert("3", 3)
	trie.Insert("three", 3)
	trie.Insert("4", 4)
	trie.Insert("four", 4)
	trie.Insert("5", 5)
	trie.Insert("five", 5)
	trie.Insert("6", 6)
	trie.Insert("six", 6)
	trie.Insert("7", 7)
	trie.Insert("seven", 7)
	trie.Insert("8", 8)
	trie.Insert("eight", 8)
	trie.Insert("9", 9)
	trie.Insert("nine", 9)

	scanner := bufio.NewScanner(os.Stdin)

	total := 0
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			continue
		}
		value := decodeCalibrationValue(line)
		fmt.Println(line, value)
		total += value
	}
	fmt.Println(total)
}

func decodeCalibrationValue(line string) int {
	firstDigit := findFirstDigit(line, 0, trie)
	lastDigit := findFirstDigit(reversed(line), 0, trie)
	return 10*firstDigit + lastDigit
}

func findFirstDigit(s string, index int, node *TrieNode) int {
	for i := index; i < len(s); i++ {
		c := rune(s[i])
		if node.children[c] == nil {
			if index > 0 {
				break
			} else {
				continue
			}
		}
		if node.children[c].end {
			return node.children[c].value
		}
		val := findFirstDigit(s, i+1, node.children[c])
		if val != -1 {
			return val
		}
	}
	return -1
}

func reversed(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
