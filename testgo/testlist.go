package main

import "fmt"

func main() {
	// START OMIT
	type Node struct {
		Next  *Node
		Value interface{}
	}
	var first *Node

	visited := make(map[*Node]bool) // HL
	for n := first; n != nil; n = n.Next {
		if visited[n] { // HL
			fmt.Println("cycle detected")
			break
		}
		visited[n] = true // HL
		fmt.Println(n.Value)
	}
	// END OMIT

}
