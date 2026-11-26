package main

import "fmt"

type Node struct {
	char byte
	freq uint
}

func countFreq(in string) []Node {
	nodes := []Node{}

	for i := 0; i < len(in); i++ {
		char := in[i]
		if char > 0 && char < 255 {
			found := false
			for j := range nodes {
				if nodes[j].char == char {
					nodes[j].freq++
					found = true
					break
				}
			}
			if !found {
				nodes = append(nodes, Node{char: char, freq: 1})
			}
		}
	}

	return nodes
}

func main() {
	// map[byte]count(uint) => Node

	// byte is 8 bit, each char is a byte
	// we need to present the char in fewer than a byte
	// A: 10, B: 8, C: 7, D:5, E:3 => total is 33
	nodes := countFreq("ADDAAAABBBBCCBDDDBEBBAACECCCCAEAA")
	var total uint

	for _, node := range nodes {
		fmt.Printf("%c: %d\n", node.char, node.freq)
		total += node.freq
	}

	fmt.Printf("total: %d\n", total)
}
