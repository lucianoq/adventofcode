package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func parse() *Node {
	fs := &Node{
		Dir:      true,
		Children: map[string]*Node{},
	}
	fs.AddDir("/")

	pwd := fs
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		ff := strings.Fields(scanner.Text())

		switch ff[0] {

		case "$":

			if ff[1] == "cd" {
				pwd = pwd.GetChild(ff[2])
			}

		case "dir":
			pwd.AddDir(ff[1])

		default:
			size, _ := strconv.Atoi(ff[0])
			pwd.AddFile(ff[1], size)
		}
	}

	fs.UpdateFolderSize()
	return fs.GetChild("/")
}

type Node struct {
	Name     string
	Parent   *Node
	Children map[string]*Node
	Dir      bool
	Size     int
}

func (n *Node) AddFile(name string, size int) {
	n.Children[name] = &Node{
		Name:   name,
		Parent: n,
		Size:   size,
	}
}

func (n *Node) AddDir(name string) {
	n.Children[name] = &Node{
		Name:     name,
		Parent:   n,
		Dir:      true,
		Children: map[string]*Node{},
	}
}

func (n *Node) GetChild(name string) *Node {
	if name == ".." {
		return n.Parent
	}
	return n.Children[name]
}

func (n *Node) GetSize() int {
	if !n.Dir {
		return n.Size
	}
	sum := 0
	for _, c := range n.Children {
		sum += c.GetSize()
	}
	return sum
}

func (n *Node) UpdateFolderSize() {
	if !n.Dir {
		return
	}
	n.Size = 0
	sum := 0
	for _, c := range n.Children {
		c.UpdateFolderSize()
		sum += c.Size
	}
	n.Size = sum
}
