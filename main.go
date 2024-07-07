package main

import (
	"fmt"
	"strings"
)

// Node represents a single node in the Radix Tree.
type Node struct {
	prefix   string
	children []*Node
	isLeaf   bool
}

// RadixTree represents the entire Radix Tree.
type RadixTree struct {
	root *Node
}

// NewRadixTree creates and returns a new Radix Tree.
func NewRadixTree() *RadixTree {
	return &RadixTree{
		root: &Node{},
	}
}

// Insert inserts a string into the Radix Tree.
func (t *RadixTree) Insert(s string) {
	t.insert(t.root, s)
}

func (t *RadixTree) insert(node *Node, s string) {
	if s == "" {
		node.isLeaf = true
		return
	}

	for i := 0; i < len(s); i++ {
		char := s[i]
		child, _ := t.findChild(node, char)
		if child == nil {
			newChild := &Node{
				prefix: s[i:],
				isLeaf: true,
			}
			node.children = append(node.children, newChild)
			return
		}

		commonPrefixLength := commonPrefixLength(s[i:], child.prefix)
		if commonPrefixLength == len(child.prefix) {
			node = child
			i += commonPrefixLength - 1
			continue
		}

		newChild := &Node{
			prefix:   child.prefix[commonPrefixLength:],
			children: child.children,
			isLeaf:   child.isLeaf,
		}
		child.prefix = child.prefix[:commonPrefixLength]
		child.children = []*Node{newChild}
		child.isLeaf = false

		if commonPrefixLength < len(s[i:]) {
			newChild := &Node{
				prefix: s[i+commonPrefixLength:],
				isLeaf: true,
			}
			child.children = append(child.children, newChild)
		} else {
			child.isLeaf = true
		}
		return
	}
}

func (t *RadixTree) findChild(node *Node, char byte) (*Node, int) {
	for i, child := range node.children {
		if child.prefix[0] == char {
			return child, i
		}
	}
	return nil, -1
}

// Search searches for a string in the Radix Tree.
func (t *RadixTree) Search(s string) bool {
	return t.search(t.root, s)
}

func (t *RadixTree) search(node *Node, s string) bool {
	if s == "" {
		return node.isLeaf
	}

	char := s[0]
	child, _ := t.findChild(node, char)
	if child == nil {
		return false
	}

	if strings.HasPrefix(s, child.prefix) {
		return t.search(child, s[len(child.prefix):])
	}

	return false
}

func commonPrefixLength(a, b string) int {
	minLen := len(a)
	if len(b) < minLen {
		minLen = len(b)
	}
	for i := 0; i < minLen; i++ {
		if a[i] != b[i] {
			return i
		}
	}
	return minLen
}

// Print prints the Radix Tree.
func (t *RadixTree) Print() {
	t.print(t.root, 0)
}

func (t *RadixTree) print(node *Node, level int) {
	if node == nil {
		return
	}
	leafMark := ""
	if node.isLeaf {
		leafMark = " 🍀"
	}
	fmt.Printf("%s%s%s\n", strings.Repeat(" ", level), node.prefix, leafMark)
	for _, child := range node.children {
		t.print(child, level+len(node.prefix))
	}
}

func main() {
	tree := NewRadixTree()
	tree.Insert("hello")
	tree.Insert("hell")
	tree.Insert("heaven")
	tree.Insert("heavy")
	tree.Insert("Goo")
	tree.Insert("Google")
	tree.Insert("Golang")
	tree.Insert("Googlerr")

	fmt.Println(tree.Search("hell"))   // true
	fmt.Println(tree.Search("hello"))  // true
	fmt.Println(tree.Search("heaven")) // true
	fmt.Println(tree.Search("heavy"))  // true
	fmt.Println(tree.Search("heav"))   // false
	fmt.Println(tree.Search("helloo")) // false
	fmt.Println(tree.Search("Googler"))
	fmt.Println(tree.Search("Googlerr"))

	fmt.Println("\nRadix Tree Structure:")
	tree.Print()
}
