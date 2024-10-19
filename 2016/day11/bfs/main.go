package main

import (
	"fmt"
)

type node struct {
	v        string
	children []*node
}

func (n node) String() string {
	return n.v
}

func main() {
	munchen := node{
		v: "München",
	}
	stuttgart := node{
		v: "Stuttgart",
	}
	augsburg := node{
		v: "Augsburg",
	}
	erfurt := node{
		v: "Erfurt",
	}
	karlsruhe := node{
		v:        "Karlsruhe",
		children: []*node{&augsburg},
	}
	nurnberg := node{
		v:        "Nürnberg",
		children: []*node{&stuttgart},
	}
	mannheim := node{
		v:        "Mannheim",
		children: []*node{&karlsruhe},
	}
	wurzburg := node{
		v:        "Würzburg",
		children: []*node{&nurnberg, &erfurt},
	}
	kassel := node{
		v:        "Kassel",
		children: []*node{&munchen},
	}
	frankfurt := node{
		v:        "Frankfurt",
		children: []*node{&mannheim, &wurzburg, &kassel},
	}

	fmt.Println(bfs(&frankfurt))
	fmt.Println(bfs_rec(&frankfurt))
}

func bfs(root *node) []*node {
	queue := make([]*node, 0, 16)
	queue = append(queue, root)

	out := make([]*node, 0, 16)

	seen := make(map[string]bool, 16)
	seen[root.v] = true

	for len(queue) > 0 {
		n := queue[0]
		queue = queue[1:]
		seen[n.v] = true

		out = append(out, n)

		for _, c := range n.children {
			if !seen[c.v] {
				queue = append(queue, c)
			}
		}
	}
	return out
}

func bfs_rec(root *node) []*node {
	queue := make([]*node, 0, 16)
	queue = append(queue, root)

	out := make([]*node, 0, 16)
	seen := make(map[string]bool, 16)

	var rec func()
	rec = func() {
		if len(queue) == 0 {
			return
		}

		n := queue[0]
		queue = queue[1:]

		seen[n.v] = true

		out = append(out, n)

		for _, c := range n.children {
			if !seen[c.v] {
				queue = append(queue, c)
			}
		}
		rec()
	}

	rec()

	return out
}
