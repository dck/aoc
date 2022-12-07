package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type NodeType int

const (
	File NodeType = iota
	Directory
)

type Tree struct {
	Parent   *Tree
	Children []*Tree
	Type     NodeType
	Name     string
	Size     uint64
}

func (t *Tree) AddFile(name string, size uint64) {
	node := &Tree{Type: File, Name: name, Size: size, Parent: t}

	t.Children = append(t.Children, node)

	temp := t
	for temp != nil {
		temp.Size += size
		temp = temp.Parent
	}
}

func (t *Tree) AddDirectory(name string) {
	node := &Tree{Type: Directory, Name: name, Parent: t}

	t.Children = append(t.Children, node)
}

func (t *Tree) GetNodeByName(name string) *Tree {
	var result *Tree = nil
	for _, n := range t.Children {
		if n.Name == name {
			result = n
			break
		}
	}

	return result
}

func (t *Tree) String() string {
	var s strings.Builder
	t.printTree(&s, 1)

	return s.String()
}

func (t *Tree) printTree(s *strings.Builder, level int) {
	if t.Type == Directory {
		fmt.Fprintf(s, "%*s %s (dir, size=%d)\n", level, "+", t.Name, t.Size)

		for _, n := range t.Children {
			n.printTree(s, level+1)
		}
	} else {
		fmt.Fprintf(s, "%*s %s (file, size=%d)\n", level, "-", t.Name, t.Size)
	}
}

func main() {
	var filesystem *Tree
	var cursor = filesystem

	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		line := scanner.Text()
		parts := strings.Fields(line)

		if parts[0] == "$" {
			if parts[1] == "cd" {
				if parts[2] == "/" {
					filesystem = &Tree{Type: Directory, Name: "/"}
					cursor = filesystem
				} else if parts[2] == ".." {
					cursor = cursor.Parent
				} else {
					subfolder := cursor.GetNodeByName(parts[2])
					if subfolder == nil {
						panic("No such subfolder")
					}

					cursor = subfolder
				}
			}
		} else {
			if parts[0] == "dir" {
				cursor.AddDirectory(parts[1])
			} else {
				size, _ := strconv.ParseUint(parts[0], 10, 64)
				cursor.AddFile(parts[1], size)
			}
		}
	}

	fmt.Println(totalSizeBelowThreshold(filesystem, 100000))
	fmt.Println(findSmallestDirToRemove(filesystem, 40000000))
}

func totalSizeBelowThreshold(t *Tree, threshold uint64) uint64 {
	queue := []*Tree{t}
	total := uint64(0)

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node.Type == Directory && node.Size < threshold {
			total += node.Size
		}
		for _, n := range node.Children {
			if node.Type == Directory {
				queue = append(queue, n)
			}
		}
	}

	return total
}

func findSmallestDirToRemove(t *Tree, targetSpace uint64) uint64 {
	root := t
	queue := []*Tree{t}
	candidates := []*Tree{}

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		if node.Type == Directory && root.Size-node.Size <= targetSpace {
			candidates = append(candidates, node)
		}

		for _, n := range node.Children {
			if node.Type == Directory {
				queue = append(queue, n)
			}
		}
	}

	sort.Slice(candidates, func(i, j int) bool {
		return candidates[i].Size < candidates[j].Size
	})

	return candidates[0].Size
}
