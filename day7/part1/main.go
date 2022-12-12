package main

import (
	u "advent-of-code-2022/utilities"
	"bufio"
	"strconv"
	"strings"
)

func main() {
	scanner := u.ScannerFromFile("day7/input")
	root := u.Node{
		Name:     "/",
		Val:      0,
		Children: nil,
		Parent:   nil,
	}

	buildTree(&root, scanner)
	updateSizes(&root)
	dirs := getDirsUnderThreshold(&root, 100000)
	sum := 0
	for _, dir := range *dirs {
		sum += dir.Val
	}
	println(sum)
}

func getDirsUnderThreshold(n *u.Node, threshold int) *[]u.Node {
	var dirs []u.Node
	for _, child := range n.Children {
		if child.Val > threshold {
			continue
		}
		dirs = append(dirs, *child)
	}
	for _, child := range n.Children {
		dirs = append(dirs, *getDirsUnderThreshold(child, threshold)...)
	}
	return &dirs
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
