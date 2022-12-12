package main

import (
	u "advent-of-code-2022/utilities"
	"bufio"
	"sort"
	"strconv"
	"strings"
)

func main() {
	scanner := u.ScannerFromFile("day7/input")
	const maxSize = 70000000
	const freeSpaceNeeded = 30000000
	root := u.Node{
		Name:     "/",
		Val:      0,
		Children: nil,
		Parent:   nil,
	}

	buildTree(&root, scanner)
	updateSizes(&root)

	currentFreeSpace := maxSize - root.Val
	freeSpaceToAdd := freeSpaceNeeded - currentFreeSpace
	sortedDirs := getSortedDirs(&root)

	for _, dir := range sortedDirs {
		if dir.Val >= freeSpaceToAdd {
			println(dir.Val)
			return
		}
	}
}

func getSortedDirs(n *u.Node) []u.Node {
	dirsArray := nodesToArray(n)
	sort.Sort(sortNodes(dirsArray))
	return dirsArray
}

func nodesToArray(n *u.Node) []u.Node {
	var nodes []u.Node
	for _, child := range n.Children {
		nodes = append(nodes, *child)
	}
	for _, child := range n.Children {
		nodes = append(nodes, nodesToArray(child)...)
	}
	return nodes
}

type sortNodes []u.Node

func (s sortNodes) Less(i, j int) bool {
	return s[i].Val < s[j].Val
}

func (s sortNodes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortNodes) Len() int {
	return len(s)
}

func updateSizes(n *u.Node) int {
	for _, child := range n.Children {
		n.Val += updateSizes(child)
	}
	return n.Val
}

func buildTree(node *u.Node, scanner *bufio.Scanner) {
	currentNode := node
	for scanner.Scan() {
		line := scanner.Text()
		if line == "$ ls" {
			continue
		}
		split := strings.Split(line, " ")
		if split[0] == "$" {
			switch split[1] {
			case "cd":
				// Change directory
				if split[2] == ".." {
					// Go up one directory
					if currentNode.Parent != nil {
						currentNode = currentNode.Parent
					}
				} else {
					// Go down one directory
					for _, child := range currentNode.Children {
						if child.Name == split[2] {
							currentNode = child
							break
						}
					}
				}
			}
		} else {
			if split[0] == "dir" {
				// Create directory
				newNode := u.Node{
					Name:     split[1],
					Val:      0,
					Children: nil,
					Parent:   currentNode,
				}
				currentNode.Children = append(currentNode.Children, &newNode)
			} else {
				// Set node with file size
				val, _ := strconv.Atoi(split[0])
				currentNode.Val += val
			}
		}
	}
}
